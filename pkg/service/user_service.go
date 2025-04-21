package service

import (
	"documentum/pkg/models"
	"documentum/pkg/storage"
	"fmt"
)

type UserService interface {
	GetUnits(function string) (string, error)
	GetGroups(function, unit string) (string, error)
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