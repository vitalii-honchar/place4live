package main

import (
	"database/sql"
	"log"
	"place4live/internal/application/port/in"
	"place4live/internal/application/usecase"
	"place4live/internal/infrastructure/config"
	"place4live/internal/infrastructure/database/postgres"
	"place4live/internal/infrastructure/database/repository"
	"place4live/internal/infrastructure/database/service"
	"place4live/internal/infrastructure/numbeo"
	"place4live/internal/infrastructure/web"
	"place4live/internal/infrastructure/web/dashboard"
)

type inPorts struct {
	getCityInPort in.GetCityInPort
}

func createInPorts(cfg *config.Config, db *sql.DB) *inPorts {
	cityRepository := repository.NewCityRepository(db)
	numbeCityService := numbeo.NewCityQueryService()
	cityDbService := service.NewCityService(cityRepository, numbeCityService)

	return &inPorts{
		getCityInPort: usecase.NewGetCityUseCase(cityDbService),
	}
}

func createHandlers(ports *inPorts) []web.ApiHandler {
	return []web.ApiHandler{
		dashboard.NewGetDashboardHandler(ports.getCityInPort),
	}
}

func main() {
	cfg, err := config.NewConfig()
	stopStartup("missed config", err)

	db, err := postgres.OpenConnection(cfg.DbConnectionStr, cfg.DbMigrationsFolder)
	stopStartup("failed open db connection", err)

	ports := createInPorts(cfg, db)
	handlers := createHandlers(ports)
	engine := web.NewApiEngine(handlers)

	engine.Run()
}

func stopStartup(reason string, err error) {
	if err != nil {
		log.Fatalf("Can't start application: reason = %s, error = %s\n", reason, err)
	}
}
