package models

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

