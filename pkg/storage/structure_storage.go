package storage

import (
	"documentum/pkg/models"
	"errors"
)

func (s *SQLStorage) GetUnits(function string) ([]models.Unit, error) {

	rows, err := s.db.Query("SELECT units_id, units.name FROM `funcs_units` JOIN `units` ON units_id = units.id WHERE `funcs_id` = ?", function)
	if err != nil {
		s.log.Error("ошибка при получении структурных подразделений в БД: %v", err)
		return nil, errors.New("ошибка при получении структурных подразделений в БД")
	}
	defer rows.Close()

	var units []models.Unit

	for rows.Next() {
		var unit models.Unit
		if err := rows.Scan(&unit.ID, &unit.Name); err != nil {
			s.log.Error("ошибка при получении структурных подразделений в БД: %v", err)
			return nil, errors.New("ошибка при получении структурных подразделений в БД")
		}
		units = append(units, unit)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return units, nil
}

func (p *SQLStorage) GetGroups(function, unit string) ([]models.Unit, error) {
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

func (p *SQLStorage) GetFuncs() ([]models.Unit, error) {

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
