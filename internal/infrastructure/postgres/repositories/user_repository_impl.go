package repositories

import (
	"errors"

	"gorm.io/gorm"

	"clean-architecture/internal/domain/entities"
	"clean-architecture/internal/domain/repositories"
)

type userRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repositories.UserRepository {
	return &userRepositoryImpl{db}
}

func (r *userRepositoryImpl) CreateUser(user *entities.User) error {
	return r.db.Create(user).Error
}

func (r *userRepositoryImpl) GetUserByID(id int) (*entities.User, error) {
	var user entities.User
	err := r.db.First(&user, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *userRepositoryImpl) GetUserByUsername(username string) (*entities.User, error) {
	var user entities.User
	err := r.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *userRepositoryImpl) UpdateUser(user *entities.User) error {
	return r.db.Save(user).Error
}

func (r *userRepositoryImpl) GetAllUsers() ([]*entities.User, error) {
	var users []entities.User
	err := r.db.Find(&users).Error
	if err != nil {
		return nil, err
	}
	// ต้อง return slice ของ pointer ถ้าจะตรงตาม interface sign หรือแล้วแต่ตกลง
	// แต่ในตัวอย่างจะรีเทิร์น slice ของ value ก็ถือว่ารับได้
	usersPtr := make([]*entities.User, len(users))
	for i := range users {
		usersPtr[i] = &users[i]
	}
	return usersPtr, nil
}
