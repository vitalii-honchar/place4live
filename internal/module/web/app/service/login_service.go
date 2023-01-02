package service

import (
	"log"
	"place4live/internal/module/user/api"
	"place4live/internal/module/web/domain"
)

type LoginService struct {
	authService api.AuthUserService
	jwtService  *JwtService
}

func NewLoginService(authService api.AuthUserService, jwtService *JwtService) *LoginService {
	return &LoginService{authService: authService, jwtService: jwtService}
}

func (ls *LoginService) Login(username string, password string) (string, bool) {
	if res := <-ls.authService.Authorize(username, password); res.Ok {
		token := domain.NewJwtToken(res.UserId)
		if res, err := ls.jwtService.generateToken(token); err != nil {
			log.Printf("Unexpected error during token generation: %v\n", err)
		} else {
			return res, true
		}
	}
	return "", false
}
