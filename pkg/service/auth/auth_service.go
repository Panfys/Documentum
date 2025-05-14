package auth

import (
	"documentum/pkg/logger"
	"documentum/pkg/models"
	"documentum/pkg/service/valid"
	"documentum/pkg/storage"
	"errors"

	"github.com/golang-jwt/jwt/v5"

	"crypto/sha256"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

type AuthService interface {
	UserRegistration(user models.User) error
	UserAuthorization(login, pass string) (int, error)
	CheckUserTokenToValid(tokenString, agent, ip string) (string, error) 
	GenerateToken(w http.ResponseWriter, authData models.AuthData) error
}

type authService struct {
	log       logger.Logger
	stor      storage.AuthStorage
	valid     valid.UserValidatService
	secretKey []byte
}

func NewAuthService(log logger.Logger, stor storage.AuthStorage, valid valid.UserValidatService, secretKey string) AuthService {
	return &authService{
		log:       log,
		stor:      stor,
		valid:     valid,
		secretKey: []byte(secretKey),
	}
}

func (s *authService) GenerateToken(w http.ResponseWriter, authData models.AuthData) error {
	now := time.Now()
	expiresAt := now.Add(24 * time.Hour)
	if authData.Remember {
		expiresAt = now.Add(72 * time.Hour)
	}

	userFP := s.generateFingerprint(authData.IP, authData.Agent)

	claims := jwt.RegisteredClaims{
		Subject:   authData.Login,
		IssuedAt:  jwt.NewNumericDate(now),
		ExpiresAt: jwt.NewNumericDate(expiresAt),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	dynamicKey := append(s.secretKey, []byte(userFP)...)
	tokenString, err := token.SignedString(dynamicKey)
	if err != nil {
		return s.log.Error(models.ErrTokenAuth, err)
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    tokenString,
		Expires:  expiresAt,
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	})

	return nil
}

func (s *authService) CheckUserTokenToValid(tokenString, agent, ip string) (string, error) {

	currentFP := s.generateFingerprint(ip, agent)

	token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, s.log.Error(models.ErrTokenAuth, token.Header["alg"])
		}
		dynamicKey := append(s.secretKey, []byte(currentFP)...)
        return dynamicKey, nil
	})

	if err != nil {
		return "", s.log.Error(models.ErrTokenAuth, err)
	}
	
	if claims, ok := token.Claims.(*jwt.RegisteredClaims); ok && token.Valid &&claims.Subject != ""{
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

	return s.stor.AddUser(user)
}

func (s *authService) UserAuthorization(login, pass string) (int, error) {

	if exists, err := s.stor.GetUserExists(login); err != nil {
		return http.StatusInternalServerError, err
	} else if !exists {
		return http.StatusUnauthorized, errors.New("authError")
	}

	userPass, err := s.stor.GetUserPassByLogin(login)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userPass), []byte(pass)); err != nil {
		return http.StatusUnauthorized, errors.New("authError")
	}

	return http.StatusOK, nil
}

func (s *authService) generateFingerprint(ip, userAgent string) string {
	h := sha256.New()
	h.Write([]byte(ip))
	h.Write([]byte(userAgent))
	return fmt.Sprintf("%x", h.Sum(nil))
}
