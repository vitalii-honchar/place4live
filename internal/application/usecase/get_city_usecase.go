package usecase

import (
	"place4live/internal/application/port/out"
	"place4live/internal/domain"
	"place4live/internal/lib"
	"time"
)

const day = 24 * time.Hour

type GetCityUseCase struct {
	dbPort      out.CityQueryOutPort
	numbeoPort  out.CityQueryOutPort
	commandPort out.CityCommandOutPort
}

func NewGetCityUseCase(dbPort out.CityQueryOutPort, numbeoPort out.CityQueryOutPort, commandPort out.CityCommandOutPort) *GetCityUseCase {
	return &GetCityUseCase{dbPort: dbPort, numbeoPort: numbeoPort, commandPort: commandPort}
}

func (gc *GetCityUseCase) getCity(name string) <-chan *domain.City {
	return lib.Async(func() *domain.City {
		city := <-gc.dbPort.GetCity(name)
		if time.Now().In(time.UTC).Sub(city.UpdatedAt) >= day {
			city = <-gc.numbeoPort.GetCity(name)
			<-gc.commandPort.Save(city)
		}
		return city
	})
}
