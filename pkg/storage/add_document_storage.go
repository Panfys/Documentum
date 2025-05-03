package storage

import (
	"documentum/pkg/models"
	"errors"
	"fmt"
)

func (s *SQLStorage) AddLookDocument(id string, name string) error {
	username := "<br>" + name

	_, err := s.db.Exec("UPDATE `doc` SET `familiar` = IF(`familiar` LIKE ?, `familiar`, CONCAT(`familiar`, ?)) WHERE `id` = ?", "%"+name+"%", username, id)
	if err != nil {
		return s.log.Error(models.ErrAddDataInDB, err)
	}
	return nil
}

func (s *SQLStorage) AddDocumentWithResolutions(doc models.Document) error {

	tx, err := s.db.Begin()
	if err != nil {
		return s.log.Error(models.ErrAddDataInDB, err)
	}

	var docID int64
	insertDocQuery := "INSERT INTO doc (type, fnum, fdate, lnum, ldate, name, sender, ispolnitel, result, familiar, count, copy, width, location, file, creator) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"

	result, err := tx.Exec(insertDocQuery, doc.Type, doc.FNum, doc.FDate, doc.LNum, doc.LDate,
		doc.Name, doc.Sender, doc.Ispolnitel, doc.Result,
		doc.Familiar, doc.Count, doc.Copy,
		doc.Width, doc.Location,
		doc.FileURL, doc.Creator)
	if err != nil {
		tx.Rollback()
		return s.log.Error(models.ErrAddDataInDB, err)
	}

	docID, err = result.LastInsertId()
	if err != nil {
		tx.Rollback()
		return s.log.Error(models.ErrAddDataInDB, err)
	}

	for _, res := range doc.Resolutions {
		insertResQuery := "INSERT INTO res (doc_id, ispolnitel, text, time, date, user, creator) VALUES (?, ?, ?, ?, ?, ?, ?)"

		if _, err := tx.Exec(insertResQuery,
			docID,
			res.Ispolnitel,
			res.Text,
			res.Time,
			res.Date,
			res.User,
			res.Creator); err != nil {
			tx.Rollback()
			return s.log.Error(models.ErrAddDataInDB, err)
		}
	}

	if err = tx.Commit(); err != nil {
		s.log.Error(models.ErrAddDataInDB, err)
	}

	return nil
}

func (d *SQLStorage) AddDocument(doc models.Document) error {
	insertDocQuery := "INSERT INTO doc (type, fnum, fdate, lnum, ldate, name, sender, ispolnitel, result, familiar, count, copy, width, location, file, creator) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"

	_, err := d.db.Exec(insertDocQuery, doc.Type, doc.FNum, doc.FDate, doc.LNum, doc.LDate, doc.Name, doc.Sender, doc.Ispolnitel, doc.Result, doc.Familiar, doc.Count, doc.Copy, doc.Width, doc.Location, doc.FileURL, doc.Creator)

	if err != nil {
		return errors.New("ошибка записи документа в БД")
	}

	return nil
}

func (d *SQLStorage) AddResolution(res models.Resolution) error {
	newRes := "INSERT INTO `res` SET `doc_id` = ?, `ispolnitel` = ?, `text` = ?, `time` = ?, `date` = ?, `user` = ?, `creator` = ?"

	_, err := d.db.Exec(newRes, res.DocID, res.Ispolnitel, res.Text, res.Time, res.Date, res.User, res.Creator)

	if err != nil {
		return fmt.Errorf("ошибка записи резолюции в БД: %s", err)
	}

	return nil
}
