package port

import "place4live/internal/domain"

type CityQueryOutPort interface {
	GetCity(name string) <-chan *domain.City
}
