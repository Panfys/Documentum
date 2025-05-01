package main

import (
	"database/sql"
	"log"
	"documentum/pkg/logger"
	"documentum/pkg/config"
	"documentum/pkg/server"
	_ "github.com/go-sql-driver/mysql" 
	
)

func main() {
	logf, err := logger.NewFileLogger("documentum.log")
	if err != nil {
		log.Fatalf("Ошибка создания логгера: %v", err)
	}
	defer logf.Close()

	// Загрузка конфигурации
	cfg := config.LoadConfig() 

	// Подключение к MySQL
	db, err := sql.Open("mysql", cfg.DBConnectionString)
	if err != nil {
		logf.Error("Ошибка подключения к БД: %v", err)
		db = nil
	} else {
		defer db.Close()  
	}

	// Проверка соединения
	if db != nil {
		if err := db.Ping(); err != nil {
			logf.Error("Ошибка проверки соединения с БД: %v", err)
			db = nil
		} else {
			logf.Info("Успешно подключено к БД")
		}
	}

	secretKey := cfg.JWTSecretKey

	// Инициализация и запуск сервера
	srv := server.NewServer(db, secretKey, logf)
	if err := srv.Run(); err != nil {
		logf.Error("Ошибка запуска сервера: %v", err)   
	}
}