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
		content.Familiar = name
		jsonContent, err := json.Marshal(content)
		if err != nil {
			return err
		}
		message.Content = jsonContent
		d.wsSrv.Broadcast(message)
	}

	return nil
}

func (s *docService) UpdateDocument(reqDoc models.Document) error {
    // Валидация документа
    doc, err := s.validSrv.ValidUpdateDocument(reqDoc)
    if err != nil {
        return err
    }

    // Валидация и обработка резолюций
    if err := s.processDocumentResolutions(&doc); err != nil {
        return err
    }

    // Обновление документа в хранилище
    if err := s.stor.UpdateDocumentWithResolutions(doc); err != nil {
        return err
    }

    // Подготовка данных для отображения
    if err := s.prepareDocumentForResponse(&doc); err != nil {
        return err
    }

    // Отправка обновления через WebSocket
    return s.sendDocumentUpdateWS(&doc)
}

func (s *docService) processDocumentResolutions(doc *models.Document) error {
    var resultBuilder string

    for i := range doc.Resolutions {
        // Валидация резолюции
        res, err := s.validSrv.ValidResolution(&doc.Resolutions[i])
        if err != nil {
            return err
        }
        doc.Resolutions[i] = res

        // Сборка результата исполнения
        if res.Result != "" {
            if resultBuilder == "" {
                resultBuilder = res.Result
            } else {
                resultBuilder += " <br> " + res.Result
            }
        }
    }

    if resultBuilder != "" {
        doc.Result = resultBuilder
    }

    return nil
}

func (s *docService) prepareDocumentForResponse(doc *models.Document) error {
    var lastIspolnitel string

    for i := range doc.Resolutions {
        // Подготовка дат резолюций
        if err := s.prepareSingleResolution(&doc.Resolutions[i]); err != nil {
            return err
        }

        // Формирование информации об исполнителе для последней резолюции
        if i == len(doc.Resolutions)-1 {
            lastIspolnitel = fmt.Sprintf(
                "<div class='table__ispolnitel--ispolnitel'>%s</div>"+
                    "<div class='table__ispolnitel--text'>&#171%s&#187</div>"+
                    "<div class='table__ispolnitel--user'>%s</div>",
                doc.Resolutions[i].Ispolnitel,
                doc.Resolutions[i].Text,
                doc.Resolutions[i].User,
            )
            doc.Resolutions[i].Creator = doc.Creator
        }
    }

    if lastIspolnitel != "" {
        doc.Ispolnitel = lastIspolnitel
    }

    return nil
}

func (s *docService) sendDocumentUpdateWS(doc *models.Document) error {
    message := models.Message{
        Action: "updDocRes",
    }

    jsonContent, err := json.Marshal(doc)
    if err != nil {
        return err
    }

    message.Content = jsonContent
    s.wsSrv.Broadcast(message)

    return nil
}