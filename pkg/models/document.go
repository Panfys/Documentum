package models

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"html/template"
	"mime/multipart"
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