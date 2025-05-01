package user

import (
	"documentum/pkg/logger"
	"documentum/pkg/models"
	"documentum/pkg/service/file"
	"documentum/pkg/service/valid"
	"documentum/pkg/storage"
	"mime/multipart"
)

type UserService interface {
	UpdateUserPassword(login, pass, newPass string) (int, error)
	UpdateUserIcon(login string, icon multipart.File, iconName string) (string, error)
	GetUserAccountData(login string) (models.AccountData, error)
}

type userService struct {
	log      logger.Logger
	stor     storage.UserStorage
	validSrv valid.UserValidatService
	fileSrv  file.FileServece
}

func NewUserService(log logger.Logger, stor storage.UserStorage, validSrv valid.UserValidatService, fileSrv file.FileServece) UserService {
	return &userService{
		log: log,
		stor:     stor,
		validSrv: validSrv,
		fileSrv:  fileSrv,
	}
}
