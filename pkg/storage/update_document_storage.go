package storage

import (
	"documentum/pkg/models"
	"time"
)

func (s *SQLStorage) UpdateDocumentWithResolutions(doc models.Document) error {

	tx, err := s.db.Begin()
	if err != nil {
		return s.log.Error(models.ErrAddDataInDB, err)
	}

	insertDocQuery := "UPDATE `inouts` SET result = IF(result = '', ?, CONCAT(result, ' <br> ', ?)) WHERE id = ?"

	if doc.Result != "" {
		_, err := tx.Exec(insertDocQuery, doc.Result, doc.Result, doc.ID)
		if err != nil {
			tx.Rollback()
			return s.log.Error(models.ErrAddDataInDB, err, " Запрос: ", insertDocQuery)
		}
	}

	for _, res := range doc.Resolutions {
		insertResQuery := "INSERT INTO resolutions (type, doc_id, ispolnitel, text, deadline, date, user, creator, createdAt) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)"

		if _, err := tx.Exec(insertResQuery,
			res.Type,
			doc.ID,
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
