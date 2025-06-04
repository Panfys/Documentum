package server

import (
	"database/sql"
	"documentum/pkg/logger"
	"documentum/pkg/routes"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
)

type Server struct {
	httpServer *http.Server
	log        logger.Logger // Используем интерфейс вместо конкретного типа
}

// NewServer создает новый экземпляр сервера с заданной конфигурацией.
func NewServer(db *sql.DB, secretKey string, log logger.Logger) *Server {
	router := routes.SetupRoutes(db, secretKey, log)

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
			Addr:    ":8000",
			Handler: handler,
		},
		log: log,
	}
}

// Run запускает HTTP сервер и логирует ошибки.
func (s *Server) Run() error {
	s.log.Info("Запуск сервера на http://localhost:8000")
	if err := s.httpServer.ListenAndServe(); err != nil {
		s.log.Error("Ошибка при запуске сервера: %v", err)
		return err
	}
	return nil
}