package valid

import (
	"database/sql"
	"documentum/pkg/models"
	"errors"
	"regexp"
	"strings"
	"unicode"
)

func (v *validatService) ValidResolution(reqRes *models.Resolution) (models.Resolution, error) {

	res := v.sanitizeResolution(reqRes)

	err := v.validResText(res.Text)
	if err != nil {
		return models.Resolution{}, err
	}

	resDate, err := v.stringToDateNullString(res.Date)
	if err != nil {
		return models.Resolution{}, errors.New("дата резолюции указана неверно")
	}

	err = v.validResDate(resDate)
	if err != nil {
		return models.Resolution{}, err
	}

	if res.Type == "task" {
		if !v.validResIspolnitel(res.Ispolnitel) {
			return models.Resolution{}, errors.New(`исполнитель документа в резолюции указан неверно, пример: "Панфилов А.П." или "Панфилов А.П., Якель Е.В."`)
		}
		res.Deadline, err = v.stringToDateNullString(res.DeadlineStr)
		if err != nil {
			return models.Resolution{}, errors.New("срок исполнения документа указан неверно")
		}
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
		return errors.New("текст резолюции не указан")
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

func (v *validatService) validResDate(date sql.NullString) error {
	if !date.Valid {
		return errors.New(`дата резолюции не указана`)
	}

	return nil
}

func (v *validatService) sanitizeResolution(res *models.Resolution) models.Resolution {

	res.Ispolnitel = v.policy.Sanitize(res.Ispolnitel)
	res.Result = v.policy.Sanitize(res.Result)
	res.Text = v.policy.Sanitize(res.Text)
	res.User = v.policy.Sanitize(res.User)
	res.Creator = v.policy.Sanitize(res.Creator)

	return *res
}
