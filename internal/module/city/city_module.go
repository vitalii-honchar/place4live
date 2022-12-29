package city

import (
	"place4live/internal/config"
	"place4live/internal/module/city/infrastructure/database/repository"
	"place4live/internal/module/city/infrastructure/database/service"
	"place4live/internal/module/city/infrastructure/numbeo"
)

type CityModule struct{}

func (cm *CityModule) Init(cfg *config.AppContext) error {
	cityRepository := repository.NewCityRepository(cfg.Db)
	numbeCityService := numbeo.NewCityQueryService()
	cityDbService := service.NewCityService(cityRepository, numbeCityService)

	return nil
}

func (cm *CityModule) Name() string {
	return "city"
}
