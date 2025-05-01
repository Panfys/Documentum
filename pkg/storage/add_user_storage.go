package storage

import (
	"documentum/pkg/models"
)

func (s *SQLStorage) AddUser(user models.User) error {
	
	newUser := "INSERT INTO `users` SET `login` = ?, `name` = ?, `func_id` = ?, `unit_id` = ?, `group_id` = ?, `pass` = ?, `status` = ?, `icon` = ?"

	_, err := s.db.Exec(newUser, user.Login, user.Name, user.Func, user.Unit, user.Group, user.Pass, "Пользователь", "")

	if err != nil {
		return s.log.Error(models.ErrAddDataInDB, err) 
	}

	return nil
}