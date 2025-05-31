package document

import (
	"documentum/pkg/models"
)

func (s *docService) GetDocuments(settings models.DocSettings) ([]models.Document, error) {
	set, err := s.settingCorrecter(settings)
	if err != nil {
		return nil, err
	}

	documents, err := s.stor.GetDocuments(set)
	if err != nil {
		return nil, err
	}

	for i := range documents {
		documents[i].Resolutions, err = s.stor.GetResolutoins(int64(documents[i].ID))
		if err != nil {
			return nil, err
		}
		doc, err := s.prepareDocument(&documents[i])
		if err != nil {
			return nil, err
		}
		documents[i] = *doc
	}

	return documents, nil
}

func (s *docService) GetDirectives(settings models.DocSettings) ([]models.Directive, error) {

	var directives []models.Directive

	set, err := s.settingCorrecter(settings)
	if err != nil {
		return []models.Directive{}, err
	}

	directives, err = s.stor.GetDirectives(set)

	if err != nil {
		return []models.Directive{}, err
	}

	for i := range directives {

		dir, err := s.prepareDirective(&directives[i])
		if err != nil {
			return nil, err
		}
		directives[i] = *dir
	}

	return directives, nil
}

func (s *docService) GetInventory(settings models.DocSettings) ([]models.Inventory, error) {
	var inventory []models.Inventory

	set, err := s.settingCorrecter(settings)
	if err != nil {
		return []models.Inventory{}, err
	}

	inventory, err = s.stor.GetInventory(set)

	if err != nil {
		return []models.Inventory{}, err
	}

	for i := range inventory {

		if inventory[i].DateCoverLetter.Valid {
			inventory[i].DateCoverLetterStr, err = s.prepareDate(inventory[i].DateCoverLetter.String)
			if err != nil {
				return []models.Inventory{}, err
			}
		}

		if inventory[i].DateSendLetter.Valid {
			inventory[i].DateSendLetterStr, err = s.prepareDate(inventory[i].DateSendLetter.String)
			if err != nil {
				return []models.Inventory{}, err
			}
		}
	}

	return inventory, nil
}

func (s *docService) settingCorrecter(settings models.DocSettings) (models.DocSettings, error) {
	if settings.DocSet == "" {
		settings.DocSet = "ASC"
	}

	if settings.DocCol == "" {
		settings.DocCol = "id"
	}

	if settings.DocDatain == "" {
		settings.DocDatain = "2000-01-01"
	}

	if settings.DocDatato == "" {
		settings.DocDatato = "3000-01-01"
	}

	return settings, nil
}
