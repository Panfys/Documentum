package storage

import (
	"documentum/pkg/models"
	"fmt"
	"time"
)

func (s *SQLStorage) AddFamiliarDocument(table, id, name string) error {

	query := fmt.Sprintf("UPDATE %s SET familiar = IF(familiar IS NULL OR familiar = '', ?, CONCAT(familiar, ', <br> ', ?)) WHERE id = ? AND (familiar IS NULL OR familiar NOT LIKE ?)", table)
	
	_, err := s.db.Exec(query, name, name, id, "%"+name+"%")

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
	insertDocQuery := "INSERT INTO inouts (type, fnum, fdate, lnum, ldate, name, sender, ispolnitel, result, familiar, count, copy, width, location, file, creator, createdAt) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"

	result, err := tx.Exec(insertDocQuery, doc.Type, doc.FNum, doc.FDate, doc.LNum, doc.LDate,
		doc.Name, doc.Sender, doc.Ispolnitel, doc.Result,
		doc.Familiar, doc.Count, doc.Copy,
		doc.Width, doc.Location,
		doc.FileURL, doc.Creator, time.Now())
	if err != nil {
		tx.Rollback()
		return s.log.Error(models.ErrAddDataInDB, err, " Запрос: ", insertDocQuery)
	}

	docID, err = result.LastInsertId()
	if err != nil {
		tx.Rollback()
		return s.log.Error(models.ErrAddDataInDB, err)
	}

	for _, res := range doc.Resolutions {
		insertResQuery := "INSERT INTO inouts (type, doc_id, ispolnitel, text, deadline, date, user, creator, createdAt) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)"

		if _, err := tx.Exec(insertResQuery,
			res.Type,
			docID,
			res.Ispolnitel,
			res.Text,
			res.Deadline,
			res.Date,
			res.User,
			res.Creator,
			time.Now(),
		); err != nil {
			tx.Rollback()
			return s.log.Error(models.ErrAddDataInDB, err, " Запрос: ", insertResQuery)
		}
	}

	if err = tx.Commit(); err != nil {
		s.log.Error(models.ErrAddDataInDB, err)
	}

	return nil
}

func (s *SQLStorage) AddDocument(doc models.Document) error {
	insertDocQuery := "INSERT INTO inouts (type, fnum, fdate, lnum, ldate, name, sender, ispolnitel, result, familiar, count, copy, width, location, file, creator, createdAt) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"

	_, err := s.db.Exec(insertDocQuery, doc.Type, doc.FNum, doc.FDate, doc.LNum, doc.LDate, doc.Name, doc.Sender, doc.Ispolnitel, doc.Result, doc.Familiar, doc.Count, doc.Copy, doc.Width, doc.Location, doc.FileURL, doc.Creator, doc.CreatedAt)

	if err != nil {
		return s.log.Error(models.ErrAddDataInDB, err, " Запрос: ", insertDocQuery)
	}

	return nil
}

func (s *SQLStorage) AddDirective(doc models.Directive) error {
	insertDirQuery := "INSERT INTO directives (number, date, name, autor, numCoverLetter, dateCoverLetter, countCopy, sender, numSendLetter, dateSendLetter, countSendCopy, familiar, location, fileURL, creator, createdAt) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"

	_, err := s.db.Exec(insertDirQuery, doc.Number, doc.Date, doc.Name, doc.Autor, doc.NumCoverLetter, doc.DateCoverLetter, doc.CountCopy, doc.Sender, doc.NumSendLetter, doc.DateSendLetter, doc.CountSendCopy, doc.Familiar, doc.Location, doc.FileURL, doc.Creator, time.Now())

	if err != nil {
		return s.log.Error(models.ErrAddDataInDB, err, " Запрос: ", insertDirQuery)
	}

	return nil
}

func (s *SQLStorage) AddInventory(doc models.Inventory) error {
	insertInvQuery := "INSERT INTO inventory (number, numCoverLetter, dateCoverLetter, name, sender, countCopy, copy, addressee, numSendLetter, dateSendLetter, sendCopy, familiar, location, fileURL, creator, createdAt) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"

	_, err := s.db.Exec(insertInvQuery, doc.Number, doc.NumCoverLetter, doc.DateCoverLetter, doc.Name, doc.Sender, doc.CountCopy, doc.Copy, doc.Addressee, doc.NumSendLetter, doc.DateSendLetter, doc.SendCopy, doc.Familiar, doc.Location, doc.FileURL, doc.Creator, time.Now())

	if err != nil {
		return s.log.Error(models.ErrAddDataInDB, err, " Запрос: ", insertInvQuery)
	}

	return nil
}

func (s *SQLStorage) AddResolution(res models.Resolution) error {
	newRes := "INSERT INTO `resolutions` SET `doc_id` = ?, `ispolnitel` = ?, `text` = ?, `time` = ?, `date` = ?, `user` = ?, `creator` = ?"

	_, err := s.db.Exec(newRes, res.DocID, res.Ispolnitel, res.Text, res.Deadline, res.Date, res.User, res.Creator)

	if err != nil {
		return s.log.Error(models.ErrAddDataInDB, err, " Запрос: ", newRes)
	}

	return nil
}
