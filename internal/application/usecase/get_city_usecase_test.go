package usecase

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"math/rand"
	"place4live/internal/domain"
	"testing"
	"time"
)

var random = rand.New(rand.NewSource(time.Now().UnixNano()))

type CityQueryMock struct {
	mock.Mock
}

func (cqm *CityQueryMock) GetCity(name string) <-chan *domain.City {
	args := cqm.Called(name)
	return args.Get(0).(<-chan *domain.City)
}

type CityCommandMock struct {
	mock.Mock
}

func (ccm *CityCommandMock) Save(city *domain.City) <-chan bool {
	args := ccm.Called(city)
	return args.Get(0).(<-chan bool)
}

func TestNewGetCityUseCase(t *testing.T) {
	t.Run("should call only database and numbeo port when cache was expired", func(t *testing.T) {
		// GIVEN
		dbPort := &CityQueryMock{}
		numbeoPort := &CityQueryMock{}
		commandPort := &CityCommandMock{}
		name := fmt.Sprintf("City_%d", random.Intn(100))
		city := &domain.City{Name: name}
		useCase := NewGetCityUseCase(dbPort, numbeoPort, commandPort)

		dbPort.On("GetCity", name).Return(testChannel(city))
		numbeoPort.On("GetCity", name).Return(testChannel(city))
		commandPort.On("Save", city).Return(testChannel(true))

		// WHEN
		actual := <-useCase.getCity(name)

		// THEN
		assert.Equal(t, city, actual)

		dbPort.AssertExpectations(t)
		numbeoPort.AssertExpectations(t)
		commandPort.AssertExpectations(t)
	})
}

func testChannel[T any](t T) <-chan T {
	ch := make(chan T, 1)
	ch <- t
	return ch
}
