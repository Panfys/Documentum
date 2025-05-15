package document

import (
	"documentum/pkg/models"
	"documentum/pkg/storage"
	"documentum/pkg/service/valid"
	"documentum/pkg/service/file"
)

type DocService interface {
	GetDocuments(settings models.DocSettings) ([]models.Document, error)
	GetDirectives(settings models.DocSettings) ([]models.Directive, error)
	GetInventory(settings models.DocSettings) ([]models.Inventory, error)
	AddFamiliarDocument(table, id, login string) error
	AddDocument(doc models.Document) (models.Document, error)
	AddDirective(dir models.Directive) (models.Directive, error)
	AddInventory(reqInv models.Inventory) (models.Inventory, error)
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