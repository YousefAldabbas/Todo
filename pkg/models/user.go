package models

type User struct {
	ID       int    `json:"id"`
	Sub      string `json:"sub"`
	Username string `json:"username"`
	Password string `json:"-"`
}
