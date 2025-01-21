package repositories

import (
	"gorm.io/gorm"

	"clean-architecture/internal/domain/entities"
	"clean-architecture/internal/domain/repositories"
)

type subjectRepositoryImpl struct {
	db *gorm.DB
}

func NewSubjectRepository(db *gorm.DB) repositories.SubjectRepository {
	return &subjectRepositoryImpl{db}
}

func (r *subjectRepositoryImpl) CreateSubject(subject *entities.Subject) error {
	return r.db.Create(subject).Error
}

func (r *subjectRepositoryImpl) GetSubjectByID(id int) (*entities.Subject, error) {
	var subject entities.Subject
	err := r.db.First(&subject, id).Error
	if err != nil {
		return nil, err
	}
	return &subject, nil
}

func (r *subjectRepositoryImpl) UpdateSubject(subject *entities.Subject) error {
	return r.db.Save(subject).Error
}

func (r *subjectRepositoryImpl) DeleteSubject(id int) error {
	return r.db.Delete(&entities.Subject{}, id).Error
}

func (r *subjectRepositoryImpl) GetAllSubjects() ([]*entities.Subject, error) {
	var subjects []*entities.Subject
	err := r.db.Find(&subjects).Error
	return subjects, err
}
