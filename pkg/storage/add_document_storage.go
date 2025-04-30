package storage

import (
	"documentum/pkg/models"
)

func (d *SQLStorage) AddLookDocument(id int, name string) error {
	username := "<br>" + name

	_, err := d.db.Exec("UPDATE `doc` SET `familiar` = IF(`familiar` LIKE ?, `familiar`, CONCAT(`familiar`, ?)) WHERE `id` = ?", "%"+name+"%", username, id)
	return err
}

func (d *SQLStorage) AddDocument(doc models.Document) error {
	return nil
}
