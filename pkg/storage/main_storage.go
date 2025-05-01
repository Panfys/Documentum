package storage

import (
	"database/sql"
	"documentum/pkg/models"
)

type SQLStorage struct {
	db *sql.DB
}

func NewSQLStorage(db *sql.DB) *SQLStorage {
	return &SQLStorage{db: db}
}

type AuthStorage interface {
	GetFuncs() ([]models.Unit, error)
	GetUserExists(login string) (bool, error)
	AddUser(user models.User) error
	GetUserPassByLogin(login string) (string, error)
}

type UserStorage interface {
	GetUserPassword(login string) (string, error)
	GetAccountData(login string) (models.AccountData, error)
	UpdateUserPassword(login, pass string) error 
	GetUserIcon(login string) (string, error)
	UpdateUserIcon(icon string, login string) error
}

type StructureStorage interface {
	GetUnits(function string) ([]models.Unit, error)
	GetGroups(function, unit string) ([]models.Unit, error)
	GetFuncs() ([]models.Unit, error)
}

type DocStorage interface {
	GetDocuments(settings models.DocSettings) ([]models.Document, error)
	GetResolutoins(id int) ([]models.Resolution, error)
	AddLookDocument(id int, name string) error
	GetUserName(login string) (string, error)
	GetAutoIncrement(table string) (int, error)
	AddDocumentWithResolutions(doc models.Document) error 
	AddDocument(doc models.Document) error
}