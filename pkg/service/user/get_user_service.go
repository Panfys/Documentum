package user

import (
	"documentum/pkg/models"
	"time"
)

func (s *userService) GetUserAccountData(login string) (models.AccountData, error) {
	var accountData models.AccountData
	accountData, err := s.stor.GetAccountData(login)

	if err != nil {
		return accountData, err
	}

	now := time.Now()

	accountData.Login = login
	accountData.ToDay = now.Format("2006-01-02")
	
	return accountData, nil
}	
