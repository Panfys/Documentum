package valid

import (
	"documentum/pkg/models"
	"errors"
	"regexp"
	"strings"
	"unicode"
)

func (v *validatService) ValidResolution(reqRes *models.Resolution) (models.Resolution, error) {

	res := v.sanitizeResolution(reqRes)

	if res.Ispolnitel != "NULL" && !v.validResIspolnitel(res.Ispolnitel) {
		return models.Resolution{}, errors.New(`исполнитель документа в резолюции указан неверно, пример: "Панфилов А.П." или "Панфилов А.П., Якель Е.В."`)
	} else if res.Ispolnitel == "NULL" {
		res.Ispolnitel = ""
	}

	err := v.validResText(res.Text)
	if err != nil {
		return models.Resolution{}, err
	}

	res.Time, err = v.stringToDateNullString(res.TimeStr)
	if err != nil {
		return models.Resolution{}, errors.New("срок исполнения документа указан неверно")
	}

	resDate, err := v.stringToDateNullString(res.Date)
	if err != nil {
		return models.Resolution{}, errors.New("дата документа указана неверно")
	}
	if !v.validDocDate(resDate) {
		return models.Resolution{}, errors.New("дата резолюции не указана")
	}

	if !v.validResUser(res.User) {
		return models.Resolution{}, errors.New(`автор резолюции указан неверно, пример: "Е.Лыков"`)
	}

	return res, nil
}

// Ispolnitel
func (v *validatService) validResIspolnitel(ispolnitel string) bool {

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
func (v *validatService) validResText(text string) error {
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
func (v *validatService) validResUser(user string) bool {
	pattern := `^[А-ЯЁ]\.[А-ЯЁ][а-яё]+$`
	re := regexp.MustCompile(pattern)
	return re.MatchString(user)
}

func (v *validatService) sanitizeResolution(res *models.Resolution) models.Resolution {

	res.Ispolnitel = v.policy.Sanitize(res.Ispolnitel)
	res.Result = v.policy.Sanitize(res.Result)
	res.Text = v.policy.Sanitize(res.Text)
	res.User = v.policy.Sanitize(res.User)
	res.Creator = v.policy.Sanitize(res.Creator)

	return *res
}
