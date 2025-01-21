package repositories

import "clean-architecture/internal/domain/entities"

type SubjectRepository interface {
	CreateSubject(subject *entities.Subject) error
	GetSubjectByID(id int) (*entities.Subject, error)
	UpdateSubject(subject *entities.Subject) error
	DeleteSubject(id int) error
	GetAllSubjects() ([]*entities.Subject, error)
}
