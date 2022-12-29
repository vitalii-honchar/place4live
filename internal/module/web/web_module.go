package web

import (
	"place4live/internal/config"
	"place4live/internal/module/web/infrastructure/rest"
)

type WebModule struct{}

func (wm *WebModule) Init(ctx *config.AppContext) error {
	ctx.ApiEngine.AddHandler(rest.NewGetDashboardHandler())
	return nil
}

func (wm *WebModule) Name() string {
	return "web"
}
