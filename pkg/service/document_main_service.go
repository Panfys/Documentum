package service

import (
	//"documentum/pkg/models"
	"documentum/pkg/models"
	"documentum/pkg/storage"
	"errors"
	"fmt"
	"strconv"
)

type DocService interface {
	GetIngoingDoc(settings models.DocSettings) (string, error)
	AddLookDocument(id int, name string) error
	AddIngoingDoc(doc models.Document) (models.Document, error)
}

type docService struct {
	storage storage.DocStorage
}

func NewDocService(storage storage.DocStorage) DocService {
	return &docService{storage: storage}
}

func (d *docService) GetIngoingDoc(settings models.DocSettings) (string, error) {
	var documents []models.Document

	documents, err := d.storage.GetDocuments(settings)

	if err != nil {
		return "", err
	}

	var response string

	for _, document := range documents {

		// Обработка даты
		var newFDate, newLDate string
		
		newFDate, err = models.ParseDate(document.FDate.String)

		if err != nil {
			return "", fmt.Errorf("ошибка валидации даты1: %s", err)
		}

		if document.LDate.Valid {
			newLDate, err = models.ParseDate(document.LDate.String)
			if err != nil {
				return "", fmt.Errorf("ошибка валидации даты2: %s", err)
			}
		}

		// Обработка резолюции

		resolutions, err := d.storage.GetResolutoins(document.ID)

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
					newTime, err = models.ParseTime(resolution.Time.String)
					if err != nil {
						return "", fmt.Errorf("ошибка валидации даты: %s", err)
					}
				}

				newDate, err = models.ParseResolutionDate(resolution.Date)
				if err != nil {
					return "", fmt.Errorf("ошибка валидации даты: %s", err)
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
			strconv.Itoa(document.Width),
			document.Location,
			document.FileURL,
			docResolutions)
	}

	return response, nil
}

func (d *docService) AddLookDocument(id int, login string) error {

	name, err := d.storage.GetUserName(login)
	if err != nil {
		return err
	}

	err = d.storage.AddLookDocument(id, name)

	if err != nil {
		return fmt.Errorf("ошибка записи просмотра документа в БД")
	}

	return nil
}

func (d *docService) AddIngoingDoc(doc models.Document) (models.Document, error) {

	cleanDoc := d.SanitizeDocument(doc)
	id, err := d.storage.GetAutoIncrement("doc")

	if err != nil {
		return cleanDoc, errors.New("ошибка получения автоинкремента")
	}
	cleanDoc.ID = id

	err = d.validIngoingDoc(cleanDoc)
	if err != nil {
		return cleanDoc, err
	}

	return cleanDoc, nil 

}
