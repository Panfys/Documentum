package document

import (
	"documentum/pkg/models"
	"encoding/json"
	"fmt"
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

	// Отправляем в ws для обновления у клиента
	if res == 1 {
		var (
			message models.Message
			content models.UpdDocFamConten
		)
		message.Action = "updDocFam"
		content.Type = types
		content.DocID = id
		content.Familiar = fmt.Sprintf(", <br>%s", name)
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

	var result, ispolnitel string

	for i := range doc.Resolutions {
		res, err := d.validSrv.ValidResolution(&doc.Resolutions[i])
		if err != nil {
			return models.Document{}, err
		}
		doc.Resolutions[i] = res

		if doc.Resolutions[i].Result != "" {
			if result == "" {
				result = doc.Resolutions[i].Result
			} else {
				result += " <br> " + doc.Resolutions[i].Result
			}
		}
	}

	if result != "" {
		doc.Result = result
	}
	
	if err := d.stor.UpdateDocumentWithResolutions(doc); err != nil {
		//d.fileSrv.DeleteFileIfExists(filepath.Join(path, newFileName))
		return models.Document{}, err
	}

	for i := range doc.Resolutions { 
		if doc.Resolutions[i].Deadline.Valid {
			doc.Resolutions[i].DeadlineStr, err = d.parseTime(doc.Resolutions[i].Deadline.String)
			if err != nil {
				return models.Document{}, err
			}
		}

		doc.Resolutions[i].Date, err = d.parseResolutionDate(doc.Resolutions[i].Date)
		if err != nil {
			return models.Document{}, err
		}

		ispolnitel = fmt.Sprintf("<div class='table__ispolnitel--ispolnitel'>%s</div>"+
			"<div class='table__ispolnitel--text'>&#171%s&#187</div>"+
			"<div class='table__ispolnitel--user'>%s</div>",
			doc.Resolutions[i].Ispolnitel, doc.Resolutions[i].Text, doc.Resolutions[i].User)
		doc.Resolutions[i].Creator = doc.Creator

	}

	if ispolnitel != "" {
		doc.Ispolnitel = ispolnitel
	}

	/*
		path := "/app/web/source/documents/"

		newFileName, err := d.fileSrv.AddFile(path, doc.FileHeader.Filename, doc.File)

		if err != nil {
			return models.Document{}, err
		}

		doc.FileURL = filepath.Join("/source/documents/", newFileName)
	*/

	// Отправляем в ws для обновления у клиента
	var (
		message models.Message
	)
	message.Action = "updDocRes"
	jsonContent, err := json.Marshal(doc)
	if err != nil {
		return doc, err
	}
	message.Content = jsonContent
	d.wsSrv.Broadcast(message)

	return doc, nil
}
