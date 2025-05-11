package handler

import (
	"documentum/pkg/logger"
	"documentum/pkg/models"
	"documentum/pkg/service/document"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
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
		h.log.Error(models.ErrRequest, nil)
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

	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		h.log.Error(models.ErrRequest, err)
		http.Error(w, models.ErrRequest, http.StatusBadRequest)
		return
	}

	jsonData := r.FormValue("document")
	var document models.Document
	if err := json.Unmarshal([]byte(jsonData), &document); err != nil {
		h.log.Error(models.ErrRequest, err)
		http.Error(w, models.ErrRequest, http.StatusBadRequest)
		return
	}

	document.File, document.FileHeader, err = r.FormFile("file")
	if err != nil {
		h.log.Error(models.ErrFileRequest, err)
		http.Error(w, models.ErrFileRequest, http.StatusBadRequest)
		return
	}
	defer document.File.Close()

	document.Creator = login

	switch document.Type {
	default:
		doc, err := h.service.AddDocument(document)
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
	}
}
