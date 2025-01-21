package usecases

import (
	"clean-architecture/internal/domain/entities"
	"clean-architecture/internal/domain/repositories"
)

type UserUsecase interface {
	GetUserByID(id int) (*entities.User, error)
	GetAllUsers() ([]*entities.User, error)
}

type userUsecase struct {
	userRepo repositories.UserRepository
}

func NewUserUsecase(ur repositories.UserRepository) UserUsecase {
	return &userUsecase{userRepo: ur}
}

func (u *userUsecase) GetUserByID(id int) (*entities.User, error) {
	return u.userRepo.GetUserByID(id)
}

func (u *userUsecase) GetAllUsers() ([]*entities.User, error) {
	return u.userRepo.GetAllUsers()
}
