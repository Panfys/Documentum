package models

import (
	"regexp"
	"strings"
	"path/filepath"
)

// Объект для работы с пользователями
type User struct {
	ID     int    `json:"id"`
	Login  string `json:"login"`
	Name   string `json:"name"`
	Func   string `json:"func"`
	Unit   string `json:"unit"`
	Group  string `json:"group"`
	Pass   string `json:"pass"`
	Status string `json:"status"`
	Icon   string `json:"icon"`
}

// Метод для валидации ФИО пользователя
func ValidName(name string) bool {

	pattern := `^[А-ЯЁ][а-яё]+ [А-ЯЁ]\.[А-ЯЁ]\.$`
	re := regexp.MustCompile(pattern)

	return re.MatchString(name)

}

// Метод для валидации логина пользователя
func ValidLogin(login string) bool {

	pattern := `^[a-zA-Z0-9](?:[a-zA-Z0-9._-]{1,10}[a-zA-Z0-9])?$`
	re := regexp.MustCompile(pattern)
	return re.MatchString(login) && len(login) >= 3 && len(login) <= 12
}

// Метод для валидации пароля пользователя
func ValidPass(pass string) bool {

	if len(pass) < 6 || len(pass) > 64 {
		return false
	}

	return true
}

func ValidIcon(iconPath string) bool {
	// Проверка расширения
	ext := strings.ToLower(filepath.Ext(iconPath))
	switch ext {
	case ".jpg", ".jpeg", ".png", ".webp", ".gif", ".ico", ".icon":
		return true
	default:
		return false
	}
}
