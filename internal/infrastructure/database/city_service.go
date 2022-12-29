package database

import (
	"place4live/internal/application/port/out"
	"place4live/internal/domain"
	"place4live/internal/lib"
	"time"
)

const day = 24 * time.Hour

type CityService struct {
	repository CityRepository
	queryPort  out.CityQueryOutPort
}

func NewCityService(r CityRepository, queryPort out.CityQueryOutPort) *CityService {
	return &CityService{repository: r, queryPort: queryPort}
}

func (cr *CityService) GetCity(name string) <-chan *domain.City {
	return lib.Async(func() *domain.City {
		city := <-cr.repository.FindByName(name)
		if city == nil || time.Now().In(time.UTC).Sub(city.UpdatedAt) >= day {
			city = <-cr.queryPort.GetCity(name)
			<-cr.repository.Save(city)
		}
		return city
	})
}
