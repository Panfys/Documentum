package service

import (
	"documentum/pkg/storage"
	"log"

	"github.com/golang-jwt/jwt"

	//"golang.org/x/crypto/bcrypt"
	"fmt"
)

type AuthService interface {
	UserAuthorization(login string, pass string, remember string) error
	CheckUserTokenToValid(token string) (string, error)
}

type authService struct {
	storage   storage.AuthStorage
	secretKey string
}

func NewAuthService(storage storage.AuthStorage, secretKey string) AuthService {
	return &authService{
		storage:   storage,
		secretKey: secretKey,
	}
}

func (s *authService) generateToken(login string) (string, error) {

	claims := jwt.MapClaims{
		"login": login,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(s.secretKey)
}

func (s *authService) CheckUserTokenToValid(requestToken string) (string, error) {

	token, err := jwt.Parse(requestToken, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("неизвестный метод шифрования: %v", token.Header["alg"])
		}
		return s.secretKey, nil
	})

	if err != nil || !token.Valid { 
		log.Printf("Ошибка при валидации токена: %s", err)
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims["login"].(string), nil
	} else {
		log.Printf("Ошибка при получении информации из токена: %s", err)
		return "", err
	} 
}

func (s *authService) UserAuthorization(login string, pass string, remember string) error {
	/*
		var exp time.Time

		if remember == "true" {
			exp = time.Now().Add(72 * time.Hour)
		}

		// Проверяем, есть ли пользователь с таким именем и паролем
		userPass, err := s.db.GetUserPass(login)

		if err != nil {
			return err
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
	*/
	return nil
}

/*
// Секретный ключ

[]byte("EAjbsdr@415tgs**405FW")
// Определите ключ для контекста
type contextKey string

const userContextKey contextKey = "user"

func generateToken(login string) (string, error) {

	claims := jwt.MapClaims{
		"login": login,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

// Метод для регистарции пользователей
func (s *AuthService) RegistrationHandler(w http.ResponseWriter, r *http.Request) {
	var user User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "ошибка обработки данных пользователя", 400)
		return
	}

	// Валидация логина
	err := s.DB.ValidUserLoginDB(user.Login)

	if err != nil {
		http.Error(w, fmt.Sprintf("%s", err), 400)
		return
	}

	if !s.validUserLogin(user.Login) {
		http.Error(w, "формат логина указан неверно", 400)
		return
	}

	// Валидация ФИО
	if !s.validUserName(user.Name) {
		http.Error(w, "формат имени указан неверно", 400)
		return
	}

	// Валидация должности
	if user.Func == "0" {
		http.Error(w, "должность не указана", 400)
		return
	}

	// Валидация пароля
	if !s.validUserPass(user.Pass) {
		http.Error(w, "формат пароля указан неверно", 400)
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Pass), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "ошибка хеширования данных", 500)
		return
	}
	user.Pass = string(hash)

	err = s.DB.AddUser(user)

	if err != nil {
		http.Error(w, "ошибка записи данных", 500)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}

// Метод для проверки токена пользователей
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

// Метод для валидации ФИО пользователя
func (s *AuthService) validUserName(name string) bool {

	pattern := `^[А-ЯЁ][а-яё]+ [А-ЯЁ]\.[А-ЯЁ]\.$`
	re := regexp.MustCompile(pattern)

	return re.MatchString(name)

}

// Метод для валидации логина пользователя
func (s *AuthService) validUserLogin(login string) bool {

	pattern := `^[a-zA-Z0-9.]{3,12}$`
	re := regexp.MustCompile(pattern)

	return re.MatchString(login)
}

// Метод для валидации пароля пользователя
func (s *AuthService) validUserPass(pass string) bool {

	pattern := `^[a-zA-Z-ЯЁа-яё0-9.]{6,30}$`
	re := regexp.MustCompile(pattern)

	return re.MatchString(pass)
}*/
