package storage

import (
	"database/sql"
	"documentum/pkg/logger"
	"documentum/pkg/models"
	"time"
	"log"
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

func ConnectToDB(connectionString string) (*sql.DB, error) {
	maxRetries := 5
	retryDelay := 10 * time.Second
	var db *sql.DB
	var err error

	for i := 0; i < maxRetries; i++ {
		db, err = sql.Open("mysql", connectionString)
		if err != nil {
			log.Printf("Попытка %d: Ошибка подключения к БД: %v", i+1, err)
			time.Sleep(retryDelay)
			continue
		}

		// Проверяем, что подключение действительно работает
		err = db.Ping()
		if err != nil {
			log.Printf("Попытка %d: Ошибка ping БД: %v", i+1, err)
			db.Close()
			time.Sleep(retryDelay)
			continue
		}

		return db, nil
	}

	return nil, err
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
	GetResolutoins(id int64) ([]models.Resolution, error)
	AddDocumentWithResolutions(doc models.Document) (int64, error)
	AddDirective(doc models.Directive) (int64, error)
	AddInventory(doc models.Inventory) (int64, error)
	UpdateDocFamiliar(types, id, name string) (int64, error)
	UpdateDocumentWithResolutions(doc models.Document) error 
}
