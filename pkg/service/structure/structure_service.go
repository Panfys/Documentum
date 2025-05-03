package structure

import (
	"documentum/pkg/models"
	"documentum/pkg/storage"
)

type StructureService interface {
	GetUnits(function string) ([]models.Unit, error) 
	GetGroups(function, unit string) ([]models.Unit, error)
	GetFuncs() ([]models.Unit, error)
}

type structureService struct {
	stor storage.StructureStorage
}

func NewstructureService(stor storage.StructureStorage) StructureService {
	return &structureService{stor: stor}
}

func (s *structureService) GetUnits(function string) ([]models.Unit, error) {

	var units []models.Unit

	units, err := s.stor.GetUnits(function)

	if err != nil {
		return []models.Unit{}, err
	}
	return units, nil
}

func (s *structureService) GetGroups(function, unit string) ([]models.Unit, error) {

	var groups []models.Unit

	groups, err := s.stor.GetGroups(function, unit)

	if err != nil {
		return []models.Unit{}, err
	}

	return groups, nil
}

func (s *structureService) GetFuncs() ([]models.Unit, error) {
	var funcs []models.Unit 
	funcs, err := s.stor.GetFuncs()
	if err != nil {
		return nil, err
	}
	return funcs, err
}