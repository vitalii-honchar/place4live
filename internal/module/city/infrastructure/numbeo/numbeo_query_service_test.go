package numbeo

import (
	"place4live/internal/module/city/domain"
	"testing"
)

const averageMonthlyNetSalary = "Average Monthly Net Salary (After Tax)"

func TestGetCity(t *testing.T) {
	t.Run("should return dashboard info when dashboard exists", func(t *testing.T) {
		// GIVEN
		cn := "Toronto"

		// WHEN
		city := <-GetCity(cn)

		// THEN
		if city.Name != cn {
			t.Errorf("Name is unexpected: expected = %s, actual = %s\n", cn, city.Name)
		}
		checkPrices(t, city.Prices.Restaurants)
		checkPrices(t, city.Prices.Markets)
		checkPrices(t, city.Prices.Transportation)
		checkPrices(t, city.Prices.Utilities)
		checkPrices(t, city.Prices.Sports)
		checkPrices(t, city.Prices.Childcare)
		checkPrices(t, city.Prices.Clothing)
		checkPrices(t, city.Prices.RentApartment)
		checkPrices(t, city.Prices.BuyApartment)
		checkPrices(t, city.Prices.Salaries)
	})

	t.Run("should return empty dashboard when dashboard not exists", func(t *testing.T) {
		// GIVEN
		cn := "Toronto1"

		// WHEN
		city := <-GetCity(cn)

		// THEN
		if city.Name != cn {
			t.Errorf("Name is unexpected: expected = %s, actual = %s\n", cn, city.Name)
		}
		if city.Prices != nil {
			t.Errorf("City prices should be nil: prices = %v\n", city.Prices)
		}
	})
}

func checkPrices(t *testing.T, props []*domain.Property) {
	for _, p := range props {
		checkProperty(t, p)
	}
}

func checkProperty(t *testing.T, p *domain.Property) {
	if p.Avg <= 0 {
		t.Errorf("Avg should be > 0: property = %+v\n", p)
	}
	if p.Name != averageMonthlyNetSalary {
		if p.Min <= 0 {
			t.Errorf("Min should be > 0: property = %+v\n", p)
		}
		if p.Max <= 0 {
			t.Errorf("Max should be > 0: property = %+v\n", p)
		}
	}
}
