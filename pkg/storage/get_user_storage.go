package storage

import (
	"documentum/pkg/models"
)

func (s *SQLStorage) GetUserIcon(login string) (string, error) {

	var icon string
	err := s.db.QueryRow("SELECT `icon` FROM `users` WHERE `login` = ?", login).Scan(&icon)

	if err != nil {
		return "", s.log.Error(models.ErrGetDataInDB, err," Запрос: ", "SELECT `icon`...")
	}

	return icon, nil
}

func (s *SQLStorage) GetAccountData(login string) (models.AccountData, error) {
	var accountData models.AccountData

	err := s.db.QueryRow("SELECT users.name, funcs.fullname_f , units.fullname_u, groups.fullname_g, users.status, users.icon FROM `users` JOIN `funcs` ON funcs.id = func_id JOIN `units` ON units.id = unit_id JOIN `groups` ON groups.id = group_id WHERE users.login = ?", login).Scan(&accountData.Name, &accountData.Func, &accountData.Unit, &accountData.Group, &accountData.Status, &accountData.Icon)

	if err != nil {
		return accountData, s.log.Error(models.ErrGetDataInDB, err," Запрос: ", "SELECT users.name, funcs.fullname_f...")
	}

	return accountData, nil
}

func (s *SQLStorage) GetUserPassByLogin(login string) (string, error) {
	var pass string

	err := s.db.QueryRow("SELECT `pass` FROM `users` WHERE `login` = ?", login).Scan(&pass)

	if err != nil {
		return "", s.log.Error(models.ErrGetDataInDB, err," Запрос: ", "SELECT `pass` FROM `users`...")
	}

	return pass, nil
}

func (s *SQLStorage) GetUserExists(login string) (bool, error) {
	var exists bool
	
	// Запрос проверяет наличие пользователя
	query := "SELECT EXISTS(SELECT 1 FROM users WHERE login = ?)"
	
	err := s.db.QueryRow(query, login).Scan(&exists)
	if err != nil {
		return false, s.log.Error(models.ErrGetDataInDB, err," Запрос: ", query)
	}
	
	return exists, nil
}

func (s *SQLStorage) GetUserName(login string) (string, error) {
	var name string

	err := s.db.QueryRow("SELECT `name` FROM `users` WHERE `login` = ?", login).Scan(&name)

	if err != nil {
		return "", s.log.Error(models.ErrGetDataInDB, err," Запрос: ", "SELECT `name` FROM `users`...")
	}

	return name, nil
}
