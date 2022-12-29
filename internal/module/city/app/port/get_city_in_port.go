package port

import "place4live/internal/module/city/domain"

type GetCityInPort interface {
	GetCity(name string) <-chan *domain.City
}
