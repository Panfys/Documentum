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

	// OPEN
	r.HandleFunc("/users/units", userHandler.GetUnits).Methods("POST")
	r.HandleFunc("/users/groups", userHandler.GetGroups).Methods("POST")
	r.HandleFunc("/users/add", authHandler.RegistrationHandler).Methods("POST")
	r.HandleFunc("/users/auth", authHandler.AuthorizationHandler).Methods("POST")

	// PROTECTED
	protect := r.PathPrefix("/protect/").Subrouter()
	protect.Use(authHandler.AuthMiddleware)

	protect.HandleFunc("/users/exit", authHandler.ExitHandler).Methods("POST")
	protect.HandleFunc("/users/updatepass", userHandler.UpdateUserPassword).Methods("POST")
	protect.HandleFunc("/users/updateicon", userHandler.UpdateUserIcon).Methods("POST")

	// ERROR 404
	r.NotFoundHandler = http.HandlerFunc(handler.NotFoundHandler)

	// STATIC
	r.PathPrefix("/static/").Handler(handler.StaticHandler())
	r.PathPrefix("/source/").Handler(authHandler.AuthMiddleware(handler.StaticHandler()))

	return r
}
