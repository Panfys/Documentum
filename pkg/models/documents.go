package models

import (
	"database/sql"
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
	FDate       sql.NullString
	LNum        string `json:"lnum"`
	LDate       sql.NullString
	Name        string `json:"name"`
	Sender      string `json:"sender"`
	Ispolnitel  string `json:"ispolnitel"`
	Result      string `json:"result"`
	Familiar    string `json:"familiar"`
	Count       int
	Copy        string `json:"copy"`
	Width       int    `json:"width"`
	Location    string `json:"location"`
	FileURL     string
	File        multipart.File
	FileHeader  *multipart.FileHeader
	Creator     string `json:"creator"`
	Resolutions []*Resolution
}

type ResolutionDB struct {
	ID         int
	DocID      int            `json:"doc_id"`
	Ispolnitel string         `json:"ispolnitel"`
	Text       string         `json:"text"`
	Time       sql.NullString `json:"time"`
	Date       string         `json:"date"`
	User       string         `json:"user"`
	Creator    string         `json:"creator"`
	Result     string         `json:"result"`
}

type Resolution struct {
	DocID      int    `json:"doc_id"`
	Ispolnitel string `json:"ispolnitel"`
	Text       string `json:"text"`
	Time       string `json:"time"`
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
