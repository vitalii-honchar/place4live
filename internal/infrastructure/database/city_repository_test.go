package database

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"place4live/internal/domain"
	"testing"
	"time"
)

var seededRand = rand.New(rand.NewSource(time.Now().UnixNano()))

func TestCityRepository_Save(t *testing.T) {
	repository := NewCityRepository(testConnection(t))

	t.Run("should find city when it was saved", func(t *testing.T) {
		// GIVEN
		city := randomCity()

		// WHEN
		saved := <-repository.Save(city)

		// THEN
		assert.True(t, saved)

		// AND WHEN
		actual := <-repository.FindByName(city.Name)

		// AND THEN
		assert.NotNil(t, actual)
		assert.Equal(t, city.Name, actual.Name)
		assert.Greater(t, actual.Id, int64(0))
		assert.Equal(t, city.Prices, actual.Prices)
	})

	t.Run("should doesn't find city when it wasn't saved", func(t *testing.T) {
		// GIVEN

		// AND WHEN
		actual := <-repository.FindByName(fmt.Sprintf("test_%d", seededRand.Intn(100)))

		// AND THEN
		assert.Nil(t, actual)
	})
}

func randomCity() *domain.City {
	return &domain.City{
		Name: fmt.Sprintf("City_%d", seededRand.Intn(100)),
		Prices: &domain.CityPrices{
			Restaurants:    domain.Restaurants{randomProperty()},
			Markets:        domain.Markets{randomProperty()},
			Transportation: domain.Transportation{randomProperty()},
			Utilities:      domain.Utilities{randomProperty()},
			Sports:         domain.Sports{randomProperty()},
			Childcare:      domain.Childcare{randomProperty()},
			Clothing:       domain.Clothing{randomProperty()},
			RentApartment:  domain.RentApartment{randomProperty()},
			BuyApartment:   domain.BuyApartment{randomProperty()},
			Salaries:       domain.Salaries{randomProperty()},
		},
	}
}

func randomProperty() *domain.Property {
	return &domain.Property{
		Name: fmt.Sprintf("Property_%d", seededRand.Intn(100)),
		Price: domain.Price{
			Avg: float64(seededRand.Intn(100)),
			Min: float64(seededRand.Intn(100)),
			Max: float64(seededRand.Intn(100)),
		},
	}
}
