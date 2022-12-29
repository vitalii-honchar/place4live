package numbeo

import (
	"fmt"
	"log"
	"net/http"
	"place4live/internal/domain"
)

const numbeoTemplate = "https://www.numbeo.com/cost-of-living/in/%s"

func GetCity(name string) <-chan *domain.City {
	c := make(chan *domain.City, 1)
	go func() {
		defer close(c)
		url := fmt.Sprintf(numbeoTemplate, name)

		r, err := http.Get(url)
		if err != nil {
			log.Printf("Unexpected error during read information about dashboard: url = %s, err = %v\n", url, err)
		} else {
			defer r.Body.Close()
			c <- parseCity(name, r.Body)
		}
	}()
	return c
}

type CityQueryService struct{}

func NewCityQueryService() *CityQueryService {
	return &CityQueryService{}
}

func (cq *CityQueryService) GetCity(name string) <-chan *domain.City {
	c := make(chan *domain.City, 1)
	go func() {
		defer close(c)
		url := fmt.Sprintf(numbeoTemplate, name)

		r, err := http.Get(url)
		if err != nil {
			log.Printf("Unexpected error during read information about dashboard: url = %s, err = %v\n", url, err)
		} else {
			defer r.Body.Close()
			c <- parseCity(name, r.Body)
		}
	}()
	return c
}
