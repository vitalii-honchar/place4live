package service

import (
	"log"
	"place4live/internal/application/port/out"
	"place4live/internal/domain"
	repository2 "place4live/internal/infrastructure/database/repository"
	"place4live/internal/lib"
	"time"
)

const day = 24 * time.Hour

type CityService struct {
	repository repository2.CityRepository
	queryPort  out.CityQueryOutPort
}

func NewCityService(r repository2.CityRepository, queryPort out.CityQueryOutPort) *CityService {
	return &CityService{repository: r, queryPort: queryPort}
}

func (cr *CityService) GetCity(name string) <-chan *domain.City {
	return lib.Async(func() *domain.City {
		city := <-cr.repository.FindByName(name)
		if city == nil || time.Now().In(time.UTC).Sub(city.UpdatedAt) >= day {
			log.Printf("Refreshing city db cache: name = %s\n", name)
			city = <-cr.queryPort.GetCity(name)
			<-cr.repository.Save(city)
		}
		return city
	})
}
