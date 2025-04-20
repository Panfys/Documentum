package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBConnectionString string
}

func LoadConfig() *Config {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Ошибка загрузки файла конфигураций")
	}

	return &Config{
		DBConnectionString: os.Getenv("DB_CONNECTION_STRING"),  
	}
}