package handler

import (
	"documentum/pkg/service/user_service"
	"documentum/pkg/models"
	"fmt"
	"net/http"
)

type UserHandler struct {
	userService user_service.UserService
}

func NewUserHandler(userService user_service.UserService) *UserHandler {
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

func (h *UserHandler) UpdateUserIcon(w http.ResponseWriter, r *http.Request) {

	login := r.Context().Value(models.LoginKey).(string);
	icon, header, err := r.FormFile("icon")
    if err != nil {
        http.Error(w, "Ошибка получения файла иконки", http.StatusBadRequest)
        return
    }
    defer icon.Close()

    newFilename, err := h.userService.UpdateUserIcon(login, icon, header.Filename)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    w.Write([]byte(newFilename))
}