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
	authStoreg := storage.NewAuthStorage(db)
	authServise := service.NewAuthService(authStoreg, secretKey)
	authHandler := handler.NewAuthHandler(authServise)


	r.HandleFunc("/", authHandler.GetHandler)
	r.HandleFunc("/users/units", userHandler.GetUnits)
	r.HandleFunc("/users/groups", userHandler.GetGroups)
	r.HandleFunc("/users/add", authHandler.RegistrationHandler)
	r.HandleFunc("/users/auth", authHandler.AuthorizationHandler)

	// ERROR 404
	r.NotFoundHandler = http.HandlerFunc(handler.NotFoundHandler)

	staticDir := "web/static"

	// Регистрация обработчика
	r.PathPrefix("/static/").Handler(handler.StaticHandler(staticDir)) 

	return r
}
