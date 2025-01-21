package repositories

import "clean-architecture/internal/domain/entities"

type TeacherRepository interface {
	CreateTeacher(teacher *entities.Teacher) error
	GetTeacherByID(id int) (*entities.Teacher, error)
	UpdateTeacher(teacher *entities.Teacher) error
	DeleteTeacher(id int) error
	GetAllTeachers() ([]*entities.Teacher, error)

	AssignSubjectsToTeacher(teacherID int, subjectIDs []int) error
	RemoveSubjectsFromTeacher(teacherID int, subjectIDs []int) error
}
