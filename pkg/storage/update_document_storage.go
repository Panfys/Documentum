package storage

import (
	"documentum/pkg/models"
	"fmt"
	"time"
)

func (s *SQLStorage) UpdateDocFamiliar(types, id, name string) (int64, error) {
	var table string 
	if types == "ingoing" || types == "outgoing" {
		table = "inouts" 
	} else {
		table = types 
	}
	query := fmt.Sprintf("UPDATE %s SET familiar = IF(familiar IS NULL OR familiar = '', ?, CONCAT(familiar, ', <br> ', ?)) WHERE id = ? AND (familiar IS NULL OR familiar NOT LIKE ?)", table)

	result, err := s.db.Exec(query, name, name, id, "%"+name+"%")
	if err != nil {
		return 0, s.log.Error(models.ErrAddDataInDB, err)
	}
	
	res, err := result.RowsAffected()
	if err != nil {
		return 0, s.log.Error(models.ErrAddDataInDB, err)
	}

	return res, nil
}

func (s *SQLStorage) UpdateDocumentWithResolutions(doc models.Document) error {

	tx, err := s.db.Begin()
	if err != nil {
		return s.log.Error(models.ErrAddDataInDB, err)
	}

	for _, res := range doc.Resolutions {
		insertResQuery := "INSERT INTO resolutions (type, doc_id, ispolnitel, text, result, deadline, date, user, creator, createdAt) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"

		if _, err := tx.Exec(insertResQuery,
			res.Type,
			doc.ID,
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
			return s.log.Error(models.ErrAddDataInDB, err, " Запрос: ", insertResQuery)
		}
	}

	if err = tx.Commit(); err != nil {
		s.log.Error(models.ErrAddDataInDB, err)
	}

	return nil
}
