package handler

import (
	"documentum/pkg/models"
	"documentum/pkg/service"
	"encoding/json"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type AuthHandler struct {
	authService service.AuthService
}

func NewAuthHandler(authService service.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

func (p AuthHandler) RegistrationHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "ошибка обработки данных пользователя", 400)
		return
	}

	err := p.authService.UserRegistration(user)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}
