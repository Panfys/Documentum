package storage

import (
	"documentum/pkg/models"
)

func (s *SQLStorage) GetUnits(function string) ([]models.Unit, error) {

	rows, err := s.db.Query("SELECT units_id, units.name FROM `funcs_units` JOIN `units` ON units_id = units.id WHERE `funcs_id` = ?", function)
	if err != nil {
		
		return nil, s.log.Error(models.ErrGetDataInDB, err)
	}
	defer rows.Close()

	var units []models.Unit

	for rows.Next() {
		var unit models.Unit
		if err := rows.Scan(&unit.ID, &unit.Name); err != nil {
			return nil, s.log.Error(models.ErrGetDataInDB, err)
		}
		units = append(units, unit)
	}

	if err := rows.Err(); err != nil {
		return nil, s.log.Error(models.ErrGetDataInDB, err)
	}

	return units, nil
}

func (s *SQLStorage) GetGroups(function, unit string) ([]models.Unit, error) {
	rows, err := s.db.Query("SELECT groups.id, groups.name FROM `funcs_groups` JOIN `groups` ON groups_id = groups.id WHERE `funcs_id` = ? AND `units_id` = ?", function, unit)

	if err != nil {
		return nil, s.log.Error(models.ErrGetDataInDB, err)
	}
	defer rows.Close()

	var groups []models.Unit

	for rows.Next() {
		var group models.Unit
		if err := rows.Scan(&group.ID, &group.Name); err != nil {
			return nil, s.log.Error(models.ErrGetDataInDB, err)
		}
		groups = append(groups, group)
	}

	if err := rows.Err(); err != nil {
		return nil, s.log.Error(models.ErrGetDataInDB, err)
	}

	return groups, nil
}

func (s *SQLStorage) GetFuncs() ([]models.Unit, error) {

	rows, err := s.db.Query("SELECT id, name FROM funcs")
	if err != nil {
		return nil, s.log.Error(models.ErrGetDataInDB, err)
	}
	defer rows.Close()

	var funcs []models.Unit

	for rows.Next() {
		var functions models.Unit
		if err := rows.Scan(&functions.ID, &functions.Name); err != nil {
			return nil, s.log.Error(models.ErrGetDataInDB, err)
		}
		funcs = append(funcs, functions)
	}

	if err := rows.Err(); err != nil {
		return nil, s.log.Error(models.ErrGetDataInDB, err)
	}

	return funcs, nil
}
