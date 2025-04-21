package service

import (
	"documentum/pkg/models"
	"documentum/pkg/storage"
)

type PageService interface {
	GetFuncs() ([]models.Unit, error)
}

type pageService struct {
	storage storage.PageStorage
}

func NewPageService(storage storage.PageStorage) PageService {
	return &pageService{storage: storage}
}

func (s *pageService) GetFuncs() ([]models.Unit, error) {
	var funcs []models.Unit 
	funcs, err := s.storage.GetFuncs()
	if err != nil {
		return nil, err
	}
	return funcs, err
}