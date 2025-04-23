package handler

import (
	"documentum/pkg/models"
	"documentum/pkg/service"
	"encoding/json"
	"net/http"
	"html/template"
	"log"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type AuthHandler struct {
	authService service.AuthService
}

func NewAuthHandler(authService service.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

func (p AuthHandler) RegistrationHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "ошибка обработки данных пользователя", 400)
		return
	}

	err := p.authService.UserRegistration(user)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}

func (p *AuthHandler) AuthorizationHandler(w http.ResponseWriter, r *http.Request) {
    // Получение данных формы
    login := r.FormValue("login")
    pass := r.FormValue("pass")
    remember := r.FormValue("remember")

    // Авторизация пользователя
    userValid, err := p.authService.UserAuthorization(login, pass)
    if err != nil {
        status := http.StatusInternalServerError
        if userValid {
            status = http.StatusUnauthorized
        }
        http.Error(w, err.Error(), status)
        return
    }

    // Генерация и установка токена
    if err := p.authService.GenerateToken(w, login, remember); err != nil {
        http.Error(w, "Ошибка генерации токена: "+err.Error(), http.StatusInternalServerError)
        return
    }

    // Получение данных аккаунта
    responseData, err := p.authService.GetAccountData(login)
    if err != nil {
        log.Printf("Ошибка получения данных пользователя %s: %v", login, err)
        http.Error(w, "Ошибка получения данных о пользователе", http.StatusInternalServerError)
        return
    }

    // Рендеринг страницы
    if err := p.renderMainPage(w, responseData); err != nil {
        log.Printf("Ошибка рендеринга для пользователя %s: %v", login, err)
        http.Error(w, "Ошибка на сервере", http.StatusInternalServerError)
    }
}

func (h *AuthHandler) GetHandler(w http.ResponseWriter, r *http.Request) {
    responseData := struct {
        UserIsValid bool
        Login       string
        Funcs       []models.Unit
    }{}

    // Проверка токена
    cookie, err := r.Cookie("token")
    if err == nil {
        responseData.Login, err = h.authService.CheckUserTokenToValid(cookie.Value)
        responseData.UserIsValid = err == nil
    }

    // Получение функций
    funcs, err := h.authService.GetFuncs()
    if err != nil {
        log.Printf("Ошибка получения функций: %v", err)
    }
    responseData.Funcs = funcs

    // Рендеринг страницы
    if err := h.renderEntrancePage(w, responseData); err != nil {
        log.Printf("Ошибка рендеринга: %v", err)
        http.Error(w, "Ошибка на сервере", http.StatusInternalServerError)
    }
}

// Вспомогательные методы для рендеринга
func (h *AuthHandler) renderMainPage(w http.ResponseWriter, data interface{}) error {
    pages := []string{
        "web/static/pages/main.html",
        "web/static/pages/main_account.html",
        "web/static/pages/main_settings.html",
    }
    return h.renderTemplates(w, "main", pages, data)
}

func (h *AuthHandler) renderEntrancePage(w http.ResponseWriter, data interface{}) error {
    pages := []string{
        "web/static/pages/global.html",
        "web/static/pages/entrance.html",
        "web/static/pages/main.html",
        "web/static/pages/main_account.html",
        "web/static/pages/main_settings.html",
    }
    return h.renderTemplates(w, "", pages, data)
}

func (h *AuthHandler) renderTemplates(w http.ResponseWriter, tmpl string, files []string, data interface{}) error {
    ts, err := template.ParseFiles(files...)
    if err != nil {
        return fmt.Errorf("ошибка парсинга шаблонов: %w", err)
    }

    if tmpl != "" {
        err = ts.ExecuteTemplate(w, tmpl, data)
    } else {
        err = ts.Execute(w, data)
    }
    
    if err != nil {
        return fmt.Errorf("ошибка выполнения шаблона: %w", err)
    }
    return nil
}