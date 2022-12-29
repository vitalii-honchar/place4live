package service

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"math/rand"
	"place4live/internal/module/city/domain"
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

func TestCityService_GetCity(t *testing.T) {
	name := fmt.Sprintf("City_%d", random.Intn(100))
	cities := []*domain.City{
		{Name: name},
		{Name: name, UpdatedAt: time.Now().Add(-24 * time.Hour)},
		{Name: name, UpdatedAt: time.Now().Add(-48 * time.Hour)},
		nil,
	}
	for _, city := range cities {
		t.Run(fmt.Sprintf("GetCity if dashboard %+v returns dashboard and saves it", city), func(t *testing.T) {
			// GIVEN
			repositoryMock := &CityRepositoryMock{}
			queryPort := &CityQueryMock{}
			queried := &domain.City{Name: name, UpdatedAt: time.Now()}
			svc := NewCityService(repositoryMock, queryPort)

			repositoryMock.On("FindByName", name).Return(testChannel(city))
			queryPort.On("GetCity", name).Return(testChannel(queried))
			repositoryMock.On("Save", queried).Return(testChannel(true))

			// WHEN
			actual := <-svc.GetCity(name)

			// THEN
			assert.Equal(t, queried, actual)

			repositoryMock.AssertExpectations(t)
			queryPort.AssertExpectations(t)
		})
	}

	t.Run("GetCity if dashboard exists and not expired returns dashboard without saving it", func(t *testing.T) {
		// GIVEN
		repositoryMock := &CityRepositoryMock{}
		queryPort := &CityQueryMock{}
		city := &domain.City{Name: name, UpdatedAt: time.Now().Add(-1 * time.Hour)}
		svc := NewCityService(repositoryMock, queryPort)

		repositoryMock.On("FindByName", name).Return(testChannel(city))

		// WHEN
		actual := <-svc.GetCity(name)

		// THEN
		assert.Equal(t, city, actual)

		repositoryMock.AssertExpectations(t)
		queryPort.AssertNotCalled(t, "GetCity", name)
		repositoryMock.AssertNotCalled(t, "Save", city)
	})
}

func testChannel[T any](t T) <-chan T {
	ch := make(chan T, 1)
	ch <- t
	return ch
}
