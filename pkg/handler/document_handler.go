package handler

import (
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
	service document.DocService
}

func NewDocHandler(service document.DocService) *DocHandler {
	return &DocHandler{service: service}
}

func (d *DocHandler) GetIngoingDoc(w http.ResponseWriter, r *http.Request) {

	var settings models.DocSettings
	settings.DocType = "Входящий"
	settings.DocCol = r.FormValue("col")
	settings.DocSet = r.FormValue("set")
	settings.DocDatain = r.FormValue("datain")
	settings.DocDatato = r.FormValue("datato")

	responceDocs, err := d.service.GetIngoingDoc(settings)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(responceDocs))
}

func (d *DocHandler) GetDocuments(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Документ"))
}

func (d *DocHandler) WievDocument(w http.ResponseWriter, r *http.Request) {
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
		SRC = template.HTML("Файл не найден.")
	} else {
		SRC = template.HTML("Ошибка при доступе к файлу: " + err.Error())
	}

	data := models.PageData{SRC: SRC}

	d.renderTemplates(w, data)
}

func (d *DocHandler) WievNewDocument(w http.ResponseWriter, r *http.Request) {
	file := r.URL.Query().Get("file")

	var SRC template.HTML

	if filepath.Ext(file) == ".pdf" {
		SRC = template.HTML(fmt.Sprintf("<object><embed src='%s'></embed></object>", file))
	} else {
		SRC = template.HTML(file)
	}

	data := models.PageData{SRC: SRC}

	d.renderTemplates(w, data)
}

func (d *DocHandler) renderTemplates(w http.ResponseWriter, data models.PageData) error {
	ts, err := template.ParseFiles("web/static/pages/main_open_doc.html")
	if err != nil {
		return fmt.Errorf("ошибка парсинга шаблонов: %w", err)
	}
	err = ts.Execute(w, data)

	if err != nil {
		return fmt.Errorf("ошибка выполнения шаблона: %w", err)
	}
	return nil
}

func (d *DocHandler) AddLookDocument(w http.ResponseWriter, r *http.Request) {

	login := r.Context().Value(models.LoginKey).(string)
	idStr := r.FormValue("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, "Ошибка обработки данных!:", http.StatusBadRequest)
		return
	}

	err = d.service.AddLookDocument(id, login)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (d *DocHandler) AddIngoingDoc(w http.ResponseWriter, r *http.Request) {

	reqError := "ошибка обработки данных"

	login := r.Context().Value(models.LoginKey).(string)

	countStr := r.FormValue("count")

	count, err := strconv.Atoi(countStr)
	if err != nil {
		http.Error(w, reqError + err.Error(), http.StatusBadRequest)
		return
	}

	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, reqError + err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	document := models.Document{
		Type:       r.FormValue("type"),
		FNum:       r.FormValue("fnum"),
		FDate:      r.FormValue("fdate"),
		LNum:       r.FormValue("lnum"),
		LDateStr:      r.FormValue("ldate"),
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
		http.Error(w, reqError + err.Error(), http.StatusBadRequest)
		return
	}

	doc, err := d.service.AddIngoingDoc(document)
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
