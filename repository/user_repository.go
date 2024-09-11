package repository

import (
	"farras/integration-test-golang/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetUsers() []model.User
	Create(user model.User) model.User
}

type userRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepositoryImpl{
		db,
	}
}

// GetUsers implements UserRepository.
func (u *userRepositoryImpl) GetUsers() []model.User {
	var users []model.User
	u.db.Find(&users)

	return users
}

// Create implements UserRepository.
func (u *userRepositoryImpl) Create(user model.User) model.User {
	u.db.Create(&user)

	return user
}
