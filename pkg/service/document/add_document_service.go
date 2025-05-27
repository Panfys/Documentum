package document

import (
	"documentum/pkg/models"
	"path/filepath"
)

func (d *docService) AddDocument(reqDoc models.Document) (models.Document, error) {

	doc, err := d.validSrv.ValidDocument(reqDoc)
	if err != nil {
		return models.Document{}, err
	}

	for i := range doc.Resolutions {
		res, err := d.validSrv.ValidResolution(&doc.Resolutions[i])
		if err != nil {
			return models.Document{}, err
		}

		res.Creator += doc.Creator
		doc.Resolutions[i] = res
	}

	path := "/app/web/source/documents/"

	newFileName, err := d.fileSrv.AddFile(path, doc.FileHeader.Filename, doc.File)

	if err != nil {
		return models.Document{}, err
	}

	doc.FileURL = filepath.Join("/source/documents/", newFileName)

	if err := d.stor.AddDocumentWithResolutions(doc); err != nil {
		d.fileSrv.DeleteFileIfExists(filepath.Join(path, newFileName))
		return models.Document{}, err
	}

	return doc, nil
}

func (d *docService) AddDirective(reqDir models.Directive) (models.Directive, error) {

	dir, err := d.validSrv.ValidDirective(reqDir)
	if err != nil {
		return models.Directive{}, err
	}

	path := "/app/web/source/documents/"

	newFileName, err := d.fileSrv.AddFile(path, dir.FileHeader.Filename, dir.File)

	if err != nil {
		return models.Directive{}, err
	}

	dir.FileURL = filepath.Join("/source/documents/", newFileName)

	if err := d.stor.AddDirective(dir); err != nil {
		d.fileSrv.DeleteFileIfExists(filepath.Join(path, newFileName))
		return models.Directive{}, err
	}

	return dir, nil
}

func (d *docService) AddInventory(reqInv models.Inventory) (models.Inventory, error) {

	inv, err := d.validSrv.ValidInventory(reqInv)
	if err != nil {
		return models.Inventory{}, err
	}
	
	path := "/app/web/source/documents/"

	newFileName, err := d.fileSrv.AddFile(path, inv.FileHeader.Filename, inv.File)

	if err != nil {
		return models.Inventory{}, err
	}

	inv.FileURL = filepath.Join("/source/documents/", newFileName)

	if err := d.stor.AddInventory(inv); err != nil {
		d.fileSrv.DeleteFileIfExists(filepath.Join(path, newFileName))
		return models.Inventory{}, err
	}

	return inv, nil
}
