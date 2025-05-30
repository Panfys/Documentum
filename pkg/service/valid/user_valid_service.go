package valid

import (
	"regexp"
	"strings"
	"path/filepath"
)

type UserValidatService interface {
	ValidUserLogin(login string) bool
	ValidUserName(name string) bool
	ValidUserIcon(iconPath string) bool
	ValidUserPass(pass string) bool
}

func (v *validatService) ValidUserLogin(login string) bool {

	pattern := `^[a-zA-Z0-9](?:[a-zA-Z0-9._-]{1,10}[a-zA-Z0-9])?$`
	re := regexp.MustCompile(pattern)
	return re.MatchString(login) && len(login) >= 3 && len(login) <= 12
}

func (v *validatService) ValidUserName(name string) bool {

	pattern := `^[А-ЯЁ][а-яё]+ [А-ЯЁ]\.[А-ЯЁ]\.$`
	re := regexp.MustCompile(pattern)

	return re.MatchString(name)

}

func (v *validatService) ValidUserIcon(iconPath string) bool {
	// Проверка расширения
	ext := strings.ToLower(filepath.Ext(iconPath))
	switch ext {
	case ".jpg", ".jpeg", ".png", ".webp", ".gif", ".ico", ".icon":
		return true
	default:
		return false
	}
} 

func (v *validatService) ValidUserPass(pass string) bool {

	if len(pass) < 6 || len(pass) > 64 {
		return false
	}

	return true
}