package handler

import (
	"documentum/pkg/service/structure"
	"fmt"
	"net/http"
)

type StructureHandler struct {
	service structure.StructureService
}

func NewStructureHandler(service structure.StructureService) *StructureHandler{
	return &StructureHandler{
		service: service,
	}
}

func (h *StructureHandler) GetUnits(w http.ResponseWriter, r *http.Request) {
	function := r.FormValue("func")

	units, err := h.service.GetUnits(function)
	if err != nil {
		http.Error(w, fmt.Sprintf("ошибка получения данных: %s", err), 500)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(units))
}

func (h *StructureHandler) GetGroups(w http.ResponseWriter, r *http.Request) {
	function := r.FormValue("func")
	unit := r.FormValue("unit")

	groups, err := h.service.GetGroups(function, unit)
	if err != nil {
		http.Error(w, fmt.Sprintf("ошибка получения данных: %s", err), 500)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(groups))
}