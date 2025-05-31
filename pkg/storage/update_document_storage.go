package storage

import (
	"documentum/pkg/models"
	"time"
)

func (s *SQLStorage) UpdateDocFamiliar(types, id, name string) (int64, error) {
    var (
        table     string
        familiar  string
        familiars []string
    )

    // Определяем таблицу
    if types == "ingoing" || types == "outgoing" {
        table = "inouts"
    } else {
        table = types
    }

    // 1. Получаем все существующие записи familiar для данного документа
    query := "SELECT docFamiliar FROM familiars WHERE docTable = ? AND docID = ?"
    rows, err := s.db.Query(query, table, id)
    if err != nil {
        return 0, s.log.Error(models.ErrGetDataInDB, err)
    }
    defer rows.Close()

    // Собираем все существующие familiars
    for rows.Next() {
        if err := rows.Scan(&familiar); err != nil {
            return 0, s.log.Error(models.ErrGetDataInDB, err)
        }
        familiars = append(familiars, familiar)
    }

    // Проверяем, есть ли ошибка после итерации
    if err := rows.Err(); err != nil {
        return 0, s.log.Error(models.ErrGetDataInDB, err)
    }

    // 2. Проверяем, существует ли уже такое имя
    nameExists := false
    for _, f := range familiars {
        if f == name {
            nameExists = true
            break
        }
    }

    // 3. Если имя не найдено, добавляем новую запись
    if !nameExists {
        insertQuery := "INSERT INTO familiars (docTable, docID, docFamiliar, createdAt) VALUES (?, ?, ?, ?)"
        result, err := s.db.Exec(insertQuery, table, id, name, time.Now()) 
        if err != nil {
            return 0, s.log.Error(models.ErrAddDataInDB, err)
        }

        // Возвращаем ID новой записи
        return result.RowsAffected()
    }

    // Если имя уже существует, возвращаем 0 и nil ошибку
    return 0, nil
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
