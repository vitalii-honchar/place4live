package numbeo

import (
	"place4live/internal/domain"
	"testing"
)

const mealInexpensiveRestaurant = "Meal, Inexpensive Restaurant"
const apartment1BedroomOutsideOfCenter = "Apartment (1 bedroom) Outside of Centre"
const apartment3BedroomOutsideOfCenter = "Apartment (3 bedrooms) Outside of Centre"
const mortgageInterestRate = "Mortgage Interest Rate in Percentages (%), Yearly, for 20 Years Fixed-Rate"

func TestGetCity(t *testing.T) {
	t.Run("should return city info when city exists", func(t *testing.T) {
		// GIVEN
		cn := "Toronto"

		// WHEN
		city := <-GetCity(cn)

		// THEN
		if city.Name != cn {
			t.Errorf("Name is unexpected: expected = %s, actual = %s\n", cn, city.Name)
		}
		props := getProperties(city)
		checkProperty(t, mealInexpensiveRestaurant, props)
		checkProperty(t, apartment1BedroomOutsideOfCenter, props)
		checkProperty(t, apartment3BedroomOutsideOfCenter, props)
		checkProperty(t, mortgageInterestRate, props)
	})

	t.Run("should return empty city when city not exists", func(t *testing.T) {
		// GIVEN
		cn := "Toronto1"

		// WHEN
		city := <-GetCity(cn)

		// THEN
		if city.Name != cn {
			t.Errorf("Name is unexpected: expected = %s, actual = %s\n", cn, city.Name)
		}
		if len(city.Properties) != 0 {
			t.Errorf("City properties should be empty: properties = %v\n", city.Properties)
		}
	})
}

func checkProperty(t *testing.T, name string, props map[string]*domain.Property) {
	if p, ok := props[name]; !ok {
		t.Errorf("Missed property: \"%s\"", name)
	} else {
		if p.Avg <= 0 {
			t.Errorf("Avg should be > 0: property = %s, avg = %f\n", name, p.Avg)
		}
		if p.Min <= 0 {
			t.Errorf("Min should be > 0: property = %s, min = %f\n", name, p.Min)
		}
		if p.Max <= 0 {
			t.Errorf("Max should be > 0: property = %s, max = %f\n", name, p.Max)
		}
	}
}

func getProperties(c *domain.City) map[string]*domain.Property {
	res := map[string]*domain.Property{}
	for _, p := range c.Properties {
		res[p.Name] = p
	}
	return res
}
