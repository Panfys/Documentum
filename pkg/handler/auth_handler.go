package handler

import (
	"context"
	"documentum/pkg/models"
	"documentum/pkg/service/auth"
	"documentum/pkg/service/user"
	"documentum/pkg/service/structure"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var pages = []string{
	"web/static/pages/global.html",
	"web/static/pages/entrance.html",
	"web/static/pages/main.html",
	"web/static/pages/main_account.html",
	"web/static/pages/main_settings.html",
	"web/static/pages/main_ingoing_doc.html",
	"web/static/pages/main_outgoing_doc.html",
	"web/static/pages/main_inventory_doc.html",
	"web/static/pages/main_directive_doc.html",
}

type AuthHandler struct {
	authSrv auth.AuthService
	userSrv	user.UserService
	structSrv structure.StructureService
}

func NewAuthHandler(authSrv auth.AuthService, userSrv	user.UserService, structSrv structure.StructureService) *AuthHandler {
	return &AuthHandler{
		authSrv: authSrv,
		userSrv: userSrv,
		structSrv: structSrv,
	}
}

func (p AuthHandler) RegistrationHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "ошибка обработки данных пользователя", 400)
		return
	}

	err := p.authSrv.UserRegistration(user)

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
	status, err := p.authSrv.UserAuthorization(login, pass)
	if err != nil {
		http.Error(w, err.Error(), status)
		return
	}

	// Генерация и установка токена
	if err := p.authSrv.GenerateToken(w, login, remember); err != nil {
		http.Error(w, "Ошибка генерации токена: "+err.Error(), http.StatusInternalServerError)
		return
	}

	responseData := struct {
		AccountData models.AccountData
	}{}

	// Получение данных аккаунта
	responseData.AccountData, err = p.userSrv.GetUserAccountData(login)
	if err != nil {
		log.Printf("Ошибка получения данных пользователя %s: %v", login, err)
		http.Error(w, "Ошибка получения данных о пользователе", http.StatusInternalServerError)
		return
	}

	// Рендеринг страницы
	if err := p.renderTemplates(w, "main", responseData); err != nil {
		log.Printf("Ошибка рендеринга для пользователя %s: %v", login, err)
		http.Error(w, "Ошибка на сервере", http.StatusInternalServerError)
	}
}

func (h *AuthHandler) GetHandler(w http.ResponseWriter, r *http.Request) {
	responseData := struct {
		UserIsValid bool
		AccountData models.AccountData
		Funcs       []models.Unit
	}{}

	// Проверка токена
	cookie, err := r.Cookie("token")
	if err == nil {
		login, err := h.authSrv.CheckUserTokenToValid(cookie.Value)
		responseData.UserIsValid = err == nil
		responseData.AccountData, err = h.userSrv.GetUserAccountData(login)
		if err != nil {
			log.Printf("Ошибка получения данных пользователя %s: %v", login, err)
			http.Error(w, "Ошибка получения данных о пользователе", http.StatusInternalServerError)
			return
		}
	}

	// Получение функций
	funcs, err := h.structSrv.GetFuncs()
	if err != nil {
		log.Printf("Ошибка получения должности: %v", err)
	}
	responseData.Funcs = funcs

	// Рендеринг страницы
	if err := h.renderTemplates(w, "", responseData); err != nil {
		log.Printf("Ошибка рендеринга: %v", err)
		http.Error(w, "Ошибка на сервере", http.StatusInternalServerError)
	}
}

func (h *AuthHandler) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Получаем токен из cookie
		cookie, err := r.Cookie("token")
		if err != nil {
			if r.Method == http.MethodGet {
				http.Redirect(w, r, "/", http.StatusSeeOther)
				return
			} else {
				http.Error(w, "Пользователь не авторизован", http.StatusUnauthorized)
				return
			}
		}

		// Проверяем валидность токена
		login, err := h.authSrv.CheckUserTokenToValid(cookie.Value)
		if err != nil {
			// Удаляем невалидный токен
			http.SetCookie(w, &http.Cookie{
				Name:     "token",
				Value:    "",
				Path:     "/",
				Expires:  time.Now().Add(-1 * time.Hour),
				HttpOnly: true,
				Secure:   false, // В продакшене должно быть true
			})
			if r.Method == http.MethodGet {
				http.Redirect(w, r, "/", http.StatusSeeOther)
				return
			} else {
				http.Error(w, "Пользователь не авторизован", http.StatusUnauthorized)
				return
			}
		}

		ctx := context.WithValue(r.Context(), models.LoginKey, login)
		
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (h *AuthHandler) ExitHandler(w http.ResponseWriter, r *http.Request) {

	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    "",                             // очищаем значение
		Path:     "/",                            // действует для всех путей
		HttpOnly: true,                           // защита от XSS
		Secure:   false,                          // только HTTPS (false для localhost)
		Expires:  time.Now().Add(-1 * time.Hour), // срок истёк
		MaxAge:   -1,                             // удалить cookie немедленно
		SameSite: http.SameSiteStrictMode,        // защита от CSRF
	})

	responseData := struct {
		UserIsValid bool
		AccountData models.AccountData
		Funcs       []models.Unit
	}{
		UserIsValid: false,
	}

	// Получение функций
	funcs, err := h.structSrv.GetFuncs()
	if err != nil {
		log.Printf("Ошибка получения должности: %v", err)
	}
	responseData.Funcs = funcs

	// Рендеринг страницы
	if err := h.renderTemplates(w, "", responseData); err != nil {
		log.Printf("Ошибка рендеринга: %v", err)
		http.Error(w, "Ошибка на сервере", http.StatusInternalServerError)
	}
}

func (h *AuthHandler) renderTemplates(w http.ResponseWriter, tmpl string, data any) error {
	ts, err := template.ParseFiles(pages...)
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
