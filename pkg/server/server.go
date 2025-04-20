package server

import (
	"log"
	"net/http"
	"documentum/pkg/routes"
	"database/sql"

	"github.com/gorilla/handlers"
	"os"
)

type Server struct {
	httpServer *http.Server
}

func NewServer(db *sql.DB) *Server {
	router := routes.SetupRoutes(db)
	
	// Создаем цепочку middleware
	handler := handlers.CompressHandler(router) // Gzip сжатие
	handler = handlers.RecoveryHandler()(handler) // Обработка паник
	handler = handlers.LoggingHandler(os.Stdout, handler) // Логирование
	handler = handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)(handler) // CORS политика

	return &Server{
		httpServer: &http.Server{ 
			Addr:    ":80",
			Handler: handler, // Используем router напрямую
		},
	}
}

func (s *Server) Run() error {
	log.Println("Запуск сервера на http://localhost:80") 
	return s.httpServer.ListenAndServe()
}