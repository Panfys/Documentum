package document

import (
	"documentum/pkg/models"
	"encoding/json"
	"path/filepath"
)

func (d *docService) AddDocument(reqDoc models.Document) error {

	doc, err := d.validSrv.ValidDocument(reqDoc)
	if err != nil {
		return err
	}

	for i := range doc.Resolutions {
		res, err := d.validSrv.ValidResolution(&doc.Resolutions[i])
		if err != nil {
			return err
		}

		res.Creator += doc.Creator
		doc.Resolutions[i] = res
	}

	path := "/app/web/source/documents/"

	newFileName, err := d.fileSrv.AddFile(path, doc.FileHeader.Filename, doc.File)

	if err != nil {
		return err
	}

	doc.FileURL = filepath.Join("/source/documents/", newFileName)

	docID, err := d.stor.AddDocumentWithResolutions(doc)
	if err != nil {
		d.fileSrv.DeleteFileIfExists(filepath.Join(path, newFileName))
		return err
	}
	doc.ID = int(docID)

	var responseDoc *models.Document

	responseDoc, err = d.prepareDocument(&doc)

	if err != nil {
		return err
	}

	message := models.Message{
		Action: "addDoc",
	}

	jsonContent, err := json.Marshal(responseDoc)
	if err != nil {
		return err
	}

	message.Content = jsonContent
	d.wsSrv.Broadcast(message)

	return nil
}

func (d *docService) AddDirective(reqDir models.Directive) error {

	dir, err := d.validSrv.ValidDirective(reqDir)
	if err != nil {
		return err
	}

	path := "/app/web/source/documents/"

	newFileName, err := d.fileSrv.AddFile(path, dir.FileHeader.Filename, dir.File)

	if err != nil {
		return err
	}

	dir.FileURL = filepath.Join("/source/documents/", newFileName)

	if err := d.stor.AddDirective(dir); err != nil {
		d.fileSrv.DeleteFileIfExists(filepath.Join(path, newFileName))
		return err
	}

	return nil
}

func (d *docService) AddInventory(reqInv models.Inventory) error {

	inv, err := d.validSrv.ValidInventory(reqInv)
	if err != nil {
		return err
	}

	path := "/app/web/source/documents/"

	newFileName, err := d.fileSrv.AddFile(path, inv.FileHeader.Filename, inv.File)

	if err != nil {
		return err
	}

	inv.FileURL = filepath.Join("/source/documents/", newFileName)

	if err := d.stor.AddInventory(inv); err != nil {
		d.fileSrv.DeleteFileIfExists(filepath.Join(path, newFileName))
		return err
	}

	return nil
}
