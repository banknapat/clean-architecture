package entities

import "time"

type Student struct {
	StudentID   int
	Gender      string
	FirstName   string
	LastName    string
	DateOfBirth time.Time
	Nationality string
	Ethnicity   string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
