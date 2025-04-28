package service

import (
	"documentum/pkg/models"
	"regexp"
	"errors"
	"strings"
	"unicode"
)

func (d *docService) validResolution(res *models.Resolution) error {

	if res.Ispolnitel != "NULL" && !d.validResIspolnitel(res.Ispolnitel) {
		return errors.New(`исполнитель документа в резолюции указан неверно, пример: "Панфилов А.П." или "Панфилов А.П., Якель Е.В."`)
	}

	err := d.validResText(res.Text) 
	if err != nil {
		return err
	}

	resDate, err := models.StringToDateNullString(res.Date)
	if err != nil {
		return errors.New("дата документа указана неверно")
	}
	if !d.validDocDate(resDate) {
		return errors.New("дата резолюции не указана")
	}

	if !d.validResUser(res.User) {
		return errors.New(`автор резолюции указан неверно, пример: "Е.Лыков"`)
	}
	
	return nil
}

// Ispolnitel
func (d *docService) validResIspolnitel(ispolnitel string) bool {

	if ispolnitel == "" {
		return false
	}

	// Основной паттерн для одного исполнителя: Фамилия И.О.
	singlePattern := `[А-ЯЁ][а-яё]+ [А-ЯЁ]\.[А-ЯЁ]\.`
	// Общий паттерн: один или несколько исполнителей через запятую с пробелом
	fullPattern := `^(` + singlePattern + `)(, ` + singlePattern + `)*$`

	re := regexp.MustCompile(fullPattern)
	return re.MatchString(ispolnitel)
}

// Text
func (d *docService) validResText(text string) error {
	trimText := strings.TrimSpace(text)

	if trimText == "" {
		return errors.New("текст резолюции не указано")
	}

	firstChar := []rune(trimText)[0]
	if !unicode.IsUpper(firstChar) && !unicode.IsDigit(firstChar) {
		return errors.New("текст резолюции должен начинаться с заглавной буквы")
	}

	if len(trimText) > 200 {
		return errors.New(`текст резолюции не должен превышать 200 символов`)
	}

	return nil
}

// User
func (d *docService) validResUser(user string) bool {
    pattern := `^[А-ЯЁ]\.[А-ЯЁ][а-яё]+$`
    re := regexp.MustCompile(pattern)
    return re.MatchString(user)
}

func (d *docService) sanitizeResolution(res *models.Resolution) models.Resolution {

	res.Ispolnitel = models.Policy.Sanitize(models.RemoveScripts(res.Ispolnitel))
	res.Result = models.Policy.Sanitize(models.RemoveScripts(res.Result))
	res.Text = models.Policy.Sanitize(models.RemoveScripts(res.Text))
	res.User = models.Policy.Sanitize(models.RemoveScripts(res.User))
	res.Creator = models.Policy.Sanitize(models.RemoveScripts(res.Creator))

	return *res
}
