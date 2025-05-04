package handler

import (
	"documentum/pkg/logger"
	"documentum/pkg/models"
	"documentum/pkg/service/document"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type DocHandler struct {
	log     logger.Logger
	service document.DocService
}

func NewDocHandler(log logger.Logger, service document.DocService) *DocHandler {
	return &DocHandler{
		log:     log,
		service: service,
	}
}

func (h *DocHandler) GetDocuments(w http.ResponseWriter, r *http.Request) {

	query := r.URL.Query()
    
    settings := models.DocSettings{
        DocType:   query.Get("type"),
        DocCol:    query.Get("col"),
        DocSet:    query.Get("set"),
        DocDatain: query.Get("datain"),
        DocDatato: query.Get("datato"),
    }
    
    // Валидация хотя бы одного обязательного параметра
    if settings.DocType == "" {
		h.log.Error(models.ErrRequest)
		http.Error(w, models.ErrRequest, 400)
        return
    }

	documents, err := h.service.GetDocuments(settings)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(documents)
}

func (h *DocHandler) AddLookDocument(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	login := r.Context().Value(models.LoginKey).(string)
	id := vars["id"]

	err := h.service.AddLookDocument(id, login)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *DocHandler) AddDocument(w http.ResponseWriter, r *http.Request) {

	login := r.Context().Value(models.LoginKey).(string)

	countStr := r.FormValue("count")

	count, err := strconv.Atoi(countStr)
	if err != nil {
		http.Error(w, models.ErrRequest, http.StatusBadRequest)
		h.log.Error(models.ErrRequest, err)
		return
	}

	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, models.ErrRequest, http.StatusBadRequest)
		h.log.Error(models.ErrRequest, err)
		return
	}
	defer file.Close()

	document := models.Document{
		Type:       r.FormValue("type"),
		FNum:       r.FormValue("fnum"),
		FDate:      r.FormValue("fdate"),
		LNum:       r.FormValue("lnum"),
		LDateStr:   r.FormValue("ldate"),
		Name:       r.FormValue("name"),
		Sender:     r.FormValue("sender"),
		Ispolnitel: r.FormValue("ispolnitel"),
		Result:     r.FormValue("result"),
		Familiar:   r.FormValue("familiar"),
		Count:      count,
		Copy:       r.FormValue("copy"),
		Width:      r.FormValue("width"),
		Location:   r.FormValue("location"),
		Creator:    login,
		File:       file,
		FileHeader: header,
	}

	switch document.Type {
	case "Входящий":
		resolutionsJSON := r.FormValue("resolutions")
		if err := json.Unmarshal([]byte(resolutionsJSON), &document.Resolutions); err != nil {
			http.Error(w, models.ErrRequest, http.StatusBadRequest)
			h.log.Error(models.ErrRequest, err)
			return
		}
		doc, err := h.service.AddIngoingDoc(document)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")

		// Кодируем данные в JSON и отправляем
		if err := json.NewEncoder(w).Encode(doc); err != nil {
			http.Error(w, "Ошибка формирования ответа", http.StatusInternalServerError)
			return
		}
	case "Исходящий":
		doc, err := h.service.AddOutgoingDoc(document)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")

		// Кодируем данные в JSON и отправляем
		if err := json.NewEncoder(w).Encode(doc); err != nil {
			http.Error(w, "Ошибка формирования ответа", http.StatusInternalServerError)
			return
		}
	default:
	}
}
