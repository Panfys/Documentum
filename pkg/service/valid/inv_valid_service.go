package valid

import (
	"documentum/pkg/models"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
	"strings"
	"unicode"
)

func (v *validatService) ValidInventory(reqInv models.Inventory) (models.Inventory, error) {

	inv := v.sanitizeInventory(&reqInv)

	err := v.validInvNumber(inv.Number)
	if err != nil {
		return models.Inventory{}, err
	}

	err = v.validInvName(inv.Name)
	if err != nil {
		return models.Inventory{}, err
	}

	err = v.validInvSender(inv.Sender)
	if err != nil {
		return models.Inventory{}, err
	}

	invCountCopy, err := strconv.Atoi(inv.CountCopy)
	if err != nil {
		return models.Inventory{}, errors.New("количество экземпляров указано некорректно")
	} 

	err = v.validDocCount(invCountCopy)
	if err != nil {
		return models.Inventory{}, err
	}

	err = v.validDocCopy(inv.Copy)
	if err != nil {
		return models.Inventory{}, err
	}

	inv.DateCoverLetter, err = v.stringToDateNullString(inv.DateCoverLetterStr)
	if err != nil {
		return models.Inventory{}, errors.New("дата сопроводительного письма указана неверно")
	}

	inv.DateSendLetter, err = v.stringToDateNullString(inv.DateSendLetterStr)
	if err != nil {
		return models.Inventory{}, errors.New("дата сопроводительного письма указана неверно")
	}

	err = v.validDocFile(inv.FileHeader)
	if err != nil {
		return models.Inventory{}, err
	}

	return inv, nil
}

func (v *validatService) validInvNumber(num string) error { 
	trimNum := strings.TrimSpace(num)

	if trimNum == "" || trimNum == "Инв. №" || trimNum == "№" {
		return errors.New(`порядковый (инвентарный) номер документа не указан`)
	}

	return nil
}

func (v *validatService) validInvName(name string) error {
	trimName := strings.TrimSpace(name)

	if trimName == "" {
		return errors.New("название издания не указано")
	}

	firstChar := []rune(trimName)[0]
	if !unicode.IsUpper(firstChar) && !unicode.IsDigit(firstChar) {
		return errors.New("название издания должно начинаться с заглавной буквы")
	}

	if len(trimName) > 200 {
		return errors.New(`название издания не должно превышать 200 символов`)
	}

	return nil
}

func (v *validatService) validInvSender(name string) error {
	trimName := strings.TrimSpace(name)

	if trimName == "" {
		return errors.New("отправитель, издатель и год издания не указаны")
	}

	if len(trimName) > 200 {
		return errors.New(`отправитель, издатель и год издания не должны превышать 200 символов`)
	}

	return nil
}

func (v *validatService) sanitizeInventory(doc *models.Inventory) models.Inventory{

	doc.Number = v.policy.Sanitize(doc.Number)
	doc.NumCoverLetter = v.policy.Sanitize(doc.NumCoverLetter)
	doc.DateCoverLetterStr = v.policy.Sanitize(doc.DateCoverLetterStr)
	doc.Name = v.policy.Sanitize(doc.Name)
	doc.Sender = v.policy.Sanitize(doc.Sender)
	doc.CountCopy = v.policy.Sanitize(doc.CountCopy)
	doc.Copy = v.policy.Sanitize(doc.Copy)
	doc.Addressee = v.policy.Sanitize(doc.Addressee)
	doc.NumSendLetter = v.policy.Sanitize(doc.NumSendLetter)
	doc.DateSendLetterStr = v.policy.Sanitize(doc.DateSendLetterStr)
	doc.SendCopy = v.policy.Sanitize(doc.SendCopy)
	doc.Familiar = v.policy.Sanitize(doc.Familiar)
	doc.Location = v.policy.Sanitize(doc.Location)
	doc.Creator = v.policy.Sanitize(doc.Creator)
	doc.FileHeader.Filename = v.policy.Sanitize(doc.FileHeader.Filename)

	return *doc
}
