package main

import (
	"fmt"
	"log"
	"place4live/internal/config"
)

type module interface {
	Start(cfg *config.AppContext) error
	Name() string
}

var modules = []module{}

//func createInPorts(db *sql.DB) *inPorts {
//	cityRepository := repository.NewCityRepository(db)
//	numbeCityService := numbeo.NewCityQueryService()
//	cityDbService := service.NewCityService(cityRepository, numbeCityService)
//
//	return &inPorts{
//		getCityInPort: usecase.NewGetCityUseCase(cityDbService),
//	}
//}
//
//func createHandlers(ports *inPorts) []web.ApiHandler {
//	return []web.ApiHandler{
//		dashboard.NewGetDashboardHandler(ports.getCityInPort),
//	}
//}

func main() {
	cfg, err := config.NewConfig()
	stopStartup("missed config", err)
	ctx, err := config.NewAppContext(cfg)
	stopStartup("failed to create app context", err)

	for _, m := range modules {
		err := m.Start(ctx)
		stopStartup(fmt.Sprintf("failed to start module '%s'", m.Name()), err)
	}
	ctx.ApiEngine.Run()
}

func stopStartup(reason string, err error) {
	if err != nil {
		log.Fatalf("Can't start application: reason = %s, error = %s\n", reason, err)
	}
}
