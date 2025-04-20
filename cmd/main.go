package main

import (
	"database/sql"
	"log"
	"documentum/pkg/config"
	"documentum/pkg/server"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Загрузка конфигурации
	cfg := config.LoadConfig()

	// Подключение к MySQL
	db, err := sql.Open("mysql", cfg.DBConnectionString)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Проверка соединения
	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	// Инициализация и запуск сервера
	srv := server.NewServer(db)
	if err := srv.Run(); err != nil {
		log.Fatalf("Server failed: %v", err) 
	}
}