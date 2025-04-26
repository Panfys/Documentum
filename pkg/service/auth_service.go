package service

import (
	"documentum/pkg/models"
	"documentum/pkg/storage"
	"errors"

	"github.com/golang-jwt/jwt/v5"

	"fmt"
	"golang.org/x/crypto/bcrypt"
	"time"
	"net/http"
)

type AuthService interface {
	UserRegistration(user models.User) error
	UserAuthorization(login, pass string) (int, error)
	CheckUserTokenToValid(token string) (string, error)
	GenerateToken(w http.ResponseWriter, login, remember string) error
	GetAccountData(login string) (models.AccountData, error)
	GetFuncs() ([]models.Unit, error)
}

type authService struct {
	storage   storage.AuthStorage
	secretKey []byte
}

func NewAuthService(storage storage.AuthStorage, secretKey string) AuthService {
	return &authService{
		storage:   storage,
		secretKey: []byte(secretKey),
	}
}

func (s *authService) GetFuncs() ([]models.Unit, error) {
	var funcs []models.Unit 
	funcs, err := s.storage.GetFuncs()
	if err != nil {
		return nil, err
	}
	return funcs, err
}

func (s *authService) GenerateToken(w http.ResponseWriter, login, remember string)  error {

	claims := jwt.RegisteredClaims{
		Subject:   login,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	responceToken, err := token.SignedString(s.secretKey) 
	if err != nil {
		return errors.New("ошибка создания токена")
	}

	if remember == "true"{
		http.SetCookie(w, &http.Cookie{
			Name:     "token", // Имя куки
			Value:    responceToken,   // Значение (наш JWT-токен)
			Expires:  time.Now().Add(72 * time.Hour),
			Path:     "/",     // Доступно для всех путей на сайте
			HttpOnly: true,    // Защита от XSS (недоступно через JavaScript)
			Secure:   false,    // Только через HTTPS (в production)
			SameSite: http.SameSiteLaxMode, // Защита от CSRF
		})
	} else {
		http.SetCookie(w, &http.Cookie{
			Name:     "token", // Имя куки
			Value:    responceToken,   // Значение (наш JWT-токен)
			Path:     "/",     // Доступно для всех путей на сайте
			HttpOnly: true,    // Защита от XSS (недоступно через JavaScript)
			Secure:   false,    // Только через HTTPS (в production)
			SameSite: http.SameSiteLaxMode, // Защита от CSRF
		})
	}

	return nil
}

func (s *authService) CheckUserTokenToValid(tokenString string) (string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("неизвестный метод шифрования токена: %v", token.Header["alg"])
		}
		return s.secretKey, nil
	})

	if err != nil {
		return "", fmt.Errorf("ошибка проверки токена: %w", err)
	}

	if claims, ok := token.Claims.(*jwt.RegisteredClaims); ok && token.Valid {
		return claims.Subject, nil
	}

	return "", errors.New("токен не валиден")
}

func (s *authService) UserRegistration(user models.User) error {

	if exists, err := s.storage.UserExists(user.Login); err != nil {
		return fmt.Errorf("ошибка при проверке существования пользователя")
	} else if exists {
		return fmt.Errorf("пользователь с логином '%s' уже существует", user.Login)
	}
	if !user.ValidLogin(user.Login) {
		return errors.New("неверный формат логина: " + user.Login)
	}

	if !user.ValidName(user.Name) {
		return errors.New("неверный формат ФИО")
	}

	if user.Func == "0" {
		return errors.New("должность не указана")
	}

	if !user.ValidPass(user.Pass) {
		return errors.New("неверный формат пароля")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Pass), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("ошибка хеширования пароля: %w", err)
	}

	user.Pass = string(hashedPassword)

	if err := s.storage.AddUser(user); err != nil {
		return fmt.Errorf("ошибка записи данных в БД: %w", err)
	}

	return nil
}

func (s *authService) UserAuthorization(login, pass string) (int, error) {
	userPass, err := s.storage.GetUserPassByLogin(login)
	if err != nil {
		return 500, fmt.Errorf("ошибка авторизации: %w", err)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userPass), []byte(pass)); err != nil {
		return 401, errors.New("неверный логин или пароль")
	}

	return 0, nil
}

func (s *authService) GetAccountData(login string) (models.AccountData, error) {
	var accountData models.AccountData
	accountData, err := s.storage.GetAccountData(login)

	if err != nil {
		return accountData, err
	}

	now := time.Now()

	accountData.Login = login
	accountData.ToDay = now.Format("2006-01-02")
	
	return accountData, nil
}	

/*

Метод для проверки токена пользователей

// Метод для авторизации пользователей
func (s *AuthService) ExitHandler(w http.ResponseWriter, r *http.Request) {

	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    "",
		Path:     "/",
		HttpOnly: true,  // Защита от доступа через JavaScript
		Secure:   false, // Используйте только по HTTPS
		Expires:  time.Now(),
	})

	s.EntranceHandler(w, r)
}

// Метод для изменения пароля пользователя
func (s *AuthService) ChangePasswordHandler(w http.ResponseWriter, r *http.Request) {
	// Извлекаем логин пользователя из контекста
	login, ok := r.Context().Value(userContextKey).(string)
	if !ok {
		http.Error(w, "не удалось получить данные пользователя", http.StatusUnauthorized)
		return
	}

	// Получаем данные из запроса
	CurrentPassword := r.FormValue("pass")
	NewPassword := r.FormValue("newpass")

	// Получаем хеш пароля пользователя из базы данных
	userPass, err := s.DB.GetUserPass(login)
	if err != nil {
		http.Error(w, "ошибка обработки данных", http.StatusUnauthorized)
		return
	}

	// Проверяем валидность текущего пароля
	if err := bcrypt.CompareHashAndPassword([]byte(userPass), []byte(CurrentPassword)); err != nil {
		http.Error(w, "Неверный текущий пароль!", http.StatusUnauthorized)
		return
	}

	// Валидация пароля
	if !s.validUserPass(NewPassword) {
		http.Error(w, "формат пароля указан неверно", 400)
		return
	}

	// Хешируем новый пароль
	newHash, err := bcrypt.GenerateFromPassword([]byte(NewPassword), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "ошибка хеширования нового пароля", http.StatusInternalServerError)
		return
	}

	// Обновляем пароль в базе данных
	err = s.DB.UpdateUserPassword(login, string(newHash))
	if err != nil {
		http.Error(w, "ошибка обновления пароля в БД", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Пароль успешно изменен!"))
}
*/