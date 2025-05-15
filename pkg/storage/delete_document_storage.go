package storage

import (
	"documentum/pkg/models"
)

func (s *SQLStorage) DeleteDocumentByID(id int) error {
	
	query := "DELETE FROM inouts WHERE id = ?"

	_, err := s.db.Exec(query, id)
	if err != nil {
		return s.log.Error(models.ErrGetDataInDB, err)
	}
	
	return nil
}

func (s *SQLStorage) DeleteResolutionByDocID(id int) error {
	query := "DELETE FROM resolutions WHERE doc_id = ?"

	_, err := s.db.Exec(query, id)
	if err != nil {
		return s.log.Error(models.ErrGetDataInDB, err)
	}
	
	return nil
}