package main

import (
	"fmt"
	"log"
	"place4live/internal/config"
	"place4live/internal/module/city"
	"place4live/internal/module/user"
	"place4live/internal/module/web"
)

type module interface {
	Init(ctx *config.AppContext) error
	Name() string
}

var modules = []module{
	&city.CityModule{},
	&user.UserModule{},
	&web.WebModule{},
}

func main() {
	cfg, err := config.NewConfig()
	stopStartup("missed config", err)
	ctx, err := config.NewAppContext(cfg)
	stopStartup("failed to create app context", err)

	for _, m := range modules {
		err := m.Init(ctx)
		stopStartup(fmt.Sprintf("failed to init module '%s'", m.Name()), err)
	}
	ctx.ApiEngine.Run()
}

func stopStartup(reason string, err error) {
	if err != nil {
		log.Fatalf("Can't start application: reason = %s, error = %s\n", reason, err)
	}
}
