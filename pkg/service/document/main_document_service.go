package document

import (
	"documentum/pkg/models"
	"documentum/pkg/storage"
	"documentum/pkg/service/valid"
	"documentum/pkg/service/file"
)

type DocService interface {
	GetIngoingDoc(settings models.DocSettings) (string, error)
	AddLookDocument(id int, name string) error
	AddIngoingDoc(doc models.Document) (models.Document, error)
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