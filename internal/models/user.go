package models

type User struct {
	ID           int    `json:"id"`
	Email        string `json:"email"`
	PasswordHash string `json:"-"` // Hidden in JSON
	CreatedAt    string `json:"created_at"`
}