package database

import "place4live/internal/domain"

type CityService struct {
	repository *CityRepository
}

func NewCityService(r *CityRepository) *CityService {
	return &CityService{repository: r}
}

func (cr *CityService) GetCity(name string) <-chan *domain.City {
	return cr.repository.FindByName(name)
}
