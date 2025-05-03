package handler

import (
	"documentum/pkg/logger"
	"documentum/pkg/models"
	"documentum/pkg/service/user"
	"encoding/json"
	"net/http"
)

type UserHandler struct {
	log     logger.Logger
	service user.UserService
}

func NewUserHandler(log logger.Logger, service user.UserService) *UserHandler {
	return &UserHandler{
		log: log,
		service: service,
	}
}

func (h *UserHandler) UpdateUserPassword(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	login := r.Context().Value(models.LoginKey).(string)

	var request models.UpdatePassRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		h.log.Error(models.ErrRequest, err)
		http.Error(w, models.ErrRequest, 400)
		return
	}

	status, err := h.service.UpdateUserPassword(login, request.Pass, request.NewPass)
	if err != nil {
		http.Error(w, err.Error(), status)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *UserHandler) UpdateUserIcon(w http.ResponseWriter, r *http.Request) {

	login := r.Context().Value(models.LoginKey).(string)
	icon, header, err := r.FormFile("icon")
	if err != nil {
		http.Error(w, "Ошибка обработки файла", http.StatusBadRequest)
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
