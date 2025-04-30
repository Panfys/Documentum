package service

import (
	"documentum/pkg/models"
	"errors"
	"fmt"
	"path/filepath"
	
)

func (d *docService) AddLookDocument(id int, login string) error {

	name, err := d.stor.GetUserName(login)
	if err != nil {
		return err
	}

	err = d.stor.AddLookDocument(id, name)

	if err != nil {
		return fmt.Errorf("ошибка записи просмотра документа в БД")
	}

	return nil
}

func (d *docService) AddIngoingDoc(reqDoc models.Document) (models.Document, error) {

	doc, err := d.valid.ValidIngoingDoc(reqDoc) 
	if err != nil {
		return models.Document{}, err
	}

	id, err := d.stor.GetAutoIncrement("doc")

	if err != nil {
		return models.Document{}, errors.New("ошибка получения автоинкремента")
	}

	doc.ID = id

	for i := range doc.Resolutions {

		res, err := d.valid.ValidResolution(doc.Resolutions[i])
		if err != nil {
			return models.Document{}, err
		}

		res.DocID = id
		res.Creator = doc.Creator
		doc.Resolutions[i] = &res

		if res.Result != "" {
			doc.Result = doc.Result
		}
	}

	path := "/app/web/source/documents/"

	newFileName, err := models.GenerateUniqueFilename(path, doc.FileHeader.Filename)
	if err != nil {
		return doc, err
	}

	filePath := filepath.Join(path, newFileName)
	if err := models.SaveFile(doc.File, filePath); err != nil {
		return doc, err
	}

	doc.FileURL = filepath.Join("/source/documents/", newFileName)
	/*
		if err := d.storage.UpdateUserIcon(storagePath, login); err != nil {
			os.Remove(filePath) // Откатываем изменения если ошибка
			return cleanDoc, err
		}
	*/
	return doc, nil
}
