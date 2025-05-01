package storage

import (
	"documentum/pkg/models"
)

func (s *SQLStorage) UpdateUserPassword(login, pass string) error {
	_, err := s.db.Exec("UPDATE `users` SET `pass` = ? WHERE `login` = ?", pass, login)
	if err != nil {
		return s.log.Error(models.ErrUpdDataInDB, err)
	} 
	return nil
}

func (s *SQLStorage) UpdateUserIcon(icon string, login string) error {

	_, err := s.db.Exec("UPDATE `users` SET `icon` = ? WHERE `login` = ?", icon, login)

	if err != nil {
		return s.log.Error(models.ErrUpdDataInDB, err)
	}

	return nil
}

