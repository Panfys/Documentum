package handler

import (
	"documentum/pkg/logger"
	"documentum/pkg/models"
	"documentum/pkg/service/document"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
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

	var settings models.DocSettings
	settings.DocType = r.FormValue("type")
	settings.DocCol = r.FormValue("col")
	settings.DocSet = r.FormValue("set")
	settings.DocDatain = r.FormValue("datain")
	settings.DocDatato = r.FormValue("datato")

	switch settings.DocType {

	case "Входящий":
		responceDocs, err := h.service.GetIngoingDoc(settings)

		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(responceDocs))

	default:
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(settings.DocType))
	}
}

func (h *DocHandler) WievDocument(w http.ResponseWriter, r *http.Request) {
	file := r.URL.Query().Get("file")

	// Безопасная проверка пути к файлу
	filePath := filepath.Join("/app/web", filepath.Clean(file))

	var SRC template.HTML
	if _, err := os.Stat(filePath); err == nil {
		// Проверяем расширение файла
		if filepath.Ext(file) == ".pdf" {
			SRC = template.HTML(fmt.Sprintf("<object><embed src='%s'></embed></object>", file))
		} else {
			SRC = template.HTML(fmt.Sprintf("<img src='%s'>", file))
		}
	} else if os.IsNotExist(err) {
		SRC = template.HTML("Файл не найден!")
	} else {
		SRC = template.HTML("Файл недоступен!")
	}

	data := models.PageData{SRC: SRC}

	h.renderTemplates(w, data)
}

func (h *DocHandler) WievNewDocument(w http.ResponseWriter, r *http.Request) {
	file := r.URL.Query().Get("file")

	var SRC template.HTML

	if filepath.Ext(file) == ".pdf" {
		SRC = template.HTML(fmt.Sprintf("<object><embed src='%s'></embed></object>", file))
	} else {
		SRC = template.HTML(file)
	}

	data := models.PageData{SRC: SRC}

	h.renderTemplates(w, data)
}

func (h *DocHandler) renderTemplates(w http.ResponseWriter, data models.PageData) error {
	ts, err := template.ParseFiles("web/static/pages/main_open_doc.html")
	if err != nil {
		return h.log.Error(models.ErrParseTMP, err)
	}
	err = ts.Execute(w, data)

	if err != nil {
		return h.log.Error(models.ErrParseTMP, err)
	}
	return nil
}

func (h *DocHandler) AddLookDocument(w http.ResponseWriter, r *http.Request) {

	login := r.Context().Value(models.LoginKey).(string)
	idStr := r.FormValue("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, models.ErrRequest, http.StatusBadRequest)
		h.log.Error(models.ErrRequest, err)
		return
	}

	err = h.service.AddLookDocument(id, login)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *DocHandler) AddIngoingDoc(w http.ResponseWriter, r *http.Request) {

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
}
