package document

import (
	"documentum/pkg/models"
	"fmt"
	"time"
)

func (s *docService) prepareDocument(doc *models.Document) (*models.Document, error) {
	var err error

	doc.FDate, err = s.prepareDate(doc.FDate)
	if err != nil {
		return nil, err
	}

	if doc.LDate.Valid {
		doc.LDateStr, err = s.prepareDate(doc.LDate.String)
		if err != nil {
			return nil, err
		}
	}

	// Обработка резолюций
	if len(doc.Resolutions) > 0 {
		if err := s.prepareResolutions(doc); err != nil {
			return nil, err
		}
	}

	return doc, nil
}

func (s *docService) prepareDirective(dir *models.Directive) (*models.Directive, error) {
	var err error

	dir.Date, err = s.prepareDate(dir.Date)
	if err != nil {
		return nil, err
	}

	if dir.DateCoverLetter.Valid {
		dir.DateCoverLetterStr, err = s.prepareDate(dir.DateCoverLetter.String)
		if err != nil {
			return nil, err
		}
	}

	if dir.DateSendLetter.Valid {
		dir.DateSendLetterStr, err = s.prepareDate(dir.DateSendLetter.String)
		if err != nil {
			return nil, err
		}
	}
	return dir, nil
}

func (s *docService) prepareInventory(inv *models.Inventory) (*models.Inventory, error) {
	var err error

	if inv.DateCoverLetter.Valid {
		inv.DateCoverLetterStr, err = s.prepareDate(inv.DateCoverLetter.String)
		if err != nil {
			return nil, err
		}
	}

	if inv.DateSendLetter.Valid {
		inv.DateSendLetterStr, err = s.prepareDate(inv.DateSendLetter.String)
		if err != nil {
			return nil, err
		}
	}
	return inv, nil 
}

func (s *docService) prepareResolutions(doc *models.Document) error {

	// Обработка последней резолюции для Ispolnitel
	s.prepareLastResolution(doc)

	// Обработка всех резолюций
	for i := range doc.Resolutions {
		if err := s.prepareSingleResolution(&doc.Resolutions[i]); err != nil {
			return err
		}

		// Сборка исполненных документов
		s.prepareResolutionResult(doc, &doc.Resolutions[i])
	}

	return nil
}

func (s *docService) prepareLastResolution(doc *models.Document) {
	resolution := doc.Resolutions[len(doc.Resolutions)-1]
	doc.Ispolnitel = fmt.Sprintf("<div class='table__ispolnitel--ispolnitel'>%s</div>"+
		"<div class='table__ispolnitel--text'>&#171%s&#187</div>"+
		"<div class='table__ispolnitel--user'>%s</div>",
		resolution.Ispolnitel, resolution.Text, resolution.User)
}

func (s *docService) prepareSingleResolution(res *models.Resolution) error {
	var err error

	if res.Deadline.Valid {
		res.DeadlineStr, err = s.prepareResolutionDeadline(res.Deadline.String)
		if err != nil {
			return err
		}
	}

	res.Date, err = s.prepareResolutionDate(res.Date)
	if err != nil {
		return err
	}

	return nil
}

func (s *docService) prepareResolutionResult(doc *models.Document, res *models.Resolution) {
	if res.Result != "" {
		if doc.Result == "" {
			doc.Result += res.Result
		} else {
			doc.Result += "<br>" + res.Result
		}
	}
}

func (s *docService) prepareDate(date string) (string, error) {
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

func (s *docService) prepareResolutionDate(date string) (string, error) {

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

func (s *docService) prepareResolutionDeadline(restime string) (string, error) {

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
