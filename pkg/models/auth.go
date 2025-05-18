package models

type contextKey string

const (
    LoginKey contextKey = "login"
    UserAgentKey contextKey = "userAgent"
    IPKey contextKey = "ip"
)

type AuthData struct {
	Login    string `json:"login"`
	Pass     string `json:"pass"`
	Remember bool   `json:"remember"`
	Agent string
	IP string
}
