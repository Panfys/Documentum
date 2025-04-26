package storage

import (
	"database/sql"
	"documentum/pkg/models"
)

type UserStorage interface {
	GetGroups(function, unit string) ([]models.Unit, error)
	GetUnits(function string) ([]models.Unit, error)
	UpdateUserPassword(login, pass string) error
	GetUserPassword(login string) (string, error)
}
type SQLUserStorage struct {
	db *sql.DB
}

func NewUserStorage(db *sql.DB) *SQLUserStorage {
	return &SQLUserStorage{db: db}
}

func (p *SQLUserStorage) GetUnits(function string) ([]models.Unit, error) {

	rows, err := p.db.Query("SELECT units_id, units.name FROM `funcs_units` JOIN `units` ON units_id = units.id WHERE `funcs_id` = ?", function)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var units []models.Unit

	for rows.Next() {
		var unit models.Unit
		if err := rows.Scan(&unit.ID, &unit.Name); err != nil {
			return nil, err
		}
		units = append(units, unit)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return units, nil
}

func (p *SQLUserStorage) GetGroups(function, unit string) ([]models.Unit, error) {
	rows, err := p.db.Query("SELECT groups.id, groups.name FROM `funcs_groups` JOIN `groups` ON groups_id = groups.id WHERE `funcs_id` = ? AND `units_id` = ?", function, unit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var groups []models.Unit

	for rows.Next() {
		var group models.Unit
		if err := rows.Scan(&group.ID, &group.Name); err != nil {
			return nil, err
		}
		groups = append(groups, group)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return groups, nil
}

func (p *SQLUserStorage) UpdateUserPassword(login, pass string) error {
	_, err := p.db.Exec("UPDATE `users` SET `pass` = ? WHERE `login` = ?", pass, login)
	if err != nil {
		return err
	} 
	return nil
}

func (p *SQLUserStorage) GetUserPassword(login string) (string, error) {
	var pass string

	err := p.db.QueryRow("SELECT `pass` FROM `users` WHERE `login` = ?", login).Scan(&pass)

	if err != nil {
		return "oшибка обработки данных пользователя в БД", err
	}

	return pass, nil
}