package handler

import (
	"documentum/pkg/service"
	"documentum/pkg/models"
	_ "github.com/go-sql-driver/mysql"
	"html/template"
	"log"
	"net/http"
)

type PageHandler struct {
	authService service.AuthService
	pageService service.PageService
}

func NewPagesHandler(authService service.AuthService, pageService service.PageService) *PageHandler {
	return &PageHandler{
		authService: authService,
		pageService: pageService,
	}
}

// Метод для вывода страницы входа/ главной страницы
func (p PageHandler) GetHandler(w http.ResponseWriter, r *http.Request) {

	var requestToken string
	var responseData struct {
		UserIsValid bool
		Login       string
		Funcs []models.Unit
	}

	cookie, err := r.Cookie("token")
	if err != nil {
		responseData.UserIsValid = false
	} else {
		requestToken = cookie.Value
		responseData.Login, err = p.authService.CheckUserTokenToValid(requestToken)
		if err != nil {
			responseData.UserIsValid = false
		} else {
			responseData.UserIsValid = true
		}
	}

	responseData.Funcs, err = p.pageService.GetFuncs()
	
	if err != nil {
		responseData.Funcs = nil
		http.Error(w, err.Error(), 500)
	}

	pages := []string{
		"web/static/pages/global.html",
		"web/static/pages/entrance.html",
		"web/static/pages/main.html",
	}

	ts, err := template.ParseFiles(pages...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	err = ts.Execute(w, responseData)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500) 
		return
	}
}
