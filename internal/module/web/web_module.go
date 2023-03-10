package web

import (
	"place4live/internal/config"
	"place4live/internal/module/user"
	"place4live/internal/module/web/app/service"
	"place4live/internal/module/web/infrastructure/rest"
	"place4live/internal/module/web/infrastructure/rest/auth"
	"place4live/internal/module/web/infrastructure/rest/login"
	"place4live/internal/module/web/infrastructure/rest/register"
)

type WebModule struct {
	UserModule *user.UserModule
}

func (wm *WebModule) Init(ctx *config.AppContext) error {
	jwtTokenQueryInPort := service.NewJwtService(ctx.ApiSecret)

	loginService := service.NewLoginService(wm.UserModule.AuthUserService, jwtTokenQueryInPort)
	registerService := service.NewRegisterService(wm.UserModule.CreateUserService)
	ctx.ApiEngine.AddHandler(login.NewHandler(loginService))
	ctx.ApiEngine.AddHandler(register.NewHandler(registerService))
	ctx.ApiEngine.AddMiddleware(auth.JwtAuthMiddleware(jwtTokenQueryInPort))
	ctx.ApiEngine.AddHandler(rest.NewGetDashboardHandler())
	return nil
}

func (wm *WebModule) Name() string {
	return "web"
}
