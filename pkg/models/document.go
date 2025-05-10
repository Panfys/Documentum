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
	Count       int    `json:"count,string"`
	Copy        string `json:"copy"`
	Width       string `json:"width"`
	Location    string `json:"location"`
	FileURL     string `json:"file"`
	File        multipart.File
	FileHeader  *multipart.FileHeader
	Creator     string       `json:"creator"`
	CreatedAt   string       `json:"createdAt"`
	Resolutions []Resolution `json:"resolutions"`
}

type Resolution struct {
	ID          string `json:"id"`
	DocID       int    `json:"doc_id"`
	Type        string `json:"type"`
	Ispolnitel  string `json:"ispolnitel"`
	Text        string `json:"text"`
	Deadline    sql.NullString
	DeadlineStr string `json:"deadline"`
	Date        string `json:"date"`
	User        string `json:"user"`
	Creator     string `json:"creator"`
	Result      string `json:"result"`
}

type DocSettings struct {
	DocType   string `json:"type"`
	DocCol    string `json:"col"`
	DocSet    string `json:"set"`
	DocDatain string `json:"datain"`
	DocDatato string `json:"datato"`
}

type PageData struct {
	SRC template.HTML
}
