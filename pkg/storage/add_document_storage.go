package storage

import (
	"documentum/pkg/models"
	"time"
)

func (s *SQLStorage) AddDocumentWithResolutions(doc models.Document) (int64, error) {

	tx, err := s.db.Begin()
	if err != nil {
		return 0, s.log.Error(models.ErrAddDataInDB, err)
	}

	var docID int64
	insertDocQuery := "INSERT INTO inouts (type, fnum, fdate, lnum, ldate, name, sender, ispolnitel, result, count, copy, width, location, file, creator, createdAt) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"

	result, err := tx.Exec(insertDocQuery, doc.Type, doc.FNum, doc.FDate, doc.LNum, doc.LDate,
		doc.Name, doc.Sender, doc.Ispolnitel, doc.Result, doc.Count, doc.Copy, doc.Width, doc.Location,
		doc.FileURL, doc.Creator, time.Now())
	if err != nil {
		tx.Rollback()
		return 0, s.log.Error(models.ErrAddDataInDB, err, " Запрос: ", insertDocQuery)
	}

	docID, err = result.LastInsertId()
	if err != nil {
		tx.Rollback()
		return 0, s.log.Error(models.ErrAddDataInDB, err)
	}

	for _, res := range doc.Resolutions {
		insertResQuery := "INSERT INTO resolutions (type, doc_id, ispolnitel, text, result, deadline, date, user, creator, createdAt) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"

		if _, err := tx.Exec(insertResQuery,
			res.Type,
			docID,
			res.Ispolnitel,
			res.Text,
			res.Result,
			res.Deadline,
			res.Date,
			res.User,
			res.Creator,
			time.Now(),
		); err != nil {
			tx.Rollback()
			return 0, s.log.Error(models.ErrAddDataInDB, err, " Запрос: ", insertResQuery)
		}
	}

	if doc.Familiar != "" {
		insertFamQuery := `INSERT INTO familiars 
            (docTable, docID, docFamiliar, createdAt) 
            VALUES (?, ?, ?, ?)`

		if _, err := tx.Exec(insertFamQuery,
			"inouts", docID, doc.Familiar, time.Now(),
		); err != nil {
			return 0, s.log.Error(models.ErrAddDataInDB, err, " Запрос: ", insertFamQuery)
		}
	}

	if err = tx.Commit(); err != nil {
		s.log.Error(models.ErrAddDataInDB, err)
	}

	return docID, nil
}

func (s *SQLStorage) AddDirective(doc models.Directive) (int64, error) {
	tx, err := s.db.Begin()
	if err != nil {
		return 0, s.log.Error(models.ErrAddDataInDB, err)
	}
	
	var docID int64

	insertDirQuery := "INSERT INTO directives (number, date, name, autor, numCoverLetter, dateCoverLetter, countCopy, sender, numSendLetter, dateSendLetter, countSendCopy, location, fileURL, creator, createdAt) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"

	result, err := tx.Exec(insertDirQuery, doc.Number, doc.Date, doc.Name, doc.Autor, doc.NumCoverLetter, doc.DateCoverLetter, doc.CountCopy, doc.Sender, doc.NumSendLetter, doc.DateSendLetter, doc.CountSendCopy, doc.Location, doc.FileURL, doc.Creator, time.Now())

	if err != nil {
		tx.Rollback()
		return 0, s.log.Error(models.ErrAddDataInDB, err, " Запрос: ", insertDirQuery)
	}

	docID, err = result.LastInsertId()
	if err != nil {
		tx.Rollback()
		return 0, s.log.Error(models.ErrAddDataInDB, err)
	}

	if doc.Familiar != "" {
		insertFamQuery := `INSERT INTO familiars 
            (docTable, docID, docFamiliar, createdAt) 
            VALUES (?, ?, ?, ?)`

		if _, err := tx.Exec(insertFamQuery,
			"inouts", docID, doc.Familiar, time.Now(),
		); err != nil {
			return 0, s.log.Error(models.ErrAddDataInDB, err, " Запрос: ", insertFamQuery)
		}
	}
	
	if err = tx.Commit(); err != nil {
		s.log.Error(models.ErrAddDataInDB, err)
	}

	return docID, nil
}

func (s *SQLStorage) AddInventory(doc models.Inventory) (int64, error) {
	tx, err := s.db.Begin()
	if err != nil {
		return 0, s.log.Error(models.ErrAddDataInDB, err)
	}

	var docID int64

	insertInvQuery := "INSERT INTO inventory (number, numCoverLetter, dateCoverLetter, name, sender, countCopy, copy, addressee, numSendLetter, dateSendLetter, sendCopy, location, fileURL, creator, createdAt) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"

	result, err := tx.Exec(insertInvQuery, doc.Number, doc.NumCoverLetter, doc.DateCoverLetter, doc.Name, doc.Sender, doc.CountCopy, doc.Copy, doc.Addressee, doc.NumSendLetter, doc.DateSendLetter, doc.SendCopy, doc.Location, doc.FileURL, doc.Creator, time.Now())

	if err != nil {
		tx.Rollback()
		return 0, s.log.Error(models.ErrAddDataInDB, err, " Запрос: ", insertInvQuery)
	}

	docID, err = result.LastInsertId()
	if err != nil {
		tx.Rollback()
		return 0, s.log.Error(models.ErrAddDataInDB, err)
	}

	if doc.Familiar != "" {
		insertFamQuery := `INSERT INTO familiars 
            (docTable, docID, docFamiliar, createdAt) 
            VALUES (?, ?, ?, ?)`

		if _, err := tx.Exec(insertFamQuery,
			"inouts", docID, doc.Familiar, time.Now(),
		); err != nil {
			return 0, s.log.Error(models.ErrAddDataInDB, err, " Запрос: ", insertFamQuery)
		}
	}
	
	if err = tx.Commit(); err != nil {
		s.log.Error(models.ErrAddDataInDB, err)
	}

	return docID, nil
}

func (s *SQLStorage) AddResolution(res models.Resolution) error {
	newRes := "INSERT INTO `resolutions` SET `doc_id` = ?, `ispolnitel` = ?, `text` = ?, `result` = ?, `time` = ?, `date` = ?, `user` = ?, `creator` = ?, `createdAt` = ?"

	_, err := s.db.Exec(newRes, res.DocID, res.Ispolnitel, res.Text, res.Result, res.Deadline, res.Date, res.User, res.Creator, time.Now())

	if err != nil {
		return s.log.Error(models.ErrAddDataInDB, err, " Запрос: ", newRes)
	}

	return nil 
}
 