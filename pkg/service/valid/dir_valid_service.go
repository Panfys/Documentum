package valid

import (
	"documentum/pkg/models"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
	"strings"
	"unicode"
)

func (v *validatService) ValidDirective(reqDir models.Directive) (models.Directive, error) {

	dir := v.sanitizeDirective(&reqDir)

	err := v.validDirNumber(dir.Number)
	if err != nil {
		return models.Directive{}, err
	}

	dirDate, err := v.stringToDateNullString(dir.Date)
	if err != nil {
		return models.Directive{}, errors.New("дата документа указана неверно")
	}

	err = v.validDocDate(dirDate)
	if err != nil {
		return models.Directive{}, err
	}

	err = v.validDirName(dir.Name)
	if err != nil {
		return models.Directive{}, err
	}

	err = v.validDirAutor(dir.Autor)
	if err != nil {
		return models.Directive{}, err
	}

	dirCountCopy, err := strconv.Atoi(dir.CountCopy)
	if err != nil {
		return models.Directive{}, errors.New("количество экземпляров указано некорректно")
	} 

	err = v.validDocCount(dirCountCopy)
	if err != nil {
		return models.Directive{}, err
	}

	dir.DateCoverLetter, err = v.stringToDateNullString(dir.DateCoverLetterStr)
	if err != nil {
		return models.Directive{}, errors.New("дата сопроводительного письма указана неверно")
	}

	dir.DateSendLetter, err = v.stringToDateNullString(dir.DateSendLetterStr)
	if err != nil {
		return models.Directive{}, errors.New("дата сопроводительного письма указана неверно")
	}

	err = v.validDocFile(dir.FileHeader)
	if err != nil {
		return models.Directive{}, err
	}

	return dir, nil
}

func (v *validatService) validDirNumber(num string) error {
	trimNum := strings.TrimSpace(num)

	if trimNum == "" || trimNum == "Приказ №" || trimNum == "№" {
		return errors.New(`порядковый номер документа не указан`)
	}

	return nil
}

func (v *validatService) validDirName(name string) error {
	trimName := strings.TrimSpace(name)

	if trimName == "" {
		return errors.New("краткое содержание не указано")
	}

	firstChar := []rune(trimName)[0]
	if !unicode.IsUpper(firstChar) && !unicode.IsDigit(firstChar) {
		return errors.New("краткое содержание должно начинаться с заглавной буквы")
	}

	if len(trimName) > 200 {
		return errors.New(`краткое содержание не должно превышать 200 символов`)
	}

	return nil
}

func (v *validatService) validDirAutor(autor string) error {
	trimAutor := strings.TrimSpace(autor)

	if trimAutor == "" {
		return errors.New(`лицо, подписавшее документ не указано`)
	}

	return nil
}

func (v *validatService) sanitizeDirective(doc *models.Directive) models.Directive {

	doc.Number = v.policy.Sanitize(doc.Number)
	doc.Date = v.policy.Sanitize(doc.Date)
	doc.Name = v.policy.Sanitize(doc.Name)
	doc.Autor = v.policy.Sanitize(doc.Autor)
	doc.NumCoverLetter = v.policy.Sanitize(doc.NumCoverLetter)
	doc.DateCoverLetterStr = v.policy.Sanitize(doc.DateCoverLetterStr)
	doc.CountCopy = v.policy.Sanitize(doc.CountCopy)
	doc.Sender = v.policy.Sanitize(doc.Sender)
	doc.NumSendLetter = v.policy.Sanitize(doc.NumSendLetter)
	doc.DateSendLetterStr = v.policy.Sanitize(doc.DateSendLetterStr)
	doc.CountSendCopy = v.policy.Sanitize(doc.CountSendCopy)
	doc.Familiar = v.policy.Sanitize(doc.Familiar)
	doc.Location = v.policy.Sanitize(doc.Location)
	doc.Creator = v.policy.Sanitize(doc.Creator)
	doc.FileHeader.Filename = v.policy.Sanitize(doc.FileHeader.Filename)

	return *doc
}
