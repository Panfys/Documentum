package handler

import (
	"documentum/pkg/logger"
	"documentum/pkg/models"
	"documentum/pkg/service/document"
	"encoding/json"
	"github.com/gorilla/mux"
	"mime/multipart"
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
	vars := mux.Vars(r)
	table := vars["table"]

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

	w.Header().Set("Content-Type", "application/json")

	switch table {
	case "directives":
		directives, err := h.service.GetDirectives(settings)

		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		json.NewEncoder(w).Encode(directives)
	case "inventory":
		inventory, err := h.service.GetInventory(settings)

		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		json.NewEncoder(w).Encode(inventory)
	default:
		documents, err := h.service.GetDocuments(settings)

		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		json.NewEncoder(w).Encode(documents)
	}
}

func (h *DocHandler) UpdateDocFamiliar(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	login := r.Context().Value(models.LoginKey).(string)
	id := vars["id"]
	types := vars["type"]

	err := h.service.UpdateDocFamiliar(types, id, login)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *DocHandler) AddDocument(w http.ResponseWriter, r *http.Request) {

	login := r.Context().Value(models.LoginKey).(string)
	table := mux.Vars(r)["table"]

	if err := r.ParseMultipartForm(10 << 20); err != nil {
		h.handleError(w, models.ErrRequest, err, http.StatusBadRequest)
		return
	}

	file, header, err := r.FormFile("file")
	if err != nil {
		h.handleError(w, models.ErrFileRequest, err, http.StatusBadRequest)
		return
	}
	defer file.Close()

	switch table {
	case "directives":
		err = h.addDirective(r, file, header, login)
	case "inventory":
		err = h.addInventory(r, file, header, login)
	default:
		err = h.addInOutDoc(r, file, header, login)
	}

	if err != nil {
		h.handleError(w, err.Error(), err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *DocHandler) UpdateDocResolutions(w http.ResponseWriter, r *http.Request) {

	login := r.Context().Value(models.LoginKey).(string)
	table := mux.Vars(r)["table"]

	if err := r.ParseMultipartForm(10 << 20); err != nil {
		h.handleError(w, models.ErrRequest, err, http.StatusBadRequest)
		return
	}

	var (
		file   multipart.File
		header *multipart.FileHeader
		err    error
	)

	if r.MultipartForm.File == nil || len(r.MultipartForm.File["file"]) == 0 {
		file = nil
		header = nil
	} else {
		file, header, err = r.FormFile("file")
		if err != nil {
			h.handleError(w, models.ErrFileRequest, err, http.StatusBadRequest)
			return
		}
		defer file.Close()
	}

	switch table {
	case "directives":
		//result, err = h.updateDirective(r, file, header, login)
		h.handleError(w, models.ErrFileRequest, err, http.StatusBadRequest)
		return
	case "inventory":
		//result, err = h.updateInventory(r, file, header, login)
		h.handleError(w, models.ErrFileRequest, err, http.StatusBadRequest)
		return
	default:
		err = h.updateInOutDoc(r, file, header, login)
	}

	if err != nil {
		h.handleError(w, err.Error(), err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *DocHandler) addDirective(r *http.Request, file multipart.File, header *multipart.FileHeader, login string) error {
	var doc models.Directive
	if err := json.Unmarshal([]byte(r.FormValue("document")), &doc); err != nil {
		return err
	}
	doc.File = file
	doc.FileHeader = header
	doc.Creator = login
	return h.service.AddDirective(doc)
}

func (h *DocHandler) addInventory(r *http.Request, file multipart.File, header *multipart.FileHeader, login string) error {
	var doc models.Inventory
	if err := json.Unmarshal([]byte(r.FormValue("document")), &doc); err != nil {
		return err
	}
	doc.File = file
	doc.FileHeader = header
	doc.Creator = login
	return h.service.AddInventory(doc)
}

func (h *DocHandler) addInOutDoc(r *http.Request, file multipart.File, header *multipart.FileHeader, login string) error {
	var doc models.Document
	if err := json.Unmarshal([]byte(r.FormValue("document")), &doc); err != nil {
		return err
	}
	doc.File = file
	doc.FileHeader = header
	doc.Creator = login
	return h.service.AddDocument(doc)
}

func (h *DocHandler) updateInOutDoc(r *http.Request, file multipart.File, header *multipart.FileHeader, login string) error {
	id := mux.Vars(r)["id"]
	var doc models.Document
	if err := json.Unmarshal([]byte(r.FormValue("document")), &doc); err != nil {
		return err
	}
	doc.File = file
	doc.FileHeader = header
	doc.Creator = login
	doc.IDStr = id
	return h.service.UpdateDocument(doc)
}

func (h *DocHandler) handleError(w http.ResponseWriter, message string, err error, status int) {
	h.log.Error(message, err)
	http.Error(w, message, status)
}
