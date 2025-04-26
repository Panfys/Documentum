package service

import (
	"documentum/pkg/models"
	"documentum/pkg/storage"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"errors"
)

type UserService interface {
	GetUnits(function string) (string, error)
	GetGroups(function, unit string) (string, error)
	UpdateUserPassword(login, pass, newPass string) (int, error)
}

type userService struct {
	storage storage.UserStorage
}

func NewUserService(storage storage.UserStorage) UserService {
	return &userService{storage: storage}
}

func (s *userService) GetUnits(function string) (string, error) {

	var units []models.Unit

	units, err := s.storage.GetUnits(function)

	if err != nil {
		return "", err
	}

	responseUnits := ""
	for _, unit := range units {
		responseUnits += fmt.Sprintf("<option value=%d>%s</option>", unit.ID, unit.Name)
	}

	return responseUnits, nil
}

func (s *userService) GetGroups(function, unit string) (string, error) {

	var groups []models.Unit

	groups, err := s.storage.GetGroups(function, unit)

	if err != nil {
		return "", err
	}

	responseGroups := ""
	for _, group := range groups {
		responseGroups += fmt.Sprintf("<option value=%d>%s</option>", group.ID, group.Name)
	}

	return responseGroups, nil
}

func (s *userService) UpdateUserPassword(login, pass, newPass string) (int, error) {
	
	userPass, err := s.storage.GetUserPassword(login)
	
	if err != nil {
		return 500, errors.New("какая-то ошибка")
	}

	// Проверяем валидность текущего пароля
	if err := bcrypt.CompareHashAndPassword([]byte(userPass), []byte(pass)); err != nil {
		return 400, errors.New("Текущий пароль неверный!")
	}

	var user models.User
	// Валидация пароля
	if !user.ValidPass(newPass) || pass == newPass{
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