package storage

import (
	"database/sql"
	"documentum/pkg/models"
)

type UserStorage interface {
	GetUser(id int) (*models.User, error)
}

// PostgresUserRepository - реализация для PostgreSQL
type SQLUserStorage struct {
	db *sql.DB
}

func NewUserStorage(db *sql.DB) *SQLUserStorage {
	return &SQLUserStorage{db: db}
}

func (r *SQLUserStorage) GetUser(id int) (*models.User, error) {
	/*var user models.User
	err := r.db.QueryRow(`
		SELECT id, name, email 
		FROM users 
		WHERE id = ?
	`, id).Scan(&user.ID, &user.Name, &user.Email)
	
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		return nil, err
	}
	return &user, nil
	*/
	return nil, nil
}
