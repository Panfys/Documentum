package handler

import (
	_ "github.com/go-sql-driver/mysql"
	"html/template"
	"log"
	"net/http"
	//"documentum/pkg/service"
)

type PageHandler struct {
	
}

func NewPagesHandler() *PageHandler {
	return &PageHandler{}
}

// Метод для вывода страницы входа
func (p PageHandler) GetHandler(w http.ResponseWriter, r *http.Request) {

	files := []string{
		"web/static/pages/global.html",
		"web/static/pages/entrance.html",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		//log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}
}