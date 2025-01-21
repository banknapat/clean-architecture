package repositories

import (
	"gorm.io/gorm"

	"clean-architecture/internal/domain/entities"
	"clean-architecture/internal/domain/repositories"
)

type studentRepositoryImpl struct {
	db *gorm.DB
}

func NewStudentRepository(db *gorm.DB) repositories.StudentRepository {
	return &studentRepositoryImpl{db}
}

func (r *studentRepositoryImpl) CreateStudent(student *entities.Student) error {
	return r.db.Create(student).Error
}

func (r *studentRepositoryImpl) GetStudentByID(id int) (*entities.Student, error) {
	var student entities.Student
	err := r.db.First(&student, id).Error
	if err != nil {
		return nil, err
	}
	return &student, nil
}

func (r *studentRepositoryImpl) UpdateStudent(student *entities.Student) error {
	return r.db.Save(student).Error
}

func (r *studentRepositoryImpl) DeleteStudent(id int) error {
	return r.db.Delete(&entities.Student{}, id).Error
}

func (r *studentRepositoryImpl) GetAllStudents() ([]*entities.Student, error) {
	var students []*entities.Student
	err := r.db.Find(&students).Error
	return students, err
}

// Assign subjects to student
func (r *studentRepositoryImpl) AssignSubjectsToStudent(studentID int, subjectIDs []int) error {
	tx := r.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	for _, sid := range subjectIDs {
		// INSERT INTO student_subjects (student_id, subject_id) VALUES ...
		err := tx.Exec(`
            INSERT INTO student_subjects (student_id, subject_id)
            VALUES (?, ?)
            ON CONFLICT DO NOTHING
        `, studentID, sid).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	return tx.Commit().Error
}

// Remove subjects from student
func (r *studentRepositoryImpl) RemoveSubjectsFromStudent(studentID int, subjectIDs []int) error {
	tx := r.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	for _, sid := range subjectIDs {
		// DELETE FROM student_subjects ...
		err := tx.Exec(`
            DELETE FROM student_subjects
            WHERE student_id = ? AND subject_id = ?
        `, studentID, sid).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	return tx.Commit().Error
}
