package logger

import (
	"fmt"
	"log"
	"os"
	"time"
	"path/filepath"
)

type Logger interface {
	Info(msg string)
	Error(format string, args ...any) error
}

type FileLogger struct {
	logger *log.Logger
	file   *os.File
}

func NewFileLogger(logFilePath string) (*FileLogger, error) {
    // Создаем все родительские директории
    if err := os.MkdirAll(filepath.Dir(logFilePath), 0755); err != nil {
        return nil, fmt.Errorf("ошибка создания лог дирректории: %v", err)
    }
    
    file, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        return nil, fmt.Errorf("ошибка открытия лог файла: %v", err)
    }

    return &FileLogger{
        logger: log.New(file, "", log.LstdFlags),
        file:   file,
    }, nil
}

func (l *FileLogger) Info(msg string) {
	l.logger.Printf("%s [INFO] %s\n", time.Now().Format("2006-01-02 15:04:05"), msg)
	log.Printf("%s [INFO] %s\n", time.Now().Format("2006-01-02 15:04:05"), msg)
}

func (l *FileLogger) Error(format string, args ...any) error {
	msg := format + ": " + fmt.Sprint(args...)

	l.logger.Printf("%s [ERROR] %s\n", time.Now().Format("2006-01-02 15:04:05"), msg)
	log.Printf("%s [ERROR] %s\n", time.Now().Format("2006-01-02 15:04:05"), msg)

	return fmt.Errorf("%s", format)
}

func (l *FileLogger) Close() error {
	return l.file.Close()
}
