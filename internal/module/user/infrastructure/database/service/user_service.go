package service

import (
	"place4live/internal/module/user/domain"
	"place4live/internal/module/user/infrastructure/database/repository"
)

type UserService struct {
	userRepository repository.UserRepository
}

func NewUserService(rep repository.UserRepository) *UserService {
	return &UserService{userRepository: rep}
}

func (us *UserService) GetUser(username string) <-chan *domain.User {
	return us.userRepository.FindByUsername(username)
}

func (us *UserService) Save(user *domain.User) <-chan bool {
	return us.userRepository.Save(user)
}
