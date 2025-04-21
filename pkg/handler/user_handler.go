package handler

import (
	"net/http"
	"documentum/pkg/service"
	"fmt"
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