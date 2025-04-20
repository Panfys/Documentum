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

	//userStoreg := storage.NewUserStorage(db)
	//userService := service.NewUserService(userStoreg)
	//userHandler := handler.NewUserHandler(userService)
	   
	authStoreg := storage.NewAuthStorage(db)
	authServise := service.NewAuthService(authStoreg, secretKey)
	authHandler := handler.NewPagesHandler(authServise)

	r.HandleFunc("/", authHandler.GetHandler)

	// ERROR 404
	r.NotFoundHandler = http.HandlerFunc(handler.NotFoundHandler)

	staticDir := "web/static"

	// Регистрация обработчика
	r.PathPrefix("/static/").Handler(handler.StaticHandler(staticDir))

	return r
}
