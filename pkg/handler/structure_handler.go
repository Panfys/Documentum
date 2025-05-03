package handler

import (
	"documentum/pkg/service/structure"
	"net/http"
	"github.com/gorilla/mux"
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
	vars := mux.Vars(r)
    function := vars["id"]

	units, err := h.service.GetUnits(function)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(units))
}

func (h *StructureHandler) GetGroups(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	function := vars["funcId"]
	unit := vars["unitId"]

	groups, err := h.service.GetGroups(function, unit)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(groups))
}