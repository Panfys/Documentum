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

	docStorage := storage.NewDocStorage(db)
	userStoreg := storage.NewUserStorage(db)
	authStoreg := storage.NewAuthStorage(db)
	userService := service.NewUserService(userStoreg)
	docService := service.NewDocService(docStorage)
	authServise := service.NewAuthService(authStoreg, secretKey)
	authHandler := handler.NewAuthHandler(authServise)
	userHandler := handler.NewUserHandler(userService)
	docHandler := handler.NewDocHandler(docService)

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
	protect.HandleFunc("/documents/getdoc", docHandler.GetDocuments).Methods("POST")
	protect.HandleFunc("/documents/ingoing", docHandler.GetIngoingDoc).Methods("POST")
	protect.HandleFunc("/documents/wievdoc", docHandler.WievDocument)
	protect.HandleFunc("/documents/wievnewdoc", docHandler.WievNewDocument)
	protect.HandleFunc("/documents/look", docHandler.AddLookDocument).Methods("POST")

	// ERROR 404
	r.NotFoundHandler = http.HandlerFunc(handler.NotFoundHandler)

	// STATIC
	r.PathPrefix("/static/").Handler(handler.StaticHandler())
	r.PathPrefix("/source/").Handler(authHandler.AuthMiddleware(handler.StaticHandler()))

	return r
}
