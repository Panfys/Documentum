package storage

import (

)

func (p *SQLStorage) UpdateUserPassword(login, pass string) error {
	_, err := p.db.Exec("UPDATE `users` SET `pass` = ? WHERE `login` = ?", pass, login)
	if err != nil {
		return err
	} 
	return nil
}

func (s *SQLStorage) UpdateUserIcon(icon string, login string) error {

	_, err := s.db.Exec("UPDATE `users` SET `icon` = ? WHERE `login` = ?", icon, login)

	if err != nil {
		return err
	}

	return nil
}

