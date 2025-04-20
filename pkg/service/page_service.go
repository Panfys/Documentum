package service

import (
	"documentum/pkg/models"
	"documentum/pkg/storage"
)

type PageService interface {
	GetFuncs() (*models.Unit, error)
	GetUnits(string) (*models.Unit, error)
	GetGroups(string, string) (*models.Unit, error)
}

type pageService struct {
	storage storage.PageStorage
}

func NewPageService(storage storage.PageStorage) PageService {
	return &pageService{storage: storage}
}

func (s *pageService) GetFuncs() (*models.Unit, error) {
	
	return s.storage.GetFuncs()
}

func (s *pageService) GetUnits(param string) (*models.Unit, error) {
	
	return s.storage.GetUnits(param)
}


func (s *pageService) GetGroups(param1, param2 string) (*models.Unit, error) {

	return s.storage.GetGroups(param1, param2)
}