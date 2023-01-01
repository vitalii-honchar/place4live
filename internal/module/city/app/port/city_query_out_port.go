package port

import "place4live/internal/module/city/domain"

type CityQueryOutPort interface {
	GetCity(name string) <-chan *domain.City
}
