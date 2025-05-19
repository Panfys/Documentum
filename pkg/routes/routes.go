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
	"documentum/pkg/service/ws"
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
	validService := valid.NewValidatService(log)
	wsService := ws.NewWebSocketService(log)
	userService := user.NewUserService(log, stor, validService, fileService)
	structService := structure.NewstructureService(stor)
	docService := document.NewDocService(stor, validService, fileService, wsService)
	authServise := auth.NewAuthService(log, stor, validService, secretKey)
	

	//Handlers
	authHandler := handler.NewAuthHandler(log, authServise, userService, structService)
	userHandler := handler.NewUserHandler(log, userService)
	docHandler := handler.NewDocHandler(log, docService)
	structHandler := handler.NewStructureHandler(structService)
	wsHandler := handler.NewWebSocketHandler(log, wsService)

	/// OPEN
	//GET
	r.HandleFunc("/", authHandler.GetHandler).Methods("GET")
	r.PathPrefix("/static/").Handler(handler.StaticHandler()).Methods("GET")
	r.PathPrefix("/source/").Handler(authHandler.AuthMiddleware(handler.StaticHandler())).Methods("GET")
	r.HandleFunc("/structures/{funcId:[0-9]+}", structHandler.GetUnits).Methods("GET")
	r.HandleFunc("/structures/{funcId:[0-9]+}/{unitId:[0-9]+}", structHandler.GetGroups).Methods("GET")
	//POST
	r.HandleFunc("/auth/register", authHandler.RegistrationHandler).Methods("POST")
	r.HandleFunc("/auth/login", authHandler.AuthorizationHandler).Methods("POST")

	/// PROTECTED
	protect := r.PathPrefix("/").Subrouter()
	protect.Use(authHandler.AuthMiddleware)
	//GET
	protect.HandleFunc("/ws", wsHandler.HandleConnection).Methods("GET")
	protect.HandleFunc("/documents/{table:[a-z]+}", docHandler.GetDocuments).Methods("GET")
	//POST
	protect.HandleFunc("/documents/{table:[a-z]+}", docHandler.AddDocument).Methods("POST")
	//PATCH
	protect.HandleFunc("/users/me/icon", userHandler.UpdateUserIcon).Methods("PATCH")
	protect.HandleFunc("/users/me/pass", userHandler.UpdateUserPassword).Methods("PATCH")
	protect.HandleFunc("/documents/{type:[a-z]+}/{id:[0-9]+}/familiar", docHandler.UpdateDocFamiliar).Methods("PATCH")
	protect.HandleFunc("/documents/{table:[a-z]+}/{id:[0-9]+}", docHandler.UpdateDocResolutions).Methods("PATCH")
	//DELETE
	protect.HandleFunc("/auth/logout", authHandler.ExitHandler)

	// ERROR 404
	r.NotFoundHandler = http.HandlerFunc(handler.NotFoundHandler)

	return r
}
