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
		log.Fatalf("Ошибка подключения к БД: %v", err)
	}
	defer db.Close()

	// Проверка соединения
	if err := db.Ping(); err != nil {
		log.Fatalf("Ошибка проверки соединения с БД: %v", err)
	}


	secretKey := cfg.JWTSecretKey

	// Инициализация и запуск сервера
	srv := server.NewServer(db, secretKey)
	if err := srv.Run(); err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)  
	}
}