package city

import (
	"place4live/internal/config"
)

type CityModule struct{}

func (cm *CityModule) Init(cfg *config.AppContext) error {
	//cityRepository := repository.NewCityRepository(cfg.Db)
	//numbeCityService := numbeo.NewCityQueryService()
	//_ := service.NewCityService(cityRepository, numbeCityService)

	return nil
}

func (cm *CityModule) Name() string {
	return "city"
}
