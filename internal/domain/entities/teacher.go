package entities

import "time"

type Teacher struct {
	TeacherID                int
	Gender                   string
	FirstName                string
	LastName                 string
	DateOfBirth              time.Time
	EducationalQualification string
	Nationality              string
	Ethnicity                string
	CreatedAt                time.Time
	UpdatedAt                time.Time
}
