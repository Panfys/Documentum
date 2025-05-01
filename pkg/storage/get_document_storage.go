package storage

import (
	"documentum/pkg/models"
	"fmt"
)

func (d *SQLStorage) GetDocuments(settings models.DocSettings) ([]models.Document, error) {

	query := fmt.Sprintf("SELECT * FROM `doc` WHERE `type` = ? AND `fdate` BETWEEN ? AND ? ORDER BY %s %s", settings.DocCol, settings.DocSet)

	rows, err := d.db.Query(query, settings.DocType, settings.DocDatain, settings.DocDatato)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var (
		documents []models.Document
	)

	for rows.Next() {
		var document models.Document


		if err := rows.Scan(&document.ID, &document.Type, &document.FNum, &document.FDate, &document.LNum, &document.LDate, &document.Name, &document.Sender, &document.Ispolnitel, &document.Result, &document.Familiar, &document.Count, &document.Copy, &document.Width, &document.Location, &document.FileURL, &document.Creator); err != nil {
			return nil, err
		}

		documents = append(documents, document)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return documents, nil
}

func (d *SQLStorage) GetResolutoins(id int) ([]models.Resolution, error) {

	rows, err := d.db.Query("SELECT * FROM `res` WHERE `doc_id` = ?", id)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var resolutions []models.Resolution

	for rows.Next() {
		var resolution models.Resolution

		if err := rows.Scan(&resolution.ID, &resolution.DocID, &resolution.Ispolnitel, &resolution.Text, &resolution.Time, &resolution.Date, &resolution.User, &resolution.Creator); err != nil {
			return nil, err
		}

		resolutions = append(resolutions, resolution)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return resolutions, nil
}

func (d *SQLStorage) GetAutoIncrement(table string) (int, error) {

	var autoIncrement int

	err := d.db.QueryRow("SELECT AUTO_INCREMENT FROM information_schema.TABLES WHERE TABLE_SCHEMA = 'documentum' AND TABLE_NAME = ?;", table).Scan(&autoIncrement)

	if err != nil {
		return 0, err
	}

	return autoIncrement, nil
}
