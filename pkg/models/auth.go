package models

type contextKey string

var LoginKey contextKey = "login"

type AuthData struct {
	Login    string `json:"login"`
	Pass     string `json:"pass"`
	Remember bool   `json:"remember"`
	Agent string
	IP string
}
