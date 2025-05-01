package models

import (
	"database/sql"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"html/template"
	"mime/multipart"
	"time"
)

// Объект для работы с документами
type Document struct {
	ID          int    `json:"id"`
	Type        string `json:"type"`
	FNum        string `json:"fnum"`
	FDate       string `json:"fdate"`
	LNum        string `json:"lnum"`
	LDateStr    string `json:"ldate"`
	LDate       sql.NullString
	Name        string `json:"name"`
	Sender      string `json:"sender"`
	Ispolnitel  string `json:"ispolnitel"`
	Result      string `json:"result"`
	Familiar    string `json:"familiar"`
	Count       int    `json:"count"`
	Copy        string `json:"copy"`
	Width       string `json:"width"`
	Location    string `json:"location"`
	FileURL     string
	File        multipart.File
	FileHeader  *multipart.FileHeader
	Creator     string `json:"creator"`
	Resolutions []*Resolution
}

type Resolution struct {
	ID         string `json:"id"`
	DocID      int    `json:"doc_id"`
	Ispolnitel string `json:"ispolnitel"`
	Text       string `json:"text"`
	Time       sql.NullString
	TimeStr    string `json:"time"`
	Date       string `json:"date"`
	User       string `json:"user"`
	Creator    string `json:"creator"`
	Result     string `json:"result"`
}

type DocSettings struct {
	DocType   string
	DocCol    string
	DocSet    string
	DocDatain string
	DocDatato string
}

type PageData struct {
	SRC template.HTML
}

func ParseDate(date string) (string, error) {
	newdate, err := time.Parse(time.RFC3339, date)
	if err != nil {
		newdate, err = time.Parse("2006-01-02", date)
		if err != nil {
			return "", err
		}
	}

	formattedDate := "<br>от " + newdate.Format("02.01.2006") + " г."
	return formattedDate, nil
}

func ParseResolutionDate(date string) (string, error) {

	newdate, err := time.Parse(time.RFC3339, date)
	if err != nil {
		newdate, err = time.Parse("2006-01-02", date)
		if err != nil {
			return "", err
		}
	}

	formateDate := newdate.Format("02.01.2006") + " г."
	return formateDate, nil
}

func ParseTime(restime string) (string, error) {

	newtime, err := time.Parse(time.RFC3339, restime)
	if err != nil {
		newtime, err = time.Parse("2006-01-02", restime)
		if err != nil {
			return "", err
		}
	}

	// Форматируем дату в нужный формат
	formateTime := "Исполнить до " + newtime.Format("02.01.2006") + " г."
	return formateTime, nil
}

func StringToDateNullString(dateStr string) (sql.NullString, error) {
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
			// Если дата распарсилась успешно - возвращаем в стандартном формате
			return sql.NullString{
				String: parsedTime.Format("2006-01-02"),
				Valid:  true,
			}, nil
		}
	}

	return sql.NullString{}, errors.New("неверный формат даты")
}
