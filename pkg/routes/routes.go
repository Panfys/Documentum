package routes

import (
	"database/sql"
	"documentum/pkg/handler"
	//"documentum/pkg/service"
	//"documentum/pkg/storage"
	"net/http"

	"github.com/gorilla/mux"
)

func SetupRoutes(db *sql.DB) http.Handler {
	r := mux.NewRouter()

	//userStoreg := storage.NewUserStorage(db)
	//userService := service.NewUserService(userStoreg)
	//userHandler := handler.NewUserHandler(userService)

	r.HandleFunc("/", handler.NewPagesHandler().GetHandler)

	// ERROR 404
	r.NotFoundHandler = http.HandlerFunc(handler.NotFoundHandler)

	staticDir := "web/static"

	// Регистрация обработчика
	r.PathPrefix("/static/").Handler(handler.StaticHandler(staticDir))

	return r
}
