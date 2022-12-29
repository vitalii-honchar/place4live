package usecase

import (
	"place4live/internal/application/port/out"
	"place4live/internal/domain"
)

type GetCityUseCase struct {
	outPort out.CityQueryOutPort
}

func NewGetCityUseCase(dbPort out.CityQueryOutPort) *GetCityUseCase {
	return &GetCityUseCase{outPort: dbPort}
}

func (gc *GetCityUseCase) getCity(name string) <-chan *domain.City {
	return gc.outPort.GetCity(name)
}
