package user

import (
	"errors"
	"mime/multipart"
	"path/filepath"
	"golang.org/x/crypto/bcrypt"
	"documentum/pkg/models"
)

func (s *userService) UpdateUserPassword(login, pass, newPass string) (int, error) {

	userPass, err := s.stor.GetUserPassByLogin(login)

	if err != nil {
		return 500, err
	}

	// Проверяем валидность текущего пароля
	if err := bcrypt.CompareHashAndPassword([]byte(userPass), []byte(pass)); err != nil {
		return 400, errors.New("текущий пароль неверный")
	}

	// Валидация пароля
	if !s.validSrv.ValidUserPass(newPass) || pass == newPass {
		return 400, errors.New("неверный формат нового пароля")
	}

	// Хешируем новый пароль
	newHash, err := bcrypt.GenerateFromPassword([]byte(newPass), bcrypt.DefaultCost)
	if err != nil {
		return 500, s.log.Error(models.ErrGetDataInDB, err)
	}

	err = s.stor.UpdateUserPassword(login, string(newHash))
	if err != nil {
		return 500, err
	}

	return 0, nil
}

// Метод для изменения иконки пользователя
func (s *userService) UpdateUserIcon(login string, icon multipart.File, iconName string) (string, error) {
	
	path := "/app/web/source/icons/"

	oldIconName, err := s.stor.GetUserIcon(login)
	if err != nil {
		return "", err
	}

	if !s.validSrv.ValidUserIcon(iconName) {
		return "", errors.New("неподдерживаемый формат файла")
	}

	newFilename, err := s.fileSrv.AddFile(path, iconName, icon)
	if err != nil {
		return "", err
	}

	filePath := filepath.Join(path, newFilename) 
	storagePath := filepath.Join("/source/icons/", newFilename)

	if err := s.stor.UpdateUserIcon(storagePath, login); err != nil {
		s.fileSrv.DeleteFileIfExists(filePath)
		return "", err
	}

	if oldIconName != "" {
		oldIconPath := filepath.Join("/app/web", oldIconName)
		s.fileSrv.DeleteFileIfExists(oldIconPath)
	}

	return storagePath, nil
}

