package valid

import (
	"database/sql"
	"documentum/pkg/models"
	"errors"
	"mime"
	"mime/multipart"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode"

	_ "github.com/go-sql-driver/mysql"
)

type DocValidatService interface {
	ValidDocument(reqDoc models.Document) (models.Document, error)
	ValidResolution(res *models.Resolution) (models.Resolution, error)
	ValidDirective(reqDir models.Directive) (models.Directive, error)
	ValidInventory(reqInv models.Inventory) (models.Inventory, error)
}

func (v *validatService) ValidDocument(reqDoc models.Document) (models.Document, error) {

	doc := v.sanitizeDocument(&reqDoc)

	if !v.validDocType(doc.Type) {
		return models.Document{}, errors.New("тип документа указан некорректно")
	}

	docCount, err := strconv.Atoi(doc.Count)

	if err != nil {
		return models.Document{}, errors.New("количество экземпляров указано некорректно")
	}	

	doc.LDate, err = v.stringToDateNullString(doc.LDateStr)
	if err != nil {
		return models.Document{}, errors.New("дата поступившего документа указана неверно")
	}

	if doc.Type == "Входящий" {
		err := v.validDocFNum(doc.FNum, "Вх. № ")
		if err != nil {
			return models.Document{}, err
		}

		err = v.validDocLdate(doc.LDate, doc.LNum, "Исх. №")
		if err != nil {
			return models.Document{}, err
		}

		err = v.validDocSender(doc.Sender)
		if err != nil {
			return models.Document{}, err
		}

		err = v.validDocCopy(doc.Copy)
		if err != nil {
			return models.Document{}, err
		}

	} else {
		err := v.validDocFNum(doc.FNum, "Исх. № 330/")
		if err != nil {
			return models.Document{}, err
		}

		err = v.validDocLdate(doc.LDate, doc.LNum, "Вх. № ")
		if err != nil {
			return models.Document{}, err
		}

		err = v.validDocSender(doc.Sender)
		if err != nil {
			return models.Document{}, err
		}

		err = v.validDocCopy(doc.Copy)
		if err != nil {
			return models.Document{}, err
		}

		if docCount > 1 && docCount < 6 {
			additionalSenders := []string{doc.Sender1, doc.Sender2, doc.Sender3, doc.Sender4}
			additionalCopyes := []string{doc.Copy1, doc.Copy2, doc.Copy3, doc.Copy4}
			for i := 0; i < docCount-1; i++ {
				sender := additionalSenders[i]
				copy := additionalCopyes[i]

				if err := v.validDocSender(sender); err != nil {
					return models.Document{}, err
				}

				err = v.validDocCopy(copy)
				if err != nil {
					return models.Document{}, err
				}

				doc.Sender += " <br> " + sender
				doc.Copy += " <br> " + copy
			}
		}
	}

	docFDate, err := v.stringToDateNullString(doc.FDate)
	if err != nil {
		return models.Document{}, errors.New("дата документа указана неверно")
	}

	err = v.validDocDate(docFDate)
	if err != nil {
		return models.Document{}, err
	}

	err = v.validDocName(doc.Name)
	if err != nil {
		return models.Document{}, err
	}

	err = v.validDocIspolnitel(doc.Ispolnitel, doc.Resolutions)
	if err != nil {
		return models.Document{}, err
	}

	err = v.validDocCount(docCount)
	if err != nil {
		return models.Document{}, err
	}

	err = v.validDocWidth(doc.Width)
	if err != nil {
		return models.Document{}, err
	}

	if doc.Location != "" && !v.validDocLocation(doc.Location) {
		return models.Document{}, errors.New(`отметка о подшивке указана неверно, примеры: "Дело 12, стр. 12", "Дело 1, стр. 12-19", "Реестр № 1", "Акт № 15дсп от 14.05.2025 г."`)
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

func (v *validatService) validDocFNum(fnum, Type string) error {
	trimFnum := strings.TrimSpace(fnum)

	if trimFnum == "" || trimFnum == Type || trimFnum == "№" {
		return errors.New(`порядковый номер документа не указан`)
	}

	if strings.HasPrefix(trimFnum, Type) {
		numPart := strings.TrimPrefix(trimFnum, Type)
		if !v.containsDigit(numPart) {
			return errors.New(`порядковый номер документа указан неверно, примеры верного номера: "` + Type + `123", "` + Type + `123дсп", "` + Type + `123/124", "` + Type + `123/654дсп"`)
		}
	} else {
		return errors.New(`порядковый номер документа указан неверно, примеры верного номера: "` + Type + `123", "` + Type + `123дсп", "` + Type + `123/124", "` + Type + `123/654дсп"`)
	}

	return nil
}

func (v *validatService) validDocLNum(lnum string, Type string) error {
	trimFnum := strings.TrimSpace(lnum)

	if trimFnum == Type {
		return errors.New(`номер документа не указан`)
	}

	if strings.HasPrefix(trimFnum, Type) {
		numPart := strings.TrimPrefix(trimFnum, Type)
		if !v.containsDigit(numPart) {
			return errors.New(`номер документа должен начинаться либо с ` + Type + `, либо с цифры`)
		}
	} else {
		if len(trimFnum) == 0 || !unicode.IsDigit(rune(trimFnum[0])) {
			return errors.New(`если номер документа не начинается с ` + Type + `, то должен начинаться с цифры`)
		}
	}

	if !v.containsDigit(trimFnum) {
		return errors.New(`номер документа указан некорректно, примеры правильного номера: "Исх. № 1", "Исх. № 317/124дсп", "330/8321", "215/1111дсп"`)
	}

	return nil
}

func (v *validatService) validDocDate(date sql.NullString) error {
	if !date.Valid {
		return errors.New(`дата учета документа не указана`)
	}

	return nil
}

func (v *validatService) validDocLdate(date sql.NullString, num, Type string) error {
	if num != "" {
		err := v.validDocLNum(num, Type)
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
		return errors.New("отправитель документа не указан")
	}

	if len(trimSender) > 100 {
		return errors.New(`отправитель документа не должно превышать 100 символов`)
	}

	return nil
}

func (v *validatService) validDocIspolnitel(ispolnitel string, resolutions []models.Resolution) error {
	trimIspolnitel := strings.TrimSpace(ispolnitel)

	if len(resolutions) == 0 {
		if trimIspolnitel == "" {
			return errors.New("исполнитель документа не указан")
		}

		pattern := `^[А-ЯЁ][а-яё]+ [А-ЯЁ]\.[А-ЯЁ]\.*$`
		re := regexp.MustCompile(pattern)

		if !re.MatchString(trimIspolnitel) {
			return errors.New(`исполнитель указан неверно, пример: "Панфилов А.П."`)
		}

		if len(trimIspolnitel) > 100 {
			return errors.New(`исполнитель документа не должно превышать 100 символов`)
		}
	}
	return nil
}

func (v *validatService) validDocCount(count int) error {

	if count < 1 {
		return errors.New(`количестов экземпляров должно быть больше нуля`)
	}

	return nil
}

func (v *validatService) validDocCopy(copy string) error {
	if copy == "" {
		return errors.New("номер экземпляра не указан")
	}

	// Проверяем что первый символ - цифра
	if len(copy) == 0 || copy[0] < '0' || copy[0] > '9' {
		return errors.New("номер экземпляра должен начинаться с цифры")
	}

	// Проверяем что первая цифра > 0
	if copy[0] == '0' {
		return errors.New("номер экземпляра должна быть больше нуля")
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
    casePattern := `^Дело (\d+), стр\. (\d+)(?:-(\d+))?$`
    // Проверяем форматы: "Реестр № <число>", "Акт № <число>[дсп][от <дата>]"
    registryPattern := `^(Реестр|Акт) № (\d+)(дсп)?(?:\sот\s\d{2}\.\d{2}\.\d{4}\sг\.)?$`
    
    reCase := regexp.MustCompile(casePattern)
    reRegistry := regexp.MustCompile(registryPattern)
    
    // Проверяем оба возможных формата
    if matches := reCase.FindStringSubmatch(location); matches != nil {
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
    } else if matches := reRegistry.FindStringSubmatch(location); matches != nil {
        // Проверяем номер реестра/акта
        _, err := strconv.Atoi(matches[2])
        if err != nil || matches[2] == "0" {
            return false
        }
        return true
    }
    
    return false
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
	doc.Sender1 = v.policy.Sanitize(doc.Sender1)
	doc.Sender2 = v.policy.Sanitize(doc.Sender2)
	doc.Sender3 = v.policy.Sanitize(doc.Sender3)
	doc.Sender4 = v.policy.Sanitize(doc.Sender4)
	doc.Ispolnitel = v.policy.Sanitize(doc.Ispolnitel)
	doc.Result = v.policy.Sanitize(doc.Result)
	doc.Familiar = v.policy.Sanitize(doc.Familiar)
	doc.Copy = v.policy.Sanitize(doc.Copy)
	doc.Copy1 = v.policy.Sanitize(doc.Copy1)
	doc.Copy2 = v.policy.Sanitize(doc.Copy2)
	doc.Copy3 = v.policy.Sanitize(doc.Copy3)
	doc.Copy4 = v.policy.Sanitize(doc.Copy4)
	doc.Width = v.policy.Sanitize(doc.Width)
	doc.Location = v.policy.Sanitize(doc.Location)
	doc.Creator = v.policy.Sanitize(doc.Creator)

	// Очищаем имя файла (если FileHeader не nil)
	if doc.FileHeader != nil {
		doc.FileHeader.Filename = v.policy.Sanitize(doc.FileHeader.Filename)
	}

	return *doc
}

func (v *validatService) stringToDateNullString(dateStr string) (sql.NullString, error) {
	if dateStr == "" {
		return sql.NullString{Valid: false}, nil
	}

	// Пробуем распарсить дату в разных форматах
	formats := []string{
		"2006-01-02",
		time.RFC3339,
	}

	var parsedTime time.Time
	var err error

	for _, format := range formats {
		parsedTime, err = time.Parse(format, dateStr)
		if err == nil {
			return sql.NullString{
				String: parsedTime.Format("2006-01-02"),
				Valid:  true,
			}, nil
		}
	}

	return sql.NullString{}, errors.New("неверный формат даты")
}
