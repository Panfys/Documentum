package service

import (
	"database/sql"
	"documentum/pkg/models"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"strings"
	"unicode"
)

func (d *docService) validIngoingDoc(doc models.Document) error {

	if !d.validDocType(doc.Type) {
		return errors.New("тип документа указан некорректно")
	}

	err := d.validDocFNum(doc.FNum)
	if err != nil {
		return err
	}

	if !d.validDocDate(doc.FDate) {
		return errors.New("дата документа не указана")
	}

	err = d.validDocLdate(doc.LDate, doc.LNum)
	if err != nil {
		return err
	}

	err = d.validDocName(doc.Name)
	if err != nil {
		return err
	}

	err = d.validDocSender(doc.Sender)
	if err != nil {
		return err
	}

	err = d.validDocIspolnitel(doc.Ispolnitel, doc.Resolutions)
	if err != nil {
		return err
	}

	err = d.validDoсCount(doc.Count)
	if err != nil {
		return err
	}

	return nil
}

func (d *docService) validDocType(docType string) bool {

	switch docType {
	case "Входящий", "Исходящий", "Директива", "Инвентарный":
		return true
	default:
		return false
	}

}

func (d *docService) validDocFNum(fnum string) error {
	trimFnum := strings.TrimSpace(fnum)

	if trimFnum == "" {
		return errors.New(`номер документа не указан`)
	}

	if trimFnum == "№" || trimFnum == "Вх. №" {
		return errors.New(`номер документа не должен содержать только "№" или "Вх. №"`)
	}

	if strings.HasPrefix(trimFnum, "Вх. № ") {
		numPart := strings.TrimPrefix(trimFnum, "Вх. № ")
		if !d.containsDigit(numPart) {
			return errors.New(`номер документа должен начинаться либо с "Вх. № ", либо с цифры`)
		}
	} else {
		if len(trimFnum) == 0 || !unicode.IsDigit(rune(trimFnum[0])) {
			return errors.New(` если номер документа не начинается с "Вх. № ", то должен начинаться с цифры`)
		}
	}

	if !d.containsDigit(trimFnum) {
		return errors.New(`номер документа указан некорректно, примеры правильного номера: "Вх. № 1", "Вх. № 317/124дсп", "312/8321", "215/1111дсп"`)
	}

	return nil
}

func (d *docService) validDocLNum(lnum string) error {
	trimFnum := strings.TrimSpace(lnum)

	if trimFnum == "" {
		return errors.New(`номер документа не указан`)
	}

	if trimFnum == "№" || trimFnum == "Исх. №" {
		return errors.New(`номер документа не должен содержать только "№" или "Исх. №"`)
	}

	if strings.HasPrefix(trimFnum, "Исх. № ") {
		numPart := strings.TrimPrefix(trimFnum, "Исх. № ")
		if !d.containsDigit(numPart) {
			return errors.New(`номер документа должен начинаться либо с "Исх. № ", либо с цифры`)
		}
	} else {
		if len(trimFnum) == 0 || !unicode.IsDigit(rune(trimFnum[0])) {
			return errors.New(` если номер документа не начинается с "Исх. № ", то должен начинаться с цифры`)
		}
	}

	if !d.containsDigit(trimFnum) {
		return errors.New(`номер документа указан некорректно, примеры правильного номера: "Исх. № 1", "Исх. № 317/124дсп", "330/8321", "215/1111дсп"`)
	}

	return nil
}

func (d *docService) validDocDate(date sql.NullString) bool {
	return date.Valid
}

func (d *docService) validDocLdate(date sql.NullString, num string) error {
	if num != "" {
		err := d.validDocLNum(num)
		if err != nil {
			return err
		}
		if !date.Valid {
			return errors.New("номер поступившего документа указан, а дата не указана")
		}
	} else {
		if date.Valid {
			return errors.New("дата поступившего документа указана, а номер не указан")
		}
	}

	return nil
}

func (d *docService) validDocName(name string) error {
	trimName := strings.TrimSpace(name)

	if trimName == "" {
		return errors.New("наименование документа не указано")
	}

	firstChar := []rune(trimName)[0]
	if !unicode.IsUpper(firstChar) && !unicode.IsDigit(firstChar) {
		return errors.New("наименование документа должно начинаться с заглавной буквы")
	}

	if len(trimName) > 200 {
		return errors.New(`наименование документа не должно превышать 200 символов`)
	}

	return nil
}

func (d *docService) validDocSender(sender string) error {
	trimSender := strings.TrimSpace(sender)

	if trimSender == "" {
		return errors.New("отправитель документа не указано")
	}

	if len(trimSender) > 100 {
		return errors.New(`отправитель документа не должно превышать 100 символов`)
	}

	return nil
}

func (d *docService) validDocIspolnitel(ispolnitel string, resolutions []*models.Resolution) error {
	trimIspolnitel := strings.TrimSpace(ispolnitel)

	if len(resolutions) == 0 {
		if trimIspolnitel == "" {
			return errors.New("исполнитель документа не указано")
		}

		if len(trimIspolnitel) > 100 {
			return errors.New(`исполнитель документа не должно превышать 100 символов`)
		}
	}
	return nil
}

func (d *docService) validDoсCount(count int) error {

	if count < 1 {
		return errors.New("количестов экземпляров не может быть меньше нуля")
	}

	return nil
}

func (d *docService) containsDigit(s string) bool {
	for _, r := range s {
		if unicode.IsDigit(r) {
			return true
		}
	}
	return false
}

func (d *docService) SanitizeDocument(doc models.Document) models.Document {

	doc.Type = models.Policy.Sanitize(models.RemoveScripts(doc.Type))
	doc.FNum = models.Policy.Sanitize(models.RemoveScripts(doc.FNum))
	doc.LNum = models.Policy.Sanitize(models.RemoveScripts(doc.LNum))
	doc.Name = models.Policy.Sanitize(models.RemoveScripts(doc.Name))
	doc.Sender = models.Policy.Sanitize(models.RemoveScripts(doc.Sender))
	doc.Ispolnitel = models.Policy.Sanitize(models.RemoveScripts(doc.Ispolnitel))
	doc.Result = models.Policy.Sanitize(models.RemoveScripts(doc.Result))
	doc.Familiar = models.Policy.Sanitize(models.RemoveScripts(doc.Familiar))
	doc.Copy = models.Policy.Sanitize(models.RemoveScripts(doc.Copy))
	doc.Location = models.Policy.Sanitize(models.RemoveScripts(doc.Location))
	doc.Creator = models.Policy.Sanitize(models.RemoveScripts(doc.Creator))

	return doc
}
