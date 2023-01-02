package user

import (
	"place4live/internal/config"
	"place4live/internal/module/user/api"
)

type UserModule struct {
	AuthUserService api.AuthUserService
}

func (um *UserModule) Init(ctx *config.AppContext) error {
	return nil
}

func (um *UserModule) Name() string {
	return "user"
}
