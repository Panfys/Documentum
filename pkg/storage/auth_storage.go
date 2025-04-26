package storage

import (
	"database/sql"
	"documentum/pkg/models"
	"fmt"
)

type AuthStorage interface {
	AddUser(user models.User) error
	GetUserPassByLogin(login string) (string, error)
	UserExists(login string) (bool, error)
	GetAccountData(login string) (models.AccountData, error)
	GetFuncs() ([]models.Unit, error)
}

type authStorage struct {
	db *sql.DB
}

func NewAuthStorage(db *sql.DB) *authStorage {
	return &authStorage{db: db}
}

func (p *authStorage) GetFuncs() ([]models.Unit, error) {
	
	rows, err := p.db.Query("SELECT id, name FROM funcs")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var funcs []models.Unit

	for rows.Next() {
		var functions models.Unit
		if err := rows.Scan(&functions.ID, &functions.Name); err != nil {
			return nil, err
		}
		funcs = append(funcs, functions)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return funcs, nil
}

// Получение данных о пользователе по токену
func (s *authStorage) GetAccountData(login string) (models.AccountData, error) {
	var accountData models.AccountData

	err := s.db.QueryRow("SELECT users.name, funcs.fullname_f , units.fullname_u, groups.fullname_g, users.status, users.icon FROM `users` JOIN `funcs` ON funcs.id = func_id JOIN `units` ON units.id = unit_id JOIN `groups` ON groups.id = group_id WHERE users.login = ?", login).Scan(&accountData.Name, &accountData.Func, &accountData.Unit, &accountData.Group, &accountData.Status, &accountData.Icon)

	if err != nil {
		return accountData, fmt.Errorf("ошибка обработки данных")
	}

	return accountData, nil
}

func (s *authStorage) AddUser(user models.User) error {
	
	newUser := "INSERT INTO `users` SET `login` = ?, `name` = ?, `func_id` = ?, `unit_id` = ?, `group_id` = ?, `pass` = ?, `status` = ?, `icon` = ?"

	_, err := s.db.Exec(newUser, user.Login, user.Name, user.Func, user.Unit, user.Group, user.Pass, "Пользователь", "")

	if err != nil {
		return err
	}

	return nil
}

func (s *authStorage) GetUserPassByLogin(login string) (string, error) {
	var pass string

	err := s.db.QueryRow("SELECT `pass` FROM `users` WHERE `login` = ?", login).Scan(&pass)

	if err != nil {
		return "", fmt.Errorf("ошибка обработки данных")
	}

	return pass, nil
}

func (s *authStorage) UserExists(login string) (bool, error) {
	var exists bool
	
	// Запрос проверяет наличие пользователя
	query := "SELECT EXISTS(SELECT 1 FROM users WHERE login = ?)"
	
	err := s.db.QueryRow(query, login).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("ошибка при проверке пользователя: %v", err)
	}
	
	return exists, nil
}
