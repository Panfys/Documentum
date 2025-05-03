package routes

import (
	"database/sql"
	"documentum/pkg/handler"
	"documentum/pkg/logger"
	"documentum/pkg/service/auth"
	"documentum/pkg/service/document"
	"documentum/pkg/service/file"
	"documentum/pkg/service/structure"
	"documentum/pkg/service/user"
	"documentum/pkg/service/valid"
	"documentum/pkg/storage"

	"net/http"

	"github.com/gorilla/mux"
)

func SetupRoutes(db *sql.DB, secretKey string, log logger.Logger) http.Handler {
	r := mux.NewRouter()

	//Storage
	stor := storage.NewSQLStorage(db, log)

	//Service
	fileService := file.NewFilesService(log)
	validService := valid.NewValidatService()
	userService := user.NewUserService(log, stor, validService, fileService)
	structService := structure.NewstructureService(stor)
	docService := document.NewDocService(stor, validService, fileService)
	authServise := auth.NewAuthService(log, stor, validService, secretKey)

	//Handlers
	authHandler := handler.NewAuthHandler(log, authServise, userService, structService)
	userHandler := handler.NewUserHandler(log, userService)
	docHandler := handler.NewDocHandler(log, docService)
	structHandler := handler.NewStructureHandler(structService)

	r.HandleFunc("/", authHandler.GetHandler)

	// STATIC
	r.PathPrefix("/static/").Handler(handler.StaticHandler())
	r.PathPrefix("/source/").Handler(authHandler.AuthMiddleware(handler.StaticHandler()))

	// OPEN
	r.HandleFunc("/funcs/{id:[0-9]+}/units", structHandler.GetUnits).Methods("GET")
	r.HandleFunc("/funcs/{funcId:[0-9]+}/{unitId:[0-9]+}/groups", structHandler.GetGroups).Methods("GET")
	r.HandleFunc("/auth/register", authHandler.RegistrationHandler).Methods("POST")
	r.HandleFunc("/auth/login", authHandler.AuthorizationHandler).Methods("POST")

	// PROTECTED
	protect := r.PathPrefix("/").Subrouter()
	protect.Use(authHandler.AuthMiddleware)
	//GET
	protect.HandleFunc("/documents", docHandler.GetDocuments).Methods("GET")
	protect.HandleFunc("/document/file", docHandler.WievDocument).Methods("GET")
	protect.HandleFunc("/document/new/file", docHandler.WievNewDocument).Methods("GET")
	//POST
	protect.HandleFunc("/document", docHandler.AddIngoingDoc).Methods("POST")
	//PATCH
	protect.HandleFunc("/user/me/icon", userHandler.UpdateUserIcon).Methods("PATCH")
	protect.HandleFunc("/user/me/pass", userHandler.UpdateUserPassword).Methods("PATCH")
	protect.HandleFunc("/document/{id:[0-9]+}/view", docHandler.AddLookDocument).Methods("PATCH")
	//DELETE
	protect.HandleFunc("/auth/logout", authHandler.ExitHandler).Methods("DELETE")

	// ERROR 404
	r.NotFoundHandler = http.HandlerFunc(handler.NotFoundHandler)

	return r
}
