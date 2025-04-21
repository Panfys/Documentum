package routes

import (
	"database/sql"
	"documentum/pkg/handler"
	"documentum/pkg/service"
	"documentum/pkg/storage"
	"net/http"

	"github.com/gorilla/mux"
)

func SetupRoutes(db *sql.DB, secretKey string) http.Handler {
	r := mux.NewRouter()

	userStoreg := storage.NewUserStorage(db)
	userService := service.NewUserService(userStoreg)
	userHandler := handler.NewUserHandler(userService)
	pageStoreg := storage.NewPageStorage(db)
	pageServise := service.NewPageService(pageStoreg)
	authStoreg := storage.NewAuthStorage(db)
	authServise := service.NewAuthService(authStoreg, secretKey)
	pageHandler := handler.NewPagesHandler(authServise, pageServise)
	authHandler := handler.NewAuthHandler(authServise)


	r.HandleFunc("/", pageHandler.GetHandler)
	r.HandleFunc("/users/units", userHandler.GetUnits)
	r.HandleFunc("/users/groups", userHandler.GetGroups)
	r.HandleFunc("/users/add", authHandler.RegistrationHandler)

	// ERROR 404
	r.NotFoundHandler = http.HandlerFunc(handler.NotFoundHandler)

	staticDir := "web/static"

	// Регистрация обработчика
	r.PathPrefix("/static/").Handler(handler.StaticHandler(staticDir)) 

	return r
}
