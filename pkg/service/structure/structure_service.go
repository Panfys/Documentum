package structure

import (
	"documentum/pkg/models"
	"documentum/pkg/storage"
	"fmt"
)

type StructureService interface {
	GetUnits(function string) (string, error)
	GetGroups(function, unit string) (string, error)
	GetFuncs() ([]models.Unit, error)
}

type structureService struct {
	stor storage.StructureStorage
}

func NewstructureService(stor storage.StructureStorage) StructureService {
	return &structureService{stor: stor}
}

func (s *structureService) GetUnits(function string) (string, error) {

	var units []models.Unit

	units, err := s.stor.GetUnits(function)

	if err != nil {
		return "", err
	}

	responseUnits := ""
	for _, unit := range units {
		responseUnits += fmt.Sprintf("<option value=%d>%s</option>", unit.ID, unit.Name)
	}

	return responseUnits, nil
}

func (s *structureService) GetGroups(function, unit string) (string, error) {

	var groups []models.Unit

	groups, err := s.stor.GetGroups(function, unit)

	if err != nil {
		return "", err
	}

	responseGroups := ""
	for _, group := range groups {
		responseGroups += fmt.Sprintf("<option value=%d>%s</option>", group.ID, group.Name)
	}

	return responseGroups, nil
}

func (s *structureService) GetFuncs() ([]models.Unit, error) {
	var funcs []models.Unit 
	funcs, err := s.stor.GetFuncs()
	if err != nil {
		return nil, err
	}
	return funcs, err
}