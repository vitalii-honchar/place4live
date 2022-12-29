package usecase

import (
	"place4live/internal/application/port/out"
	"place4live/internal/domain"
)

type GetUserUseCase struct {
	outPort out.UserQueryOutPort
}

func NewGetUserUseCase(outPort out.UserQueryOutPort) *GetUserUseCase {
	return &GetUserUseCase{outPort: outPort}
}

func (gu *GetUserUseCase) GetUser(username string) <-chan *domain.User {
	return gu.outPort.GetUser(username)
}
