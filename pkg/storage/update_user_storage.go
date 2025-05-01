package storage

import "errors"

func (s *SQLStorage) UpdateUserPassword(login, pass string) error {
	_, err := s.db.Exec("UPDATE `users` SET `pass` = ? WHERE `login` = ?", pass, login)
	if err != nil {
		s.log.Error("ошибка при обновлении пароля пользователя в БД: %v", err)
		return errors.New("ошибка при обновлении пароля пользователя в БД")
	} 
	return nil
}

func (s *SQLStorage) UpdateUserIcon(icon string, login string) error {

	_, err := s.db.Exec("UPDATE `users` SET `icon` = ? WHERE `login` = ?", icon, login)

	if err != nil {
		s.log.Error("ошибка при обновлении иконки пользователя в БД", err)
		return errors.New("ошибка при обновлении иконки пользователя в БД")
	}

	return nil
}

