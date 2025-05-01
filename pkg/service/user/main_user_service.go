package user

import (
	"documentum/pkg/service/valid"
	"documentum/pkg/service/file"
	"documentum/pkg/storage"
	"documentum/pkg/models"
	"mime/multipart"
)

type UserService interface {
	UpdateUserPassword(login, pass, newPass string) (int, error)
	UpdateUserIcon(login string, icon multipart.File, iconName string) (string, error)
	GetUserAccountData(login string) (models.AccountData, error)
}

type userService struct {
	stor  storage.UserStorage
	validSrv valid.UserValidatService
	fileSrv file.FileServece
}

func NewUserService(stor storage.UserStorage, validSrv valid.UserValidatService, fileSrv file.FileServece) UserService {
	return &userService{
		stor:  stor,
		validSrv: validSrv,
		fileSrv: fileSrv,
	}
}
