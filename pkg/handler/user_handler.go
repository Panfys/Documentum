package handler

import (
	"documentum/pkg/service"
	"documentum/pkg/models"
	"fmt"
	"net/http"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) GetUnits(w http.ResponseWriter, r *http.Request) {
	function := r.FormValue("func")
	
	units, err := h.userService.GetUnits(function)
	if err != nil {
		http.Error(w, fmt.Sprintf("ошибка получения данных: %s", err), 500)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(units))
}

func (h *UserHandler) GetGroups(w http.ResponseWriter, r *http.Request) {
	function := r.FormValue("func")
	unit := r.FormValue("unit")

	groups, err := h.userService.GetGroups(function, unit)
	if err != nil {
		http.Error(w, fmt.Sprintf("ошибка получения данных: %s", err), 500)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(groups))
}

func (h *UserHandler) UpdateUserPassword(w http.ResponseWriter, r *http.Request) {

	login := r.Context().Value(models.LoginKey).(string);
	pass := r.FormValue("pass")
	newPass := r.FormValue("newpass")
	status, err := h.userService.UpdateUserPassword(login, pass, newPass)
	if err != nil {
		http.Error(w, err.Error(), status)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}