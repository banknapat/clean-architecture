package entities

import "time"

type Subject struct {
	SubjectID   int
	SubjectName string
	CreditHours int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
