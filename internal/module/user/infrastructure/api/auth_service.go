package api

import (
	"place4live/internal/lib"
	"place4live/internal/module/user/api"
	"place4live/internal/module/user/app/port"
)

type AuthService struct {
	inPort port.GetUserInPort
}

func NewAuthService(inPort port.GetUserInPort) *AuthService {
	return &AuthService{inPort: inPort}
}

func (au *AuthService) Authorize(username string, password string) <-chan api.AuthResult {
	return lib.Async(func() api.AuthResult {
		user := <-au.inPort.GetUser(username)
		if user != nil && user.CheckPassword(password) {
			return api.AuthResult{Ok: true, UserId: user.Id}
		}
		return api.AuthResult{}
	})
}
