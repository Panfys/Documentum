package auth

import (
	"documentum/pkg/models"
	"documentum/pkg/service/valid"
	"documentum/pkg/storage"
	"documentum/pkg/logger"
	"errors"

	"github.com/golang-jwt/jwt/v5"

	"fmt"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	UserRegistration(user models.User) error
	UserAuthorization(login, pass string) (int, error)
	CheckUserTokenToValid(token string) (string, error)
	GenerateToken(w http.ResponseWriter, login, remember string) error
}

type authService struct {
	log  logger.Logger
	stor storage.AuthStorage 
	valid valid.UserValidatService 
	secretKey []byte
}

func NewAuthService(log  logger.Logger, stor storage.AuthStorage, valid valid.UserValidatService , secretKey string) AuthService {
	return &authService{
		log: log,
		stor: stor,
		valid: valid,
		secretKey: []byte(secretKey),
	}
}

func (s *authService) GenerateToken(w http.ResponseWriter, login, remember string)  error {

	claims := jwt.RegisteredClaims{
		Subject:   login,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	responceToken, err := token.SignedString(s.secretKey) 
	if err != nil {
		return s.log.Error(models.ErrTokenAuth, err)
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
			return nil, s.log.Error(models.ErrTokenAuth, token.Header["alg"])
		}
		return s.secretKey, nil
	})

	if err != nil {
		return "", s.log.Error(models.ErrTokenAuth, err)
	}

	if claims, ok := token.Claims.(*jwt.RegisteredClaims); ok && token.Valid {
		return claims.Subject, nil
	}

	return "", s.log.Error(models.ErrTokenAuth, err)
}

func (s *authService) UserRegistration(user models.User) error {

	if exists, err := s.stor.GetUserExists(user.Login); err != nil {
		return err
	} else if exists {
		return fmt.Errorf("пользователь с логином '%s' уже существует", user.Login)
	}
	if !s.valid.ValidUserLogin(user.Login) {
		return errors.New("неверный формат логина: " + user.Login)
	}

	if !s.valid.ValidUserName(user.Name) {
		return errors.New("неверный формат ФИО")
	}

	if user.Func == "0" {
		return errors.New("должность не указана")
	}

	if !s.valid.ValidUserPass(user.Pass) {
		return errors.New("неверный формат пароля")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Pass), bcrypt.DefaultCost)
	if err != nil {
		return s.log.Error(models.ErrHashPass, err)
	}

	user.Pass = string(hashedPassword)

	if err := s.stor.AddUser(user); err != nil {
		return err
	}

	return nil
}

func (s *authService) UserAuthorization(login, pass string) (int, error) {
	userPass, err := s.stor.GetUserPassByLogin(login)
	if err != nil {
		return 500, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userPass), []byte(pass)); err != nil {
		return 401, errors.New("неверный логин или пароль")
	}

	return 0, nil 
}
