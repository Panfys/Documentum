package handler

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// Страница не найдена
func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		http.Redirect(w, r, "/", http.StatusFound)
	case "POST":
		http.Error(w, "404 - Not Found (POST)", http.StatusNotFound)
	case "PUT":
		http.Error(w, "404 - Not Found (PUT)", http.StatusNotFound)
	case "DELETE":
		http.Error(w, "404 - Not Found (DELETE)", http.StatusNotFound)
	default:
		http.Error(w, "404 - Not Found (Other)", http.StatusNotFound)
	}
}

// CacheDuration - время кеширования статических файлов (1 год)
const CacheDuration = 365 * 24 * time.Hour

func StaticHandler() http.Handler {
	// Базовый путь к папке web
	webDir := "/app/web"
	absPath, err := filepath.Abs(webDir)
	if err != nil {

		log.Fatalf("Ошибка получения абсолютного пути: %v", err)
	}

	// Создаем файловый сервер для базовой директории
	fs := http.FileServer(http.Dir(absPath))

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Разрешаем доступ только к static/ и source/
		if !strings.HasPrefix(r.URL.Path, "/static/") && !strings.HasPrefix(r.URL.Path, "/source/") {
			NotFoundHandler(w, r)
			return
		}

		// Блокируем доступ к скрытым файлам
		if strings.HasPrefix(filepath.Base(r.URL.Path), ".") {
			http.Error(w, "404 - Not Found", http.StatusNotFound)
			return
		}
		// Полный путь к запрашиваемому файлу
		reqPath := filepath.Join(absPath, filepath.Clean(r.URL.Path))
		if strings.HasSuffix(r.URL.Path, "/") && !strings.HasSuffix(reqPath, "/") {
			reqPath += "/"
		}
		// Проверяем существование файла
		info, err := os.Stat(reqPath)
		if err != nil {
			if os.IsNotExist(err) {
				log.Printf("Файл/директория не найдены: %s", reqPath)
				if strings.HasSuffix(r.URL.Path, "/") {
					NotFoundHandler(w, r)
				} else {
                   http.Error(w, "404 - Not Found", http.StatusNotFound)
                }
			} else {
				log.Printf("Ошибка доступа: %s, error: %v", reqPath, err)
				http.Error(w, "404 - Not Found", http.StatusNotFound)
			}
			return
		}

		// Запрещаем доступ к директориям
		if info.IsDir() {
			NotFoundHandler(w, r)
			return
		}

		// Кеширование для статических файлов
		if strings.HasPrefix(r.URL.Path, "/static/") {
			w.Header().Set("Cache-Control", "public, max-age=31536000")
			w.Header().Set("Expires", time.Now().Add(365*24*time.Hour).UTC().Format(http.TimeFormat))
		}

		// Отдаем файл без префикса /static/ или /source/
		http.StripPrefix("/", fs).ServeHTTP(w, r)
	})
}
