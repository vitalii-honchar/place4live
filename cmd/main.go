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

var cityModule = &city.CityModule{}
var userModule = &user.UserModule{}
var webModule = &web.WebModule{UserModule: userModule}

var modules = []module{
	cityModule,
	userModule,
	webModule,
}

func main() {
	cfg, err := config.NewConfig()
	stopStartup("missed config", err)
	ctx, err := config.NewAppContext(cfg)
	stopStartup("failed to create app context", err)

	log.Printf("Started modules: %v\n", startModules(ctx))
	ctx.ApiEngine.Run()
}

func startModules(ctx *config.AppContext) []string {
	var started []string
	for _, m := range modules {
		err := m.Init(ctx)
		stopStartup(fmt.Sprintf("failed to init module '%s'", m.Name()), err)
		started = append(started, m.Name())
	}
	return started
}

func stopStartup(reason string, err error) {
	if err != nil {
		log.Fatalf("Can't start application: reason = %s, error = %s\n", reason, err)
	}
}
