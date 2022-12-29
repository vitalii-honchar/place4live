package port

import "place4live/internal/domain"

type GetCityInPort interface {
	GetCity(name string) <-chan *domain.City
}
