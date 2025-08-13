package models

type User struct {
	UserID        string
	Email         string
	PasswordHash  []byte
	EmailVerified bool
	GoogleID      string
}
