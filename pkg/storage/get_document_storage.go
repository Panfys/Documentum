package storage

import (
	"documentum/pkg/models"
	"fmt"
	"strings"
)

func (s *SQLStorage) GetDocuments(settings models.DocSettings) ([]models.Document, error) {
    query := fmt.Sprintf("SELECT * FROM `inouts` WHERE `type` = ? AND `fdate` BETWEEN ? AND ? ORDER BY %s %s", 
        settings.DocCol, settings.DocSet)

    rows, err := s.db.Query(query, settings.DocType, settings.DocDatain, settings.DocDatato)
    if err != nil {
        return nil, s.log.Error(models.ErrGetDataInDB, err, " Запрос: ", query)
    }
    defer rows.Close()

    var documents []models.Document
    var docIDs []int

    for rows.Next() {
        var document models.Document
        if err := rows.Scan(&document.ID, &document.Type, &document.FNum, &document.FDate, 
            &document.LNum, &document.LDate, &document.Name, &document.Sender, 
            &document.Ispolnitel, &document.Result, &document.Count, &document.Copy, 
            &document.Width, &document.Location, &document.FileURL, 
            &document.Creator, &document.CreatedAt); err != nil {
            return nil, s.log.Error(models.ErrGetDataInDB, err)
        }
        documents = append(documents, document)
        docIDs = append(docIDs, document.ID)
    }

    if err := rows.Err(); err != nil {
        return nil, s.log.Error(models.ErrGetDataInDB, err)
    }

    // Получаем familiars для всех документов
    familiarsMap, err := s.getFamiliars("inouts", docIDs)
    if err != nil {
        return nil, err
    }

    // Наполняем документы familiars
    for i := range documents {
        documents[i].Familiars = familiarsMap[documents[i].ID]
    }
    
    return documents, nil
}

func (s *SQLStorage) GetDirectives(settings models.DocSettings) ([]models.Directive, error) {
    query := fmt.Sprintf("SELECT * FROM `directives` WHERE `date` BETWEEN ? AND ? ORDER BY %s %s", 
        settings.DocCol, settings.DocSet)

    rows, err := s.db.Query(query, settings.DocDatain, settings.DocDatato)
    if err != nil {
        return nil, s.log.Error(models.ErrGetDataInDB, err, " Запрос: ", query)
    }
    defer rows.Close()

    var directives []models.Directive
    var docIDs []int

    for rows.Next() {
        var directive models.Directive
        if err := rows.Scan(&directive.ID, &directive.Number, &directive.Date, &directive.Name, 
            &directive.Autor, &directive.NumCoverLetter, &directive.DateCoverLetter, 
            &directive.CountCopy, &directive.Sender, &directive.NumSendLetter, 
            &directive.DateSendLetter, &directive.CountSendCopy, &directive.Location, 
            &directive.FileURL, &directive.Creator, &directive.CreatedAt); err != nil {
            return nil, s.log.Error(models.ErrGetDataInDB, err)
        }
        directives = append(directives, directive)
        docIDs = append(docIDs, directive.ID)
    }

    if err := rows.Err(); err != nil {
        return nil, s.log.Error(models.ErrGetDataInDB, err)
    }

    // Получаем familiars для всех директив
    familiarsMap, err := s.getFamiliars("directives", docIDs)
    if err != nil {
        return nil, err
    }

    // Наполняем директивы familiars
    for i := range directives {
        directives[i].Familiars = familiarsMap[directives[i].ID]
    }
 
    return directives, nil
}

func (s *SQLStorage) GetInventory(settings models.DocSettings) ([]models.Inventory, error) {
    query := fmt.Sprintf("SELECT * FROM `inventory` ORDER BY %s %s", settings.DocCol, settings.DocSet)

    rows, err := s.db.Query(query)
    if err != nil {
        return nil, s.log.Error(models.ErrGetDataInDB, err, " Запрос: ", query)
    }
    defer rows.Close()

    var inventoryDocs []models.Inventory
    var docIDs []int

    for rows.Next() {
        var inventoryDoc models.Inventory
        if err := rows.Scan(
            &inventoryDoc.ID,
            &inventoryDoc.Number,
            &inventoryDoc.NumCoverLetter,
            &inventoryDoc.DateCoverLetter,
            &inventoryDoc.Name,
            &inventoryDoc.Sender,
            &inventoryDoc.CountCopy,
            &inventoryDoc.Copy,
            &inventoryDoc.Addressee,
            &inventoryDoc.NumSendLetter,
            &inventoryDoc.DateSendLetter,
            &inventoryDoc.SendCopy,
            &inventoryDoc.Location,
            &inventoryDoc.FileURL,
            &inventoryDoc.Creator,
            &inventoryDoc.CreatedAt,
        ); err != nil {
            return nil, s.log.Error(models.ErrGetDataInDB, err)
        }
        inventoryDocs = append(inventoryDocs, inventoryDoc)
        docIDs = append(docIDs, inventoryDoc.ID)
    }

    if err := rows.Err(); err != nil {
        return nil, s.log.Error(models.ErrGetDataInDB, err)
    }

    // Получаем familiars для всех инвентарных документов
    familiarsMap, err := s.getFamiliars("inventory", docIDs)
    if err != nil {
        return nil, err
    }

    // Наполняем документы familiars
    for i := range inventoryDocs {
        inventoryDocs[i].Familiars = familiarsMap[inventoryDocs[i].ID]
    }

    return inventoryDocs, nil
}

func (s *SQLStorage) GetResolutoins(id int64) ([]models.Resolution, error) {

	rows, err := s.db.Query("SELECT * FROM `resolutions` WHERE `doc_id` = ?", id)

	if err != nil {
		return nil, s.log.Error(models.ErrGetDataInDB, err, " Запрос: ", "SELECT * FROM `resolutions`...")
	}

	defer rows.Close()

	var resolutions []models.Resolution

	for rows.Next() {
		var resolution models.Resolution

		if err := rows.Scan(&resolution.ID, &resolution.Type, &resolution.DocID, &resolution.Ispolnitel, &resolution.Text, &resolution.Result, &resolution.Deadline, &resolution.Date, &resolution.User, &resolution.Creator, &resolution.CreatedAt); err != nil {
			return nil, s.log.Error(models.ErrGetDataInDB, err)
		}

		resolutions = append(resolutions, resolution)
	}

	if err := rows.Err(); err != nil {
		return nil, s.log.Error(models.ErrGetDataInDB, err)
	}

	return resolutions, nil
}

func (s *SQLStorage) getFamiliars(table string, docIDs []int) (map[int][]string, error) {
    if len(docIDs) == 0 {
        return make(map[int][]string), nil
    }

    // Создаем список параметров для IN
    params := make([]any, len(docIDs)+1)
    params[0] = table
    for i, id := range docIDs {
        params[i+1] = id
    }

    // Создаем строку с плейсхолдерами
    placeholders := "?" + strings.Repeat(",?", len(docIDs)-1)

    query := fmt.Sprintf("SELECT docID, docFamiliar FROM familiars WHERE docTable = ? AND docID IN (%s)", placeholders)
    
    rows, err := s.db.Query(query, params...)
    if err != nil {
        return nil, s.log.Error(models.ErrGetDataInDB, err, " Запрос familiars: ", query)
    }
    defer rows.Close()

    familiarsMap := make(map[int][]string)
    for rows.Next() {
        var docID int
        var familiar string
        if err := rows.Scan(&docID, &familiar); err != nil {
            return nil, s.log.Error(models.ErrGetDataInDB, err)
        }
        familiarsMap[docID] = append(familiarsMap[docID], familiar)
    }

    if err := rows.Err(); err != nil {
        return nil, s.log.Error(models.ErrGetDataInDB, err)
    }

    return familiarsMap, nil
}