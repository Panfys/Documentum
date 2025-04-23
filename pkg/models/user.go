package models

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"regexp"
)

// Объект для работы с подразделениями
type Unit struct {
	ID   int
	Name string
}

// Объект для работы с пользователями
type User struct {
	ID     int    `json:"id"`
	Login  string `json:"login"`
	Name   string `json:"name"`
	Func   string `json:"func"`
	Unit   string `json:"unit"`
	Group  string `json:"group"`
	Pass   string `json:"pass"`
	Status string `json:"status"`
	Icon   string `json:"icon"`
}

type AccountData struct {
	Login  string
	Name   string 
	Func   string 
	Unit   string 
	Group  string 
	Status string
	Icon   string
	ToDay  string
}

// Объект для работы с документами
type Document struct {
	ID          int    `json:"id"`
	Type        string `json:"type"`
	FNum        string `json:"fnum"`
	FDate       string `json:"fdate"`
	LNum        string `json:"lnum"`
	LDate       sql.NullString
	Name        string `json:"name"`
	Sender      string `json:"sender"`
	Ispolnitel  string `json:"ispolnitel"`
	Result      string `json:"result"`
	Familiar    string `json:"familiar"`
	Count       int    `json:"count"`
	Copy        string `json:"copy"`
	Width       int    `json:"width"`
	Location    string `json:"location"`
	File        string
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

// Метод для валидации ФИО пользователя
func (u *User) ValidName(name string) bool {

	pattern := `^[А-ЯЁ][а-яё]+ [А-ЯЁ]\.[А-ЯЁ]\.$`
	re := regexp.MustCompile(pattern)

	return re.MatchString(name)

}

// Метод для валидации логина пользователя
func (u *User) ValidLogin(login string) bool {

	pattern := `^[a-zA-Z0-9](?:[a-zA-Z0-9._-]{1,10}[a-zA-Z0-9])?$`
    re := regexp.MustCompile(pattern)
    return re.MatchString(login) && len(login) >= 3 && len(login) <= 12
}

// Метод для валидации пароля пользователя
func (u *User) ValidPass(pass string) bool {

	pattern := `^[a-zA-Z-ЯЁа-яё0-9.]{6,30}$`
	re := regexp.MustCompile(pattern)

	return re.MatchString(pass)
}
