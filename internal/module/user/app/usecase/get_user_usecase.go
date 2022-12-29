package usecase

import (
	"place4live/internal/module/user/app/port"
	"place4live/internal/module/user/domain"
)

type GetUserUseCase struct {
	outPort port.UserQueryOutPort
}

func NewGetUserUseCase(outPort port.UserQueryOutPort) *GetUserUseCase {
	return &GetUserUseCase{outPort: outPort}
}

func (gu *GetUserUseCase) GetUser(username string) <-chan *domain.User {
	return gu.outPort.GetUser(username)
}
