package repositories

import "clean-architecture/internal/domain/entities"

type StudentRepository interface {
	CreateStudent(student *entities.Student) error
	GetStudentByID(id int) (*entities.Student, error)
	UpdateStudent(student *entities.Student) error
	DeleteStudent(id int) error
	GetAllStudents() ([]*entities.Student, error)

	AssignSubjectsToStudent(studentID int, subjectIDs []int) error
	RemoveSubjectsFromStudent(studentID int, subjectIDs []int) error
}
