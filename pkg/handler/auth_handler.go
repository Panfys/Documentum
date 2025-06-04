package handler

import (
	"context"
	"documentum/pkg/logger"
	"documentum/pkg/models"
	"documentum/pkg/service/auth"
	"documentum/pkg/service/structure"
	"documentum/pkg/service/user"
	"encoding/json"
	"html/template"
	"net/http"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var pages = []string{
	"app/web/static/pages/global.html",
	"app/web/static/pages/entrance.html",
	"app/web/static/pages/main.html",
	"app/web/static/pages/main_account.html",
	"app/web/static/pages/main_settings.html",
	"app/web/static/pages/main_ingoing_doc.html",
	"app/web/static/pages/main_outgoing_doc.html",
	"app/web/static/pages/main_inventory_doc.html",
	"app/web/static/pages/main_directive_doc.html",
}

type AuthHandler struct {
	log       logger.Logger
	authSrv   auth.AuthService
	userSrv   user.UserService
	structSrv structure.StructureService
}

func NewAuthHandler(log logger.Logger, authSrv auth.AuthService, userSrv user.UserService, structSrv structure.StructureService) *AuthHandler {
	return &AuthHandler{
		log:       log,
		authSrv:   authSrv,
		userSrv:   userSrv,
		structSrv: structSrv,
	}
}

func (h AuthHandler) RegistrationHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		h.log.Error(models.ErrRequest, err)
		http.Error(w, models.ErrRequest, 400)
		return
	}

	err := h.authSrv.UserRegistration(user)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *AuthHandler) AuthorizationHandler(w http.ResponseWriter, r *http.Request) {
	// Получение данных формы
	var authData models.AuthData

	if err := json.NewDecoder(r.Body).Decode(&authData); err != nil {
		h.log.Error(models.ErrRequest, err)
		http.Error(w, models.ErrRequest, 400)
		return
	}

	authData.Agent = r.UserAgent()
	authData.IP = h.getUserIP(r)

	// Авторизация пользователя
	status, err := h.authSrv.UserAuthorization(authData.Login, authData.Pass)
	if err != nil {
		http.Error(w, err.Error(), status)
		return
	}

	// Генерация и установка токена
	if err := h.authSrv.GenerateToken(w, authData); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	responseData := struct {
		AccountData models.AccountData
	}{}

	// Получение данных аккаунта
	responseData.AccountData, err = h.userSrv.GetUserAccountData(authData.Login)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Рендеринг страницы
	if err := h.renderTemplates(w, "main", responseData); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *AuthHandler) GetHandler(w http.ResponseWriter, r *http.Request) {
	responseData := struct {
		UserIsValid bool
		AccountData models.AccountData
		Funcs       []models.Unit
	}{}

	var authData models.AuthData
	authData.Agent = r.UserAgent()
	authData.IP = h.getUserIP(r)

	// Проверка токена
	cookie, err := r.Cookie("token")
	if err == nil {
		login, err := h.authSrv.CheckUserTokenToValid(cookie.Value, authData.Agent, authData.IP)
		if err == nil {
			responseData.AccountData, err = h.userSrv.GetUserAccountData(login)
			if err != nil {
				h.ExitHandler(w, r)
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			responseData.UserIsValid = err == nil
		} else {
			responseData.AccountData = models.AccountData{}
			h.ExitHandler(w, r)
		}
	}

	// Получение функций
	funcs, err := h.structSrv.GetFuncs()
	if err != nil {
		funcs = []models.Unit{models.Unit{ID: 1, Name: err.Error()}}
	}
	responseData.Funcs = funcs

	// Рендеринг страницы
	if err := h.renderTemplates(w, "", responseData); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
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

		var authData models.AuthData
		authData.Agent = r.UserAgent()
		authData.IP = h.getUserIP(r)

		// Проверяем валидность токена
		login, err := h.authSrv.CheckUserTokenToValid(cookie.Value, authData.Agent, authData.IP)
		if err != nil {
			h.ExitHandler(w, r)
			if r.Method == http.MethodGet {
				http.Redirect(w, r, "/", http.StatusSeeOther)
				return
			} else {
				http.Error(w, "пользователь не авторизован", http.StatusUnauthorized)
				return
			}
		}

		ctx := r.Context()
        ctx = context.WithValue(ctx, models.LoginKey, login)
        ctx = context.WithValue(ctx, models.UserAgentKey, authData.Agent)
        ctx = context.WithValue(ctx, models.IPKey, authData.IP)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (h *AuthHandler) ExitHandler(w http.ResponseWriter, r *http.Request) {

	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
		Expires:  time.Now().Add(-1 * time.Hour),
		MaxAge:   -1,
		SameSite: http.SameSiteStrictMode,
	})
}

func (h *AuthHandler) renderTemplates(w http.ResponseWriter, tmpl string, data any) error {
	ts, err := template.ParseFiles(pages...)
	if err != nil {
		return h.log.Error(models.ErrParseTMP, err)
	}

	if tmpl != "" {
		err = ts.ExecuteTemplate(w, tmpl, data)
	} else {
		err = ts.Execute(w, data)
	}

	if err != nil {
		return h.log.Error(models.ErrParseTMP, err)
	}
	return nil
}

func (h *AuthHandler) getUserIP(r *http.Request) string {
	// Проверяем заголовки, которые могут содержать реальный IP за прокси
	if ip := r.Header.Get("X-Forwarded-For"); ip != "" {
		return strings.Split(ip, ",")[0]
	}
	if ip := r.Header.Get("X-Real-IP"); ip != "" {
		return ip
	}
	// Если заголовков нет, берем RemoteAddr
	return strings.Split(r.RemoteAddr, ":")[0]
}
