package usecases

import (
	"clean-architecture/internal/domain/entities"
	"clean-architecture/internal/domain/repositories"
)

type SubjectUsecase interface {
	CreateSubject(subject *entities.Subject) error
	GetSubjectByID(id int) (*entities.Subject, error)
	UpdateSubject(subject *entities.Subject) error
	DeleteSubject(id int) error
	GetAllSubjects() ([]*entities.Subject, error)
}

type subjectUsecase struct {
	subjectRepo repositories.SubjectRepository
}

func NewSubjectUsecase(sr repositories.SubjectRepository) SubjectUsecase {
	return &subjectUsecase{subjectRepo: sr}
}

func (u *subjectUsecase) CreateSubject(subject *entities.Subject) error {
	return u.subjectRepo.CreateSubject(subject)
}

func (u *subjectUsecase) GetSubjectByID(id int) (*entities.Subject, error) {
	return u.subjectRepo.GetSubjectByID(id)
}

func (u *subjectUsecase) UpdateSubject(subject *entities.Subject) error {
	return u.subjectRepo.UpdateSubject(subject)
}

func (u *subjectUsecase) DeleteSubject(id int) error {
	return u.subjectRepo.DeleteSubject(id)
}

func (u *subjectUsecase) GetAllSubjects() ([]*entities.Subject, error) {
	return u.subjectRepo.GetAllSubjects()
}
