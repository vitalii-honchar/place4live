package rest

import (
	"github.com/stretchr/testify/assert"
	"place4live/internal/domain"
	"testing"
)

func TestNewUiDashboard(t *testing.T) {
	// GIVEN
	dashboard := &domain.Dashboard{
		Id: 10,
		Cities: map[int64]*domain.DashboardCity{
			1: {Order: 10, City: domain.City{Id: 1, Name: "Toronto"}},
			2: {Order: 7, City: domain.City{Id: 2, Name: "Calgary"}},
			3: {Order: 5, City: domain.City{Id: 3, Name: "Kyiv"}},
			4: {Order: 23, City: domain.City{Id: 4, Name: "Edmonton"}},
		},
	}

	// WHEN
	actual := newUiDashboard(dashboard)

	// THEN
	assert.Equal(t, dashboard, convertToDomain(actual))
}

func convertToDomain(d *uiDashboard) *domain.Dashboard {
	res := &domain.Dashboard{
		Id:     d.Id,
		Cities: make(map[int64]*domain.DashboardCity),
	}
	for _, city := range d.Cities {
		res.Cities[city.Id] = &domain.DashboardCity{
			Order: city.Order,
			City:  city.City,
		}
	}
	return res
}
