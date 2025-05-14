package storage

import (
	"documentum/pkg/models"
	"fmt"
)

func (s *SQLStorage) GetDocuments(settings models.DocSettings) ([]models.Document, error) {

	query := fmt.Sprintf("SELECT * FROM `documents` WHERE `type` = ? AND `fdate` BETWEEN ? AND ? ORDER BY %s %s", settings.DocCol, settings.DocSet)

	rows, err := s.db.Query(query, settings.DocType, settings.DocDatain, settings.DocDatato)

	if err != nil {
		return nil, s.log.Error(models.ErrGetDataInDB, err," Запрос: ", query)
	}

	defer rows.Close()

	var	documents []models.Document

	for rows.Next() {
		var document models.Document

		if err := rows.Scan(&document.ID, &document.Type, &document.FNum, &document.FDate, &document.LNum, &document.LDate, &document.Name, &document.Sender, &document.Ispolnitel, &document.Result, &document.Familiar, &document.Count, &document.Copy, &document.Width, &document.Location, &document.FileURL, &document.Creator, &document.CreatedAt); err != nil {
			return nil, s.log.Error(models.ErrGetDataInDB, err)
		}

		documents = append(documents, document)
	}

	if err := rows.Err(); err != nil {
		return nil, s.log.Error(models.ErrGetDataInDB, err)
	}

	return documents, nil
}

func (s *SQLStorage) GetDirectives(settings models.DocSettings) ([]models.Directive, error) {

	query := fmt.Sprintf("SELECT * FROM `directives` WHERE `date` BETWEEN ? AND ? ORDER BY %s %s", settings.DocCol, settings.DocSet)

	rows, err := s.db.Query(query, settings.DocDatain, settings.DocDatato)

	if err != nil {
		return nil, s.log.Error(models.ErrGetDataInDB, err," Запрос: ", query)
	}

	defer rows.Close()

	var	directives []models.Directive

	for rows.Next() {
		var directive models.Directive

		if err := rows.Scan(&directive.ID ,&directive.Number, &directive.Date, &directive.Name, &directive.Autor, &directive.NumCoverLetter, &directive.DateCoverLetter, &directive.CountCopy, &directive.Sender, &directive.NumSendLetter, &directive.DateSendLetter, &directive.CountSendCopy, &directive.Familiar, &directive.Location, &directive.FileURL, &directive.Creator, &directive.CreatedAt); err != nil {
			return nil, s.log.Error(models.ErrGetDataInDB, err)
		}

		directives = append(directives, directive)
	}

	if err := rows.Err(); err != nil {
		return nil, s.log.Error(models.ErrGetDataInDB, err)
	}

	return directives, nil
}

func (s *SQLStorage) GetResolutoins(id int) ([]models.Resolution, error) {

	rows, err := s.db.Query("SELECT * FROM `resolutions` WHERE `doc_id` = ?", id)

	if err != nil {
		return nil, s.log.Error(models.ErrGetDataInDB, err," Запрос: ", "SELECT * FROM `resolutions`...")
	}

	defer rows.Close()

	var resolutions []models.Resolution

	for rows.Next() {
		var resolution models.Resolution

		if err := rows.Scan(&resolution.ID, &resolution.Type, &resolution.DocID, &resolution.Ispolnitel, &resolution.Text, &resolution.Deadline, &resolution.Date, &resolution.User, &resolution.Creator, &resolution.CreatedAt); err != nil {
			return nil, s.log.Error(models.ErrGetDataInDB, err)
		}

		resolutions = append(resolutions, resolution)
	}

	if err := rows.Err(); err != nil {
		return nil, s.log.Error(models.ErrGetDataInDB, err)
	}

	return resolutions, nil
}
