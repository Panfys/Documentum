package document

import (
	"documentum/pkg/models"
	"fmt"
	"time"
)

func (s *docService) GetDocuments(settings models.DocSettings) ([]models.Document, error) {
	var documents []models.Document

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
	
	documents, err := s.stor.GetDocuments(settings)

	if err != nil {
		return []models.Document{}, err
	}

	for i := range documents {

		documents[i].FDate, err = s.parseDate(documents[i].FDate)

		if err != nil {
			return []models.Document{}, err
		}

		if documents[i].LDate.Valid{
			documents[i].LDateStr, err = s.parseDate(documents[i].LDate.String)
			if err != nil {
				return []models.Document{}, err
			}
		}

		// Обработка резолюции

		documents[i].Resolutions, err = s.stor.GetResolutoins(documents[i].ID)

		if err != nil {
			return []models.Document{}, err
		}

		if len(documents[i].Resolutions) > 0 {
			resolution := documents[i].Resolutions[len(documents[i].Resolutions)-1]

			// Сборка исполнителя
			documents[i].Ispolnitel = fmt.Sprintf("<div class='table__ispolnitel--ispolnitel'>%s</div>"+
				"<div class='table__ispolnitel--text'>&#171%s&#187</div>"+
				"<div class='table__ispolnitel--user'>%s</div>",
				resolution.Ispolnitel, resolution.Text, resolution.User)

			// Сборка резолюций
			for j := range documents[i].Resolutions {

				if documents[i].Resolutions[j].Deadline.Valid {
					documents[i].Resolutions[j].DeadlineStr, err = s.parseTime(documents[i].Resolutions[j].Deadline.String)
					if err != nil {
						return []models.Document{}, err
					}
				}

				documents[i].Resolutions[j].Date, err = s.parseResolutionDate(documents[i].Resolutions[j].Date)
				if err != nil {
					return []models.Document{}, err
				}
			}
		}
	}

	return documents, nil
}

func (s *docService) GetDirectives(settings models.DocSettings) ([]models.Directive, error) {
	var directives []models.Directive

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
	
	directives, err := s.stor.GetDirectives(settings)

	if err != nil {
		return []models.Directive{}, err
	}

	for i := range directives {

		directives[i].Date, err = s.parseDate(directives[i].Date)

		if err != nil {
			return []models.Directive{}, err
		}

		if directives[i].DateCoverLetter.Valid{
			directives[i].DateCoverLetterStr, err = s.parseDate(directives[i].DateCoverLetter.String)
			if err != nil {
				return []models.Directive{}, err
			}
		}

		if directives[i].DateSendLetter.Valid{
			directives[i].DateSendLetterStr, err = s.parseDate(directives[i].DateSendLetter.String)
			if err != nil {
				return []models.Directive{}, err
			}
		}
	}

	return directives, nil
}

func (s *docService) GetInventory(settings models.DocSettings) ([]models.Inventory, error) {
	var inventory []models.Inventory

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
	
	inventory, err := s.stor.GetInventory(settings)

	if err != nil {
		return []models.Inventory{}, err
	}

	for i := range inventory {

		if inventory[i].DateCoverLetter.Valid{
			inventory[i].DateCoverLetterStr, err = s.parseDate(inventory[i].DateCoverLetter.String)
			if err != nil {
				return []models.Inventory{}, err
			}
		}

		if inventory[i].DateSendLetter.Valid{
			inventory[i].DateSendLetterStr, err = s.parseDate(inventory[i].DateSendLetter.String)
			if err != nil {
				return []models.Inventory{}, err
			}
		}
	}

	return inventory, nil
}

func (s *docService) parseDate(date string) (string, error) {
	newdate, err := time.Parse(time.RFC3339, date)
	if err != nil {
		newdate, err = time.Parse("2006-01-02", date)
		if err != nil {
			return "", err
		}
	}

	formattedDate := "от " + newdate.Format("02.01.2006") + " г."
	return formattedDate, nil
}

func (s *docService) parseResolutionDate(date string) (string, error) {

	newdate, err := time.Parse(time.RFC3339, date)
	if err != nil {
		newdate, err = time.Parse("2006-01-02", date)
		if err != nil {
			return "", err
		}
	}

	formateDate := newdate.Format("02.01.2006") + " г."
	return formateDate, nil
}

func (s *docService) parseTime(restime string) (string, error) {

	newtime, err := time.Parse(time.RFC3339, restime)
	if err != nil {
		newtime, err = time.Parse("2006-01-02", restime)
		if err != nil {
			return "", err
		}
	}

	// Форматируем дату в нужный формат
	formateTime := "Исполнить до " + newtime.Format("02.01.2006") + " г."
	return formateTime, nil
}
