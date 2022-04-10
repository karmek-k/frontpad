package models

type Session struct {
	Id             string `json:"id"`
	PasswordHashed string `json:"-"`
}
