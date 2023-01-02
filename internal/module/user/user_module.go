package user

import (
	"place4live/internal/config"
	"place4live/internal/module/user/api"
	"place4live/internal/module/user/app/usecase"
	internalApi "place4live/internal/module/user/infrastructure/api"
	"place4live/internal/module/user/infrastructure/database/repository"
	"place4live/internal/module/user/infrastructure/database/service"
)

type UserModule struct {
	AuthUserService api.AuthUserService
}

func (um *UserModule) Init(ctx *config.AppContext) error {
	rep := repository.NewUserRepository(ctx.Db)
	userOutPort := service.NewUserService(rep)
	userInPort := usecase.NewGetUserUseCase(userOutPort)

	um.AuthUserService = internalApi.NewAuthService(userInPort)
	return nil
}

func (um *UserModule) Name() string {
	return "user"
}
