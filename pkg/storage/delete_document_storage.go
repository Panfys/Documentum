package storage

import (
	"errors"
)

func (d *SQLStorage) DeleteDocumentByID(id int) error {
	query := "DELETE FROM doc WHERE id = ?"

	_, err := d.db.Exec(query, id)
	if err != nil {
		return errors.New("ошибка удаления документа из БД")
	}
	
	return nil
}

func (d *SQLStorage) DeleteResolutionByDocID(id int) error {
	query := "DELETE FROM res WHERE doc_id = ?"

	_, err := d.db.Exec(query, id)
	if err != nil {
		return errors.New("ошибка удаления резольции документа из БД")
	}
	
	return nil
}