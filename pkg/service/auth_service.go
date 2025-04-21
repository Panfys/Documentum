package service

import (
	"documentum/pkg/models"
	"documentum/pkg/storage"
	"errors"

	"github.com/golang-jwt/jwt/v5"

	"fmt"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type AuthService interface {
	UserRegistration(user models.User) error
	UserAuthorization(login, pass, remember string) (string, error)
	CheckUserTokenToValid(token string) (string, error)
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

func (s *authService) generateToken(login string, remember bool) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour) 
	if remember {
		expirationTime = time.Now().Add(72 * time.Hour)
	}

	claims := jwt.RegisteredClaims{
		Subject:   login,
		ExpiresAt: jwt.NewNumericDate(expirationTime),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(s.secretKey)
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
	if err := validateUser(user); err != nil {
		return err
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

func (s *authService) UserAuthorization(login, pass, remember string) (string, error) {
	userPass, err := s.storage.GetUserPass(login)
	if err != nil {
		return "", fmt.Errorf("ошибка авторизации: %w", err)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userPass), []byte(pass)); err != nil {
		return "errPass", errors.New("неверный пароль")
	}

	rememberMe := remember == "true"
	token, err := s.generateToken(login, rememberMe)
	if err != nil {
		return "", fmt.Errorf("token generation failed: %w", err)
	}

	return token, nil
}

// Вспомогательная функция для валидации пользователя
func validateUser(user models.User) error {
	if !user.ValidLogin(user.Login) {
		return errors.New("неверный формат логина " + user.Login)
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

	return nil
}

/*

Метод для проверки токена пользователей
func (s *AuthService) CheckTokenHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Извлечение токена из заголовка Authorization
		cookie, err := r.Cookie("token")
		if err != nil {
			server.EntranceHandler(w, r)
			return
		}

		userToken := cookie.Value

		// Проверяем токен
		token, err := jwt.Parse(userToken, func(token *jwt.Token) (interface{}, error) {
			// Проверяем метод подписи
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("неизвестный метод шифрования: %v", token.Header["alg"])
			}
			return secretKey, nil
		})

		if err != nil || !token.Valid {
			s.EntranceHandler(w, r)
			return
		}

		// Извлечение полезных данных из токена
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			// Например, извлекаем поле "user_id"
			login := claims["login"].(string)

			// Создаем новый контекст с данными пользователя
			ctx := context.WithValue(r.Context(), userContextKey, login)

			// Обновляем запрос с новым контекстом
			r = r.WithContext(ctx)
		}

		next.ServeHTTP(w, r)
	})
}

// Метод для авторизации пользователей
func (s *AuthService) AuthorizationHandler(w http.ResponseWriter, r *http.Request) {
	login := r.FormValue("login")
	pass := r.FormValue("pass")
	remember := r.FormValue("remember")
	var exp time.Time

	if remember == "true" {
		exp = time.Now().Add(72 * time.Hour)
	}

	// Проверяем, есть ли пользователь с таким именем и паролем
	userPass, err := s.DB.GetUserPass(login)

	if err != nil {
		http.Error(w, "ошибка обработки данных", http.StatusUnauthorized)
		return
	}

	// Проверка валидности пароля
	if err := bcrypt.CompareHashAndPassword([]byte(userPass), []byte(pass)); err != nil {
		http.Error(w, "Неверный логин или пароль!", http.StatusUnauthorized)
		return
	}

	token, err := generateToken(login)

	if err != nil {
		http.Error(w, "ошибка создания токена", 500)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    token,
		Path:     "/",
		HttpOnly: true,  // Защита от доступа через JavaScript
		Secure:   false, // Используйте только по HTTPS
		Expires:  exp,
	})

	// Создаем новый контекст с данными пользователя
	ctx := context.WithValue(r.Context(), userContextKey, login)

	// Обновляем запрос с новым контекстом
	r = r.WithContext(ctx)

	s.MainHandler(w, r)
}

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