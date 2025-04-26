package handler

import (
	"documentum/pkg/models"
	"documentum/pkg/service"
	"net/http"
	"html/template"
	"os"
	"path/filepath"
	"fmt"
	"strconv"
)

type DocHandler struct {
	docService service.DocService
}

func NewDocHandler(docService service.DocService) *DocHandler {
	return &DocHandler{docService: docService}
}

func (d *DocHandler) GetIngoingDoc(w http.ResponseWriter, r *http.Request) {
	
	var settings models.DocSettings
	settings.DocType = "Входящий"
	settings.DocCol = r.FormValue("col")
	settings.DocSet = r.FormValue("set")
	settings.DocDatain = r.FormValue("datain")
	settings.DocDatato = r.FormValue("datato")
	
	responceDocs, err := d.docService.GetIngoingDoc(settings)
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

// PageData структура для передачи данных в шаблон
type PageData struct {
    SRC template.HTML
}

// WievDocumentHandler метод для просмотра документа
func (d *DocHandler) WievDocumentHandler(w http.ResponseWriter, r *http.Request) {
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

    // Парсим HTML-шаблон
    ts, err := template.ParseFiles("web/static/pages/main_open_doc.html")
    if err != nil {
        http.Error(w, "Ошибка сервера", http.StatusInternalServerError)
        return
    }

    // Выполняем шаблон с данными
    data := PageData{SRC: SRC}
    err = ts.Execute(w, data)
    if err != nil {
        http.Error(w, "Ошибка сервера", http.StatusInternalServerError)
        return
    }
}

func (d *DocHandler) AddLookDocument(w http.ResponseWriter, r *http.Request) {
	
	login := r.Context().Value(models.LoginKey).(string);
	idStr := r.FormValue("id")
	id, err := strconv.Atoi(idStr)
	
	if err != nil {
		http.Error(w, "Ошибка обработки данных!:", http.StatusBadRequest)
		return
	}

	err = d.docService.AddLookDocument(id, login)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.WriteHeader(http.StatusOK)
}