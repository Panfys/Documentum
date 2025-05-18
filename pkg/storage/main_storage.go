package storage

import (
	"database/sql"
	"documentum/pkg/logger"
	"documentum/pkg/models"
)

type SQLStorage struct {
	db  *sql.DB
	log logger.Logger
}

func NewSQLStorage(db *sql.DB, log logger.Logger) *SQLStorage {
	return &SQLStorage{
		db:  db,
		log: log,
	}
}

type AuthStorage interface {
	GetFuncs() ([]models.Unit, error)
	GetUserExists(login string) (bool, error)
	AddUser(user models.User) error
	GetUserPassByLogin(login string) (string, error)
}

type UserStorage interface {
	GetUserPassByLogin(login string) (string, error)
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
	GetUserName(login string) (string, error)
	GetDocuments(settings models.DocSettings) ([]models.Document, error)
	GetDirectives(settings models.DocSettings) ([]models.Directive, error)
	GetInventory(settings models.DocSettings) ([]models.Inventory, error)
	GetResolutoins(id int) ([]models.Resolution, error)
	UpdateDocFamiliar(table, id, name string) (int64, error)
	AddDocumentWithResolutions(doc models.Document) error
	AddDirective(doc models.Directive) error
	AddInventory(doc models.Inventory) error
	UpdateDocumentWithResolutions(doc models.Document) error 
}
