package usecase

import (
	"log"
	"place4live/internal/lib"
	"place4live/internal/module/user/app/port"
	"place4live/internal/module/user/domain"
)

type CreateUserUseCase struct {
	outPort port.UserCommandOutPort
}

func NewCreateUserUseCase(outPort port.UserCommandOutPort) *CreateUserUseCase {
	return &CreateUserUseCase{outPort: outPort}
}

func (cu *CreateUserUseCase) Create(username string, password string) <-chan bool {
	return lib.Async(func() bool {
		user, err := domain.NewUser(username, password)
		if err != nil {
			log.Printf("Unexpected error during creating user: %v\n", err)
			return false
		}
		return <-cu.outPort.Save(user)
	})
}
