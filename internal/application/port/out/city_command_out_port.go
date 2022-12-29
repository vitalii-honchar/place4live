package out

import "place4live/internal/domain"

type CityCommandOutPort interface {
	Save(city *domain.City) <-chan bool
}
