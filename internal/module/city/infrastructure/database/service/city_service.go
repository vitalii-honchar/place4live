package service

import (
	"log"
	"place4live/internal/lib"
	"place4live/internal/module/city/app/port"
	"place4live/internal/module/city/domain"
	"place4live/internal/module/city/infrastructure/database/repository"
	"time"
)

const day = 24 * time.Hour

type CityService struct {
	cityRepository repository.CityRepository
	queryPort      port.CityQueryOutPort
}

func NewCityService(r repository.CityRepository, queryPort port.CityQueryOutPort) *CityService {
	return &CityService{cityRepository: r, queryPort: queryPort}
}

func (cr *CityService) GetCity(name string) <-chan *domain.City {
	return lib.Async(func() *domain.City {
		city := <-cr.cityRepository.FindByName(name)
		if city == nil || time.Now().In(time.UTC).Sub(city.UpdatedAt) >= day {
			log.Printf("Refreshing city db cache: name = %s\n", name)
			city = <-cr.queryPort.GetCity(name)
			<-cr.cityRepository.Save(city)
		}
		return city
	})
}
