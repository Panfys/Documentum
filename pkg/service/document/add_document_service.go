package document

import (
	"documentum/pkg/models"
	"path/filepath"
)

func (d *docService) AddLookDocument(id int, login string) error {

	name, err := d.stor.GetUserName(login)
	if err != nil {
		return err
	}

	err = d.stor.AddLookDocument(id, name)

	if err != nil {
		return err
	}

	return nil
}

func (d *docService) AddIngoingDoc(reqDoc models.Document) (models.Document, error) {

	doc, err := d.validSrv.ValidIngoingDoc(reqDoc)
	if err != nil {
		return models.Document{}, err
	}

	for i := range doc.Resolutions {

		res, err := d.validSrv.ValidResolution(doc.Resolutions[i])
		if err != nil {
			return models.Document{}, err
		}

		res.Creator = doc.Creator
		doc.Resolutions[i] = &res

		if res.Result != "" {
			doc.Result = res.Result
		}
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
