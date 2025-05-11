package document

import (
	"documentum/pkg/models"
	"documentum/pkg/storage"
	"documentum/pkg/service/valid"
	"documentum/pkg/service/file"
)

type DocService interface {
	GetDocuments(settings models.DocSettings) ([]models.Document, error)
	AddLookDocument(id string, name string) error
	AddDocument(doc models.Document) (models.Document, error)
}

type docService struct {
	stor storage.DocStorage
	validSrv valid.DocValidatService
	fileSrv file.FileServece
}

func NewDocService(stor storage.DocStorage, validSrv valid.DocValidatService, fileSrv file.FileServece) DocService {
	return &docService{
		stor: stor,
		validSrv: validSrv,
		fileSrv: fileSrv, 
	}
}