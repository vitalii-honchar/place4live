package database

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

type CityRepositoryMock struct {
	mock.Mock
}

func (cqm *CityRepositoryMock) FindByName(name string) <-chan *domain.City {
	args := cqm.Called(name)
	return args.Get(0).(<-chan *domain.City)
}

func (cqm *CityRepositoryMock) Save(city *domain.City) <-chan bool {
	args := cqm.Called(city)
	return args.Get(0).(<-chan bool)
}

type CityQueryMock struct {
	mock.Mock
}

func (cqm *CityQueryMock) GetCity(name string) <-chan *domain.City {
	args := cqm.Called(name)
	return args.Get(0).(<-chan *domain.City)
}

func TestNewGetCityUseCase(t *testing.T) {
	t.Run("should call only database and numbeo port when cache was expired", func(t *testing.T) {
		// GIVEN
		repositoryMock := &CityRepositoryMock{}
		queryPort := &CityQueryMock{}
		name := fmt.Sprintf("City_%d", random.Intn(100))
		city := &domain.City{Name: name}
		svc := NewCityService(repositoryMock, queryPort)

		repositoryMock.On("FindByName", name).Return(testChannel(city))
		queryPort.On("GetCity", name).Return(testChannel(city))
		repositoryMock.On("Save", city).Return(testChannel(true))

		// WHEN
		actual := <-svc.GetCity(name)

		// THEN
		assert.Equal(t, city, actual)

		repositoryMock.AssertExpectations(t)
		queryPort.AssertExpectations(t)
	})
}

func testChannel[T any](t T) <-chan T {
	ch := make(chan T, 1)
	ch <- t
	return ch
}
