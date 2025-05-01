package storage

import (
	"documentum/pkg/models"
	"fmt"
)

func (s *SQLStorage) GetDocuments(settings models.DocSettings) ([]models.Document, error) {

	query := fmt.Sprintf("SELECT * FROM `doc` WHERE `type` = ? AND `fdate` BETWEEN ? AND ? ORDER BY %s %s", settings.DocCol, settings.DocSet)

	rows, err := s.db.Query(query, settings.DocType, settings.DocDatain, settings.DocDatato)

	if err != nil {
		return nil, s.log.Error(models.ErrGetDataInDB, err)
	}

	defer rows.Close()

	var	documents []models.Document

	for rows.Next() {
		var document models.Document

		if err := rows.Scan(&document.ID, &document.Type, &document.FNum, &document.FDate, &document.LNum, &document.LDate, &document.Name, &document.Sender, &document.Ispolnitel, &document.Result, &document.Familiar, &document.Count, &document.Copy, &document.Width, &document.Location, &document.FileURL, &document.Creator); err != nil {
			return nil, s.log.Error(models.ErrGetDataInDB, err)
		}

		documents = append(documents, document)
	}

	if err := rows.Err(); err != nil {
		return nil, s.log.Error(models.ErrGetDataInDB, err)
	}

	return documents, nil
}

func (s *SQLStorage) GetResolutoins(id int) ([]models.Resolution, error) {

	rows, err := s.db.Query("SELECT * FROM `res` WHERE `doc_id` = ?", id)

	if err != nil {
		return nil, s.log.Error(models.ErrGetDataInDB, err)
	}

	defer rows.Close()

	var resolutions []models.Resolution

	for rows.Next() {
		var resolution models.Resolution

		if err := rows.Scan(&resolution.ID, &resolution.DocID, &resolution.Ispolnitel, &resolution.Text, &resolution.Time, &resolution.Date, &resolution.User, &resolution.Creator); err != nil {
			return nil, s.log.Error(models.ErrGetDataInDB, err)
		}

		resolutions = append(resolutions, resolution)
	}

	if err := rows.Err(); err != nil {
		return nil, s.log.Error(models.ErrGetDataInDB, err)
	}

	return resolutions, nil
}
