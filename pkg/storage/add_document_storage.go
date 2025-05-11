package storage

import (
	"documentum/pkg/models"
	"time"
)

func (s *SQLStorage) AddLookDocument(id string, name string) error {
	username := "<br>" + name

	_, err := s.db.Exec("UPDATE `documents` SET `familiar` = IF(`familiar` LIKE ?, `familiar`, CONCAT(`familiar`, ?)) WHERE `id` = ?", "%"+name+"%", username, id)
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
	insertDocQuery := "INSERT INTO documents (type, fnum, fdate, lnum, ldate, name, sender, ispolnitel, result, familiar, count, copy, width, location, file, creator, createdAt) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"

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
		insertResQuery := "INSERT INTO resolutions (type, doc_id, ispolnitel, text, deadline, date, user, creator, createdAt) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)"

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
	insertDocQuery := "INSERT INTO documents (type, fnum, fdate, lnum, ldate, name, sender, ispolnitel, result, familiar, count, copy, width, location, file, creator, createdAt) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"

	_, err := s.db.Exec(insertDocQuery, doc.Type, doc.FNum, doc.FDate, doc.LNum, doc.LDate, doc.Name, doc.Sender, doc.Ispolnitel, doc.Result, doc.Familiar, doc.Count, doc.Copy, doc.Width, doc.Location, doc.FileURL, doc.Creator, doc.CreatedAt)

	if err != nil {
		return s.log.Error(models.ErrAddDataInDB, err)
	}

	return nil
}

func (s *SQLStorage) AddDirective(doc models.Directive) error {
	insertDirQuery := "INSERT INTO directives (number, date, name, autor, numCoverLetter, dateCoverLetter, countCopy, sender, numSendLetter, dateSendLetter, countSendCopy, familiar, location, fileURL, creator, createdAt) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"

	_, err := s.db.Exec(insertDirQuery, doc.Number, doc.Date, doc.Name, doc.Autor, doc.NumCoverLetter, doc.DateCoverLetter, doc.CountCopy, doc.Sender, doc.NumSendLetter, doc.DateSendLetter, doc.CountSendCopy, doc.Familiar, doc.Location, doc.FileURL, doc.Creator, time.Now())

	if err != nil {
		return s.log.Error(models.ErrAddDataInDB, err)
	}

	return nil
}

func (s *SQLStorage) AddResolution(res models.Resolution) error {
	newRes := "INSERT INTO `resolutions` SET `doc_id` = ?, `ispolnitel` = ?, `text` = ?, `time` = ?, `date` = ?, `user` = ?, `creator` = ?"

	_, err := s.db.Exec(newRes, res.DocID, res.Ispolnitel, res.Text, res.Deadline, res.Date, res.User, res.Creator)

	if err != nil {
		return s.log.Error(models.ErrAddDataInDB, err)
	}

	return nil
}
