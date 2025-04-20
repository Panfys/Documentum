package service

import (
	"documentum/pkg/models"
	"documentum/pkg/storage"
)

type UserService interface {
	GetUser(id int) (*models.User, error)
}

type userService struct {
	storage storage.UserStorage
}

func NewUserService(storage storage.UserStorage) UserService {
	return &userService{storage: storage}
}

func (s *userService) GetUser(id int) (*models.User, error) {
	// Здесь может быть дополнительная бизнес-логика
	// Например, проверка прав доступа, кэширование и т.д.
	return s.storage.GetUser(id)
}