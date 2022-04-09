package models

import "net/http"

type Session struct {
	Id             string `json:"id"`
	PasswordHashed string `json:"-"`
}

func (s Session) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}