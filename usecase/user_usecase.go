package usecase

import (
	"farras/integration-test-golang/model"
	"farras/integration-test-golang/repository"
)

type UserUseCase interface {
	GetUsers() []model.User
	Create(user model.User) model.User
}

type userUseCaseImpl struct {
	userRepository repository.UserRepository
}

func NewUserUseCase(userRepository repository.UserRepository) UserUseCase {
	return &userUseCaseImpl{
		userRepository,
	}
}

func (u *userUseCaseImpl) GetUsers() []model.User {
	return u.userRepository.GetUsers()
}

// Create implements UserUseCase.
func (u *userUseCaseImpl) Create(user model.User) model.User {
	return u.userRepository.Create(user)
}
