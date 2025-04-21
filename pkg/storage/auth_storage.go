package storage

import (
	"database/sql"
	"documentum/pkg/models"
	"fmt"
)

type AuthStorage interface {
	AddUser(user models.User) error
	GetUserPass(login string) (string, error)
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

func (s *JVTAuthStorage) AddUser(user models.User) error {
	
	newUser := "INSERT INTO `users` SET `login` = ?, `name` = ?, `func_id` = ?, `unit_id` = ?, `group_id` = ?, `pass` = ?, `status` = ?, `icon` = ?"

	_, err := s.db.Exec(newUser, user.Login, user.Name, user.Func, user.Unit, user.Group, user.Pass, "Пользователь", "")

	if err != nil {
		return err
	}

	return nil
}

func (s *JVTAuthStorage) GetUserPass(login string) (string, error) {
	var pass string

	err := s.db.QueryRow("SELECT `pass` FROM `users` WHERE `login` = ?", login).Scan(&pass)

	if err != nil {
		return "", fmt.Errorf("ошибка обработки данных")
	}

	return pass, nil
}
