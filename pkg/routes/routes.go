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
	fileService := file.NewFilesService()
	validService := valid.NewValidatService()
	userService := user.NewUserService(stor, validService, fileService)
	structService := structure.NewstructureService(stor)
	docService := document.NewDocService(stor, validService, fileService)
	authServise := auth.NewAuthService(log, stor, validService, secretKey)

	//Handlers
	authHandler := handler.NewAuthHandler(log, authServise, userService, structService)
	userHandler := handler.NewUserHandler(userService)
	docHandler := handler.NewDocHandler(docService)
	structHandler := handler.NewStructureHandler(structService)

	r.HandleFunc("/", authHandler.GetHandler)

	// OPEN
	r.HandleFunc("/users/units", structHandler.GetUnits).Methods("POST")
	r.HandleFunc("/users/groups", structHandler.GetGroups).Methods("POST")
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
	protect.HandleFunc("/documents/add", docHandler.AddIngoingDoc).Methods("POST")

	// ERROR 404
	r.NotFoundHandler = http.HandlerFunc(handler.NotFoundHandler)

	// STATIC
	r.PathPrefix("/static/").Handler(handler.StaticHandler())
	r.PathPrefix("/source/").Handler(authHandler.AuthMiddleware(handler.StaticHandler()))

	return r
}
