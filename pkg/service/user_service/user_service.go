package user_service

import (
	"documentum/pkg/storage"
)

type UserService interface {
	GetUserService
	UpdateUserService
}

type userService struct {
	storage storage.UserStorage
}

func NewUserService(storage storage.UserStorage) UserService {
	return &userService{storage: storage}
}
