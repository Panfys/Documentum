package user_service

import (
	"documentum/pkg/models"
	"errors"
	"mime/multipart"
	"os"
	"path/filepath"
	"golang.org/x/crypto/bcrypt"
)

type UpdateUserService interface {
	UpdateUserPassword(login, pass, newPass string) (int, error)
	UpdateUserIcon(login string, icon multipart.File, iconName string) (string, error)
}

func (s *userService) UpdateUserPassword(login, pass, newPass string) (int, error) {

	userPass, err := s.storage.GetUserPassword(login)

	if err != nil {
		return 500, err
	}

	// Проверяем валидность текущего пароля
	if err := bcrypt.CompareHashAndPassword([]byte(userPass), []byte(pass)); err != nil {
		return 400, errors.New("Текущий пароль неверный!")
	}

	// Валидация пароля
	if !models.ValidPass(newPass) || pass == newPass {
		return 400, errors.New("Неверный формат нового пароля!")
	}

	// Хешируем новый пароль
	newHash, err := bcrypt.GenerateFromPassword([]byte(newPass), bcrypt.DefaultCost)
	if err != nil {
		return 500, errors.New("ошибка хеширования нового пароля")
	}

	err = s.storage.UpdateUserPassword(login, string(newHash))
	if err != nil {
		return 500, errors.New("ошибка обновления пароля в БД")
	}

	return 0, nil
}

// Метод для изменения иконки пользователя
func (s *userService) UpdateUserIcon(login string, icon multipart.File, iconName string) (string, error) {
	path := "/app/web/source/icons/"

	oldIconName, err := s.storage.GetUserIcon(login)
	if err != nil {
		return "", err
	}

	if !models.ValidIcon(iconName) {
		return "", errors.New("неподдерживаемый формат файла")
	}

	newFilename, err := models.GenerateUniqueFilename(path, iconName)
	if err != nil {
		return "", err
	}

	filePath := filepath.Join(path, newFilename)
	if err := models.SaveFile(icon, filePath); err != nil {
		return "", err
	}

	storagePath := filepath.Join("/source/icons/", newFilename)

	if err := s.storage.UpdateUserIcon(storagePath, login); err != nil {
		os.Remove(filePath) // Откатываем изменения если ошибка
		return "", err
	}

	if oldIconName != "" {
		oldIconPath := filepath.Join("/app/web", oldIconName)
		models.DeleteFileIfExists(oldIconPath)
	}

	return storagePath, nil
}

