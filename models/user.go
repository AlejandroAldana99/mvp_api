package models

type UserData struct {
	FirstName string   `json:"firstname"`
	LastName  string   `json:"lastname"`
	Email     string   `json:"email"`
	Password  string   `json:"password"`
	Orders    []string `json:"orders"`
}
