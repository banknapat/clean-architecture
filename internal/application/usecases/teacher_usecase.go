package usecases

import (
	"clean-architecture/internal/domain/entities"
	"clean-architecture/internal/domain/repositories"
)

type TeacherUsecase interface {
	CreateTeacher(teacher *entities.Teacher) error
	GetTeacherByID(id int) (*entities.Teacher, error)
	UpdateTeacher(teacher *entities.Teacher) error
	DeleteTeacher(id int) error
	GetAllTeachers() ([]*entities.Teacher, error)

	AssignSubjectsToTeacher(teacherID int, subjectIDs []int) error
}

type teacherUsecase struct {
	teacherRepo repositories.TeacherRepository
}

func NewTeacherUsecase(tr repositories.TeacherRepository) TeacherUsecase {
	return &teacherUsecase{teacherRepo: tr}
}

func (u *teacherUsecase) CreateTeacher(teacher *entities.Teacher) error {
	return u.teacherRepo.CreateTeacher(teacher)
}

func (u *teacherUsecase) GetTeacherByID(id int) (*entities.Teacher, error) {
	return u.teacherRepo.GetTeacherByID(id)
}

func (u *teacherUsecase) UpdateTeacher(teacher *entities.Teacher) error {
	return u.teacherRepo.UpdateTeacher(teacher)
}

func (u *teacherUsecase) DeleteTeacher(id int) error {
	return u.teacherRepo.DeleteTeacher(id)
}

func (u *teacherUsecase) GetAllTeachers() ([]*entities.Teacher, error) {
	return u.teacherRepo.GetAllTeachers()
}

func (u *teacherUsecase) AssignSubjectsToTeacher(teacherID int, subjectIDs []int) error {
	return u.teacherRepo.AssignSubjectsToTeacher(teacherID, subjectIDs)
}
