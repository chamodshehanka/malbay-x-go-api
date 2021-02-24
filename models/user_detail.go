package models

type UserDetail struct {
	Password string `json:"password"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Address  string `json:"address"`
}
