package api

import (
	"place4live/internal/lib"
	"place4live/internal/module/user/api"
	"place4live/internal/module/user/app/port"
)

type UserService struct {
	getUserInPort    port.GetUserInPort
	createUserInPort port.CreateUserInPort
}

func NewUserService(inPort port.GetUserInPort, createUserInPort port.CreateUserInPort) *UserService {
	return &UserService{getUserInPort: inPort, createUserInPort: createUserInPort}
}

func (us *UserService) Authorize(username string, password string) <-chan api.AuthResult {
	return lib.Async(func() api.AuthResult {
		user := <-us.getUserInPort.GetUser(username)
		if user != nil && user.CheckPassword(password) {
			return api.AuthResult{Ok: true, UserId: user.Id}
		}
		return api.AuthResult{}
	})
}

func (us *UserService) Create(username string, password string) <-chan bool {
	return us.createUserInPort.Create(username, password)
}
