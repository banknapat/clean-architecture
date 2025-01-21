package repositories

import "clean-architecture/internal/domain/entities"

type UserRepository interface {
	CreateUser(user *entities.User) error
	GetUserByID(id int) (*entities.User, error)
	GetUserByUsername(username string) (*entities.User, error)
	UpdateUser(user *entities.User) error

	// เพิ่มเติมเพื่อให้รองรับ getAllUser
	GetAllUsers() ([]*entities.User, error)
}
