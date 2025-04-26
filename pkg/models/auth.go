package models


type contextKey string

var LoginKey contextKey = "login"

// Объект для работы с подразделениями
type Unit struct {
	ID   int
	Name string
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
