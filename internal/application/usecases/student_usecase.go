package usecases

import (
	"clean-architecture/internal/domain/entities"
	"clean-architecture/internal/domain/repositories"
)

type StudentUsecase interface {
	CreateStudent(student *entities.Student) error
	GetStudentByID(id int) (*entities.Student, error)
	UpdateStudent(student *entities.Student) error
	DeleteStudent(id int) error
	GetAllStudents() ([]*entities.Student, error)

	UpdateStudentAndAssignSubjects(studentID int, firstName, lastName string, subjectIDs []int) error
}

type studentUsecase struct {
	studentRepo repositories.StudentRepository
}

func NewStudentUsecase(sr repositories.StudentRepository) StudentUsecase {
	return &studentUsecase{studentRepo: sr}
}

func (u *studentUsecase) CreateStudent(student *entities.Student) error {
	return u.studentRepo.CreateStudent(student)
}

func (u *studentUsecase) GetStudentByID(id int) (*entities.Student, error) {
	return u.studentRepo.GetStudentByID(id)
}

func (u *studentUsecase) UpdateStudent(student *entities.Student) error {
	return u.studentRepo.UpdateStudent(student)
}

func (u *studentUsecase) DeleteStudent(id int) error {
	return u.studentRepo.DeleteStudent(id)
}

func (u *studentUsecase) GetAllStudents() ([]*entities.Student, error) {
	return u.studentRepo.GetAllStudents()
}

func (u *studentUsecase) UpdateStudentAndAssignSubjects(studentID int, firstName, lastName string, subjectIDs []int) error {
	student, err := u.studentRepo.GetStudentByID(studentID)
	if err != nil {
		return err
	}
	if student == nil {
		return nil // or return error("not found")
	}

	student.FirstName = firstName
	student.LastName = lastName

	if err := u.studentRepo.UpdateStudent(student); err != nil {
		return err
	}

	if err := u.studentRepo.AssignSubjectsToStudent(studentID, subjectIDs); err != nil {
		return err
	}
	return nil
}
