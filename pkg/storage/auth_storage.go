package storage

import (
	"database/sql"
	"documentum/pkg/models"
	"fmt"
)

type AuthStorage interface {
	GetAccountData(login string) (models.User, error)
}

type JVTAuthStorage struct {
	db *sql.DB
}

func NewAuthStorage(db *sql.DB) *JVTAuthStorage {
	return &JVTAuthStorage{db: db}
}

// Получение данных о пользователе по токену
func (s *JVTAuthStorage) GetAccountData(login string) (models.User, error) {
	var accountData models.User

	err := s.db.QueryRow("SELECT users.name, funcs.fullname_f , units.fullname_u, groups.fullname_g, users.status, users.icon FROM `users` JOIN `funcs` ON funcs.id = func_id JOIN `units` ON units.id = unit_id JOIN `groups` ON groups.id = group_id WHERE users.login = ?", login).Scan(&accountData.Name, &accountData.Func, &accountData.Unit, &accountData.Group, &accountData.Status, &accountData.Icon)

	if err != nil {
		return accountData, fmt.Errorf("ошибка обработки данных")
	}

	accountData.Login = login

	return accountData, nil
}