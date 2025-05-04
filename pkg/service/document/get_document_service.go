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

				if documents[i].Resolutions[j].Time.Valid {
					documents[i].Resolutions[j].TimeStr, err = s.parseTime(documents[i].Resolutions[j].Time.String)
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

/*
func (s *docService) GetDocuments(settings models.DocSettings) (string, error) {
	var documents []models.Document

	documents, err := s.stor.GetDocuments(settings)

	if err != nil {
		return "", err
	}

	var response string

	for _, document := range documents {

		// Обработка даты
		var newFDate, newLDate string

		newFDate, err = s.parseDate(document.FDate)

		if err != nil {
			return "", err
		}

		if document.LDate.Valid{
			newLDate, err = s.parseDate(document.LDate.String)
			if err != nil {
				return "", err
			}
		}

		// Обработка резолюции

		resolutions, err := s.stor.GetResolutoins(document.ID)

		if err != nil {
			return "", err
		}

		var newTime, newDate, docResolution string

		if len(resolutions) > 0 {
			resolution := resolutions[len(resolutions)-1]

			// Сборка исполнителя
			document.Ispolnitel = fmt.Sprintf("<div class='table__ispolnitel--ispolnitel'>%s</div>"+
				"<div class='table__ispolnitel--text'>&#171%s&#187</div>"+
				"<div class='table__ispolnitel--user'>%s</div>",
				resolution.Ispolnitel, resolution.Text, resolution.User)

			// Сборка резолюций
			for _, resolution := range resolutions {

				if resolution.Time.Valid {
					newTime, err = s.parseTime(resolution.Time.String)
					if err != nil {
						return "", err
					}
				}

				newDate, err = s.parseResolutionDate(resolution.Date)
				if err != nil {
					return "", err
				}

				docResolution += fmt.Sprintf("<div class='table__resolution' id='ingoing-resolution'> "+
					"<div class='table__resolution--ispolnitel'>%s</div>"+
					"<div class='table__resolution--text'>&#171%s&#187</div>"+
					"<div class='table__resolution--time'>%s</div>"+
					"<div class='table__resolution--user'>%s</div>"+
					"<div class='table__resolution--date'>%s</div></div>",
					resolution.Ispolnitel,
					resolution.Text,
					newTime,
					resolution.User,
					newDate)
			}
		}

		docResolutions := fmt.Sprintf("<div class='table__resolution-panel' id='resolution-panel-%d'>%s</div>", document.ID, docResolution)

		// Сборка документа
		response += fmt.Sprintf(
			"<table class='tubs__table tubs__table--document' id='document-table-%d' document-id='%d'>"+
				"<tr>"+
				"<td class='table__column--number'>%s %s</td>"+
				"<td class='table__column--number'>%s %s</td>"+
				"<td class='table__column--name'>%s</td>"+
				"<td class='table__column--sender'>%s</td>"+
				"<td class='table__column--ispolnitel'>%s</td>"+
				"<td class='table__column--result'>%s</td>"+
				"<td class='table__column--familiar'>%s</td>"+
				"<td class='table__column--count'>%s</td>"+
				"<td class='table__column--copy'>%s</td>"+
				"<td class='table__column--width'>%s</td>"+
				"<td class='table__column--location'>%s</td>"+
				"<td class='table__column--button'><button class='table__btn--opendoc' file=%s></button></td>"+
				"</tr>"+
				"</table>%s",
			document.ID, document.ID,
			document.FNum, newFDate,
			document.LNum, newLDate,
			document.Name,
			document.Sender,
			document.Ispolnitel,
			document.Result,
			document.Familiar,
			strconv.Itoa(document.Count),
			document.Copy,
			document.Width,
			document.Location,
			document.FileURL,
			docResolutions)
	}

	return response, nil
} */