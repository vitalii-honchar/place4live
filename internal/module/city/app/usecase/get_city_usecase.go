package usecase

import (
	"place4live/internal/module/city/app/port"
	"place4live/internal/module/city/domain"
)

type GetCityUseCase struct {
	outPort port.CityQueryOutPort
}

func NewGetCityUseCase(dbPort port.CityQueryOutPort) *GetCityUseCase {
	return &GetCityUseCase{outPort: dbPort}
}

func (gc *GetCityUseCase) GetCity(name string) <-chan *domain.City {
	return gc.outPort.GetCity(name)
}
