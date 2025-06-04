package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBConnectionString     string
	DBRootConnectionString string
	JWTSecretKey           string
}

func LoadConfig() *Config {

	err := godotenv.Load("app/.env") // Явный путь к файлу
	if err != nil {
		err := godotenv.Load(".env")
		if err != nil {
			log.Printf("Не удалось загрузить .env файл:", err)
		}
	}

	return &Config{
		DBConnectionString:     os.Getenv("DB_CONNECTION_STRING"),
		DBRootConnectionString: os.Getenv("DB_ROOT_CONNECTION_STRING"),
		JWTSecretKey:           os.Getenv("JWT_SECRET_KEY"),
	}
}
