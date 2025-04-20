package handler

import (
	"documentum/pkg/service"
	_ "github.com/go-sql-driver/mysql"
	"html/template"
	"log"
	"net/http"
)

type PageHandler struct {
	authService service.AuthService
}

func NewPagesHandler(authService service.AuthService) *PageHandler {
	return &PageHandler{
		authService: authService,
	}
}

// Метод для вывода страницы входа
func (p PageHandler) GetHandler(w http.ResponseWriter, r *http.Request) {

	var requestToken string
	var responseData struct {
		UserIsValid bool
		Login       string
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
