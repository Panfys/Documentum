package handler

import (
	"documentum/pkg/models"
	"documentum/pkg/service/user"
	"net/http"
)

type UserHandler struct {
	service user.UserService
}

func NewUserHandler(service user.UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (h *UserHandler) UpdateUserPassword(w http.ResponseWriter, r *http.Request) {

	login := r.Context().Value(models.LoginKey).(string)
	pass := r.FormValue("pass")
	newPass := r.FormValue("newpass")
	
	status, err := h.service.UpdateUserPassword(login, pass, newPass)
	if err != nil {
		http.Error(w, err.Error(), status)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}

func (h *UserHandler) UpdateUserIcon(w http.ResponseWriter, r *http.Request) {

	login := r.Context().Value(models.LoginKey).(string)
	icon, header, err := r.FormFile("icon") 
	if err != nil {
		http.Error(w, "Ошибка получения файла иконки", http.StatusBadRequest)
		return
	}
	defer icon.Close()

	newFilename, err := h.service.UpdateUserIcon(login, icon, header.Filename)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(newFilename))
}
