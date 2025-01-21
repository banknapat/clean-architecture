package repositories

import (
	"gorm.io/gorm"

	"clean-architecture/internal/domain/entities"
	"clean-architecture/internal/domain/repositories"
)

type teacherRepositoryImpl struct {
	db *gorm.DB
}

func NewTeacherRepository(db *gorm.DB) repositories.TeacherRepository {
	return &teacherRepositoryImpl{db}
}

func (r *teacherRepositoryImpl) CreateTeacher(teacher *entities.Teacher) error {
	return r.db.Create(teacher).Error
}

func (r *teacherRepositoryImpl) GetTeacherByID(id int) (*entities.Teacher, error) {
	var teacher entities.Teacher
	err := r.db.First(&teacher, id).Error
	if err != nil {
		return nil, err
	}
	return &teacher, nil
}

func (r *teacherRepositoryImpl) UpdateTeacher(teacher *entities.Teacher) error {
	return r.db.Save(teacher).Error
}

func (r *teacherRepositoryImpl) DeleteTeacher(id int) error {
	return r.db.Delete(&entities.Teacher{}, id).Error
}

func (r *teacherRepositoryImpl) GetAllTeachers() ([]*entities.Teacher, error) {
	var teachers []*entities.Teacher
	err := r.db.Find(&teachers).Error
	return teachers, err
}

func (r *teacherRepositoryImpl) AssignSubjectsToTeacher(teacherID int, subjectIDs []int) error {
	tx := r.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	for _, sid := range subjectIDs {
		err := tx.Exec(`
            INSERT INTO teacher_subjects (teacher_id, subject_id)
            VALUES (?, ?)
            ON CONFLICT DO NOTHING
        `, teacherID, sid).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	return tx.Commit().Error
}

func (r *teacherRepositoryImpl) RemoveSubjectsFromTeacher(teacherID int, subjectIDs []int) error {
	tx := r.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	for _, sid := range subjectIDs {
		err := tx.Exec(`
            DELETE FROM teacher_subjects
            WHERE teacher_id = ? AND subject_id = ?
        `, teacherID, sid).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	return tx.Commit().Error
}
