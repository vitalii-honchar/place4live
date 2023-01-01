package user

import "place4live/internal/config"

type UserModule struct{}

func (um *UserModule) Init(ctx *config.AppContext) error {
	return nil
}

func (um *UserModule) Name() string {
	return "user"
}
