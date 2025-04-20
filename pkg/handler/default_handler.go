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

func StaticHandler(staticDir string) http.Handler {
	// Получаем абсолютный путь к статике
	absPath, err := filepath.Abs(staticDir)
	if err != nil {
		log.Println(err.Error())
		panic("Ошибка получения абсолютного пути: " + err.Error())
	}

	fs := http.FileServer(http.Dir(absPath))

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Блокируем доступ к скрытым файлам
		if strings.HasPrefix(filepath.Base(r.URL.Path), ".") {
			log.Println(err.Error())
			NotFoundHandler(w, r)
			return
		}

		// Проверяем существование файла
		reqPath := filepath.Join(absPath, strings.TrimPrefix(r.URL.Path, "/static/"))
		info, err := os.Stat(reqPath)
		if err != nil || info.IsDir() {
			log.Println(err.Error())
			NotFoundHandler(w, r)
			return
		}

		// Кеширование
		w.Header().Set("Cache-Control", "public, max-age=31536000")
		w.Header().Set("Expires", time.Now().Add(CacheDuration).UTC().Format(http.TimeFormat))

		// Отдаем файл
		http.StripPrefix("/static/", fs).ServeHTTP(w, r)
	})
}
