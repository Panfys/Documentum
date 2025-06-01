package logger

import (
	"fmt"
	"log"
	"os"
	"time"
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
	file, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}

	return &FileLogger{
		logger: log.New(file, "", 0),
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
