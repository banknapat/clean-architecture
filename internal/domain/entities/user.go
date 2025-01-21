package entities

import "time"

type User struct {
	UserID       int
	Username     string
	Email        string
	PasswordHash string
	RefreshToken string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
