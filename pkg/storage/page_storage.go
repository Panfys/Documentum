package storage

import (
	"database/sql"
	"documentum/pkg/models"
)

type PageStorage interface {
	GetFuncs() ([]models.Unit, error)
}
type SQLPageStorage struct {
	db *sql.DB
}

func NewPageStorage(db *sql.DB) *SQLPageStorage {
	return &SQLPageStorage{db: db}
}

func (p *SQLPageStorage) GetFuncs() ([]models.Unit, error) {
	
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