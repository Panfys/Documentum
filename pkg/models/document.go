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
	IDStr       string 
	Type        string `json:"type"`
	FNum        string `json:"fnum"`
	FDate       string `json:"fdate"`
	LNum        string `json:"lnum"`
	LDateStr    string `json:"ldate"`
	LDate       sql.NullString
	Name        string `json:"name"`
	Sender      string `json:"sender"`
	Sender1     string `json:"sender1"`
	Sender2     string `json:"sender2"`
	Sender3     string `json:"sender3"`
	Sender4     string `json:"sender4"`
	Ispolnitel  string `json:"ispolnitel"`
	Result      string `json:"result"`
	Familiar    string `json:"familiar"`
	Count       string `json:"count"`
	Copy        string `json:"copy"`
	Copy1       string `json:"copy1"`
	Copy2       string `json:"copy2"`
	Copy3       string `json:"copy3"`
	Copy4       string `json:"copy4"`
	Width       string `json:"width"`
	Location    string `json:"location"`
	FileURL     string `json:"file"`
	File        multipart.File
	FileHeader  *multipart.FileHeader
	Creator     string `json:"creator"`
	CreatedAt   time.Time
	Resolutions []Resolution `json:"resolutions"`
}

type Directive struct {
	ID                 int    `json:"id"`
	Number             string `json:"number"`
	Date               string `json:"date"`
	Name               string `json:"name"`
	Autor              string `json:"autor"`
	NumCoverLetter     string `json:"numCoverLetter"`
	DateCoverLetterStr string `json:"dateCoverLetter"`
	DateCoverLetter    sql.NullString
	CountCopy          string `json:"countCopy"`
	Sender             string `json:"sender"`
	NumSendLetter      string `json:"numSendLetter"`
	DateSendLetterStr  string `json:"dateSendLetter"`
	DateSendLetter     sql.NullString
	CountSendCopy      string `json:"countSendCopy"`
	Familiar           string `json:"familiar"`
	Location           string `json:"location"`
	FileURL            string `json:"fileURL"`
	File               multipart.File
	FileHeader         *multipart.FileHeader
	Creator            string `json:"creator"`
	CreatedAt          time.Time
}

type Inventory struct {
	ID                 int    `json:"id"`
	Number             string `json:"number"`
	NumCoverLetter     string `json:"numCoverLetter"`
	DateCoverLetterStr string `json:"dateCoverLetter"`
	DateCoverLetter    sql.NullString
	Name               string `json:"name"`
	Sender             string `json:"sender"`
	CountCopy          string `json:"countCopy"`
	Copy               string `json:"copy"`
	Addressee          string `json:"addressee"`
	NumSendLetter      string `json:"numSendLetter"`
	DateSendLetterStr  string `json:"dateSendLetter"`
	DateSendLetter     sql.NullString
	SendCopy           string `json:"sendCopy"`
	Familiar           string `json:"familiar"`
	Location           string `json:"location"`
	FileURL            string `json:"fileURL"`
	File               multipart.File
	FileHeader         *multipart.FileHeader
	Creator            string `json:"creator"`
	CreatedAt          time.Time
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
	CreatedAt   time.Time
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
