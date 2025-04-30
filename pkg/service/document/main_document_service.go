package service

import (
	"documentum/pkg/models"
	"documentum/pkg/storage"
	"documentum/pkg/service/valid"
)

type DocService interface {
	GetIngoingDoc(settings models.DocSettings) (string, error)
	AddLookDocument(id int, name string) error
	AddIngoingDoc(doc models.Document) (models.Document, error)
}

type docService struct {
	stor storage.DocStorage
	valid valid.DocValidator
}

func NewDocService(stor storage.DocStorage, valid valid.DocValidator) DocService {
	return &docService{
		stor: stor,
		valid: valid,
	}
}