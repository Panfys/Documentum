package valid

import (
	"database/sql"
	"documentum/pkg/models"
	"errors"
	"path/filepath"
	_ "github.com/go-sql-driver/mysql"
	"mime"
	"mime/multipart"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

type DocValidatService interface {
	ValidIngoingDoc(reqDoc models.Document) (models.Document, error)
	ValidResolution(res *models.Resolution) (models.Resolution, error)
}

func (v *validatService) ValidIngoingDoc(reqDoc models.Document) (models.Document, error) {

	doc := v.sanitizeDocument(&reqDoc)

	if !v.validDocType(doc.Type) {
		return models.Document{}, errors.New("тип документа указан некорректно")
	}

	err := v.validDocFNum(doc.FNum)
	if err != nil {
		return models.Document{}, err
	}

	docFDate, err := models.StringToDateNullString(doc.FDate)
	if err != nil {
		return models.Document{}, errors.New("дата документа указана неверно")
	}

	if !v.validDocDate(docFDate) {
		return models.Document{}, errors.New("дата документа не указана")
	}

	doc.LDate, err = models.StringToDateNullString(doc.LDateStr)
	if err != nil {
		return models.Document{}, errors.New("дата поступившего документа указана неверно")
	}

	err = v.validDocLdate(doc.LDate, doc.LNum)
	if err != nil {
		return models.Document{}, err
	}

	err = v.validDocName(doc.Name)
	if err != nil {
		return models.Document{}, err
	}

	err = v.validDocSender(doc.Sender)
	if err != nil {
		return models.Document{}, err
	}

	err = v.validDocIspolnitel(doc.Ispolnitel, doc.Resolutions)
	if err != nil {
		return models.Document{}, err
	}

	err = v.validDocCount(doc.Count)
	if err != nil {
		return models.Document{}, err
	}

	err = v.validDocWidth(doc.Width)
	if err != nil {
		return models.Document{}, err
	}

	if doc.Location != "" && !v.validDocLocation(doc.Location) {
		return models.Document{}, errors.New(`отметка о подшивке указана неверно, примеры: "Дело 12, стр. 12", "Дело 1, стр. 12-19"`)
	}

	err = v.validDocFile(doc.FileHeader)
	if err != nil {
		return models.Document{}, err
	}

	return doc, nil
}

func (v *validatService) validDocType(docType string) bool {

	switch docType {
	case "Входящий", "Исходящий", "Директива", "Инвентарный":
		return true
	default:
		return false
	}
}

func (v *validatService) validDocFNum(fnum string) error {
	trimFnum := strings.TrimSpace(fnum)

	if trimFnum == "" {
		return errors.New(`номер документа не указан`)
	}

	if trimFnum == "№" || trimFnum == "Вх. №" {
		return errors.New(`номер документа не должен содержать только "№" или "Вх. №"`)
	}

	if strings.HasPrefix(trimFnum, "Вх. № ") {
		numPart := strings.TrimPrefix(trimFnum, "Вх. № ")
		if !v.containsDigit(numPart) {
			return errors.New(`номер документа должен начинаться либо с "Вх. № ", либо с цифры`)
		}
	} else {
		if len(trimFnum) == 0 || !unicode.IsDigit(rune(trimFnum[0])) {
			return errors.New(` если номер документа не начинается с "Вх. № ", то должен начинаться с цифры`)
		}
	}

	if !v.containsDigit(trimFnum) {
		return errors.New(`номер документа указан некорректно, примеры правильного номера: "Вх. № 1", "Вх. № 317/124дсп", "312/8321", "215/1111дсп"`)
	}

	return nil
}

func (v *validatService) validDocLNum(lnum string) error {
	trimFnum := strings.TrimSpace(lnum)

	if trimFnum == "" {
		return errors.New(`номер документа не указан`)
	}

	if trimFnum == "№" || trimFnum == "Исх. №" {
		return errors.New(`номер документа не должен содержать только "№" или "Исх. №"`)
	}

	if strings.HasPrefix(trimFnum, "Исх. № ") {
		numPart := strings.TrimPrefix(trimFnum, "Исх. № ")
		if !v.containsDigit(numPart) {
			return errors.New(`номер документа должен начинаться либо с "Исх. № ", либо с цифры`)
		}
	} else {
		if len(trimFnum) == 0 || !unicode.IsDigit(rune(trimFnum[0])) {
			return errors.New(` если номер документа не начинается с "Исх. № ", то должен начинаться с цифры`)
		}
	}

	if !v.containsDigit(trimFnum) {
		return errors.New(`номер документа указан некорректно, примеры правильного номера: "Исх. № 1", "Исх. № 317/124дсп", "330/8321", "215/1111дсп"`)
	}

	return nil
}

func (v *validatService) validDocDate(date sql.NullString) bool {
	return date.Valid
}

func (v *validatService) validDocLdate(date sql.NullString, num string) error {
	if num != "" {
		err := v.validDocLNum(num)
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

func (v *validatService) validDocName(name string) error {
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

func (v *validatService) validDocSender(sender string) error {
	trimSender := strings.TrimSpace(sender)

	if trimSender == "" {
		return errors.New("отправитель документа не указано")
	}

	if len(trimSender) > 100 {
		return errors.New(`отправитель документа не должно превышать 100 символов`)
	}

	return nil
}

func (v *validatService) validDocIspolnitel(ispolnitel string, resolutions []*models.Resolution) error {
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

func (v *validatService) validDocCount(count int) error {

	if count < 1 {
		return errors.New("количестов экземпляров не может быть меньше единицы")
	}

	return nil
}

func (v *validatService) validDocWidth(width string) error {
	if width == "" {
		return errors.New("количество листов не указано")
	}

	parts := strings.Split(width, "/")

	// Проверяем что частей не больше 2
	if len(parts) > 2 {
		return errors.New(`количество листов указано неверно, пример: "1", "1/25"`)
	}

	// Проверяем каждую часть
	for _, part := range parts {
		num, err := strconv.Atoi(part)
		if err != nil {
			return errors.New(`количество листов указано неверно, пример: "1", "1/25"`)
		}
		if num < 1 {
			return errors.New("количество листов не может быть меньше единицы")
		}
	}

	return nil
}

func (v *validatService) validDocLocation(location string) bool {

	// Проверяем общий формат: "Дело <число>, стр. <число>[-<число>]"
	pattern := `^Дело (\d+), стр\. (\d+)(?:-(\d+))?$`
	re := regexp.MustCompile(pattern)
	matches := re.FindStringSubmatch(location)

	if matches == nil {
		return false
	}

	// Проверяем, что все числа валидны (положительные)
	_, err := strconv.Atoi(matches[1]) // Номер дела
	if err != nil || matches[1] == "0" {
		return false
	}

	_, err = strconv.Atoi(matches[2]) // Первая страница
	if err != nil || matches[2] == "0" {
		return false
	}

	// Если есть диапазон (например, "12-19"), проверяем второе число
	if matches[3] != "" {
		pageEnd, err := strconv.Atoi(matches[3])
		if err != nil || matches[3] == "0" {
			return false
		}

		pageStart, _ := strconv.Atoi(matches[2])
		if pageEnd <= pageStart {
			return false
		}
	}

	return true
}

func (v *validatService) validDocFile(file *multipart.FileHeader) error {
	if file == nil {
		return errors.New("файл не указан")
	}

	if file.Size == 0 {
		return errors.New("файл не должен быть пустым")
	}

	// Получаем MIME-тип файла
	ext := strings.ToLower(filepath.Ext(file.Filename))
	mimeType := mime.TypeByExtension(ext)

	// Проверяем, что файл PDF или изображение
	isPDF := ext == ".pdf"
	isImage := strings.HasPrefix(mimeType, "image/") &&
		(ext == ".jpg" || ext == ".jpeg" || ext == ".png" || ext == ".gif")

	if !isPDF && !isImage {
		return errors.New("поддерживаются только файлы PDF и изображения")
	}

	return nil
}

func (v *validatService) containsDigit(s string) bool {
	for _, r := range s {
		if unicode.IsDigit(r) {
			return true
		}
	}
	return false
}

func (v *validatService) sanitizeDocument(doc *models.Document) models.Document {

	doc.Type = v.policy.Sanitize(doc.Type)
	doc.FNum = v.policy.Sanitize(doc.FNum)
	doc.LNum = v.policy.Sanitize(doc.LNum)
	doc.Name = v.policy.Sanitize(doc.Name)
	doc.Sender = v.policy.Sanitize(doc.Sender)
	doc.Ispolnitel = v.policy.Sanitize(doc.Ispolnitel)
	doc.Result = v.policy.Sanitize(doc.Result)
	doc.Familiar = v.policy.Sanitize(doc.Familiar)
	doc.Copy = v.policy.Sanitize(doc.Copy)
	doc.Width = v.policy.Sanitize(doc.Width)
	doc.Location = v.policy.Sanitize(doc.Location)
	doc.FileHeader.Filename = v.policy.Sanitize(doc.FileHeader.Filename)
	doc.Creator = v.policy.Sanitize(doc.Creator)

	return *doc
}
