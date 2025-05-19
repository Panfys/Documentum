package document

import (
	"documentum/pkg/models"
	"fmt"
	"encoding/json"
)

func (d *docService) UpdateDocFamiliar(types, id, login string) error {

	name, err := d.stor.GetUserName(login)
	if err != nil {
		return err
	}

	res, err := d.stor.UpdateDocFamiliar(types, id, name)

	if err != nil {
		return err
	}

	if res == 1 {
		var (
			message models.Message
			content models.UpdDocFamConten
		)
		message.Action = "updDocFam"
		content.Type = types
		content.DocID = id
		content.Familiar = fmt.Sprintf("<br>%s", name)
		jsonContent, err := json.Marshal(content)
		if err != nil {
			return err
		}
		message.Content = jsonContent
		d.wsSrv.Broadcast(message)
	}

	return nil
}

func (d *docService) UpdateDocument(reqDoc models.Document) (models.Document, error) {

	doc, err := d.validSrv.ValidUpdateDocument(reqDoc)
	if err != nil {
		return models.Document{}, err
	}

	var result string

	for i := range doc.Resolutions {
		res, err := d.validSrv.ValidResolution(&doc.Resolutions[i])
		if err != nil {
			return models.Document{}, err
		}

		res.Creator += doc.Creator
		doc.Resolutions[i] = res

		if res.Result != "" {
			if result == "" {
				result = res.Result
			} else {
				result += " <br> " + res.Result
			}
		}
	}

	if result != "" {
		doc.Result = result
	}

	/*
		path := "/app/web/source/documents/"

		newFileName, err := d.fileSrv.AddFile(path, doc.FileHeader.Filename, doc.File)

		if err != nil {
			return models.Document{}, err
		}

		doc.FileURL = filepath.Join("/source/documents/", newFileName)
	*/

	if err := d.stor.UpdateDocumentWithResolutions(doc); err != nil {
		//d.fileSrv.DeleteFileIfExists(filepath.Join(path, newFileName))
		return models.Document{}, err
	}

	return doc, nil
}
