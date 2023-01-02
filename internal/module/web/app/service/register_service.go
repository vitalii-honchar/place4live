package service

import "place4live/internal/module/user/api"

type RegisterService struct {
	userService api.CreateUserService
}

func NewRegisterService(userService api.CreateUserService) *RegisterService {
	return &RegisterService{userService: userService}
}

func (rs *RegisterService) Register(username string, password string) bool {
	return <-rs.userService.Create(username, password)
}
