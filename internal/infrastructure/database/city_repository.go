package database

import (
	"database/sql"
	"log"
	"place4live/internal/domain"
	"place4live/internal/lib"
	"time"
)

const insertCity = "INSERT INTO city (name) VALUES ($1) ON CONFLICT(name) DO NOTHING RETURNING id"

const insertPrice = `
INSERT INTO city_price
	(city_id, category, name, avg, min, max, updated_at)
VALUES 
	($1, $2, $3, $4, $5, $6, $7)
ON CONFLICT (city_id, category, name)
DO UPDATE SET avg = $4, min = $5, max = $6, updated_at = $7 
`

const selectCityByName = `
	SELECT
	    c.id,
	    c.name,
	    cp.category,
	    cp.name,
	    cp.avg,
	    cp.min,
	    cp.max
	FROM city c
	INNER JOIN city_price cp ON cp.city_id = c.id
	WHERE c.name = $1
`

const (
	categoryRestaurants    = "RESTAURANTS"
	categoryMarkets        = "MARKETS"
	categoryTransportation = "TRANSPORTATION"
	categoryUtilities      = "UTILITIES"
	categorySports         = "SPORTS"
	categoryChildcare      = "CHILDCARE"
	categoryClothing       = "CLOTHING"
	categoryRentApartment  = "RENT_APARTMENT"
	categoryBuyApartment   = "BUY_APARTMENT"
	categorySalaries       = "SALARIES"
)

var categoryGetters = map[string]func(prices *domain.CityPrices) []*domain.Property{
	categoryRestaurants: func(cp *domain.CityPrices) []*domain.Property {
		return cp.Restaurants
	},
	categoryMarkets: func(cp *domain.CityPrices) []*domain.Property {
		return cp.Markets
	},
	categoryTransportation: func(cp *domain.CityPrices) []*domain.Property {
		return cp.Transportation
	},
	categoryUtilities: func(cp *domain.CityPrices) []*domain.Property {
		return cp.Utilities
	},
	categorySports: func(cp *domain.CityPrices) []*domain.Property {
		return cp.Sports
	},
	categoryChildcare: func(cp *domain.CityPrices) []*domain.Property {
		return cp.Childcare
	},
	categoryClothing: func(cp *domain.CityPrices) []*domain.Property {
		return cp.Clothing
	},
	categoryRentApartment: func(cp *domain.CityPrices) []*domain.Property {
		return cp.RentApartment
	},
	categoryBuyApartment: func(cp *domain.CityPrices) []*domain.Property {
		return cp.BuyApartment
	},
	categorySalaries: func(cp *domain.CityPrices) []*domain.Property {
		return cp.Salaries
	},
}

var categorySetters = map[string]func(cp *domain.CityPrices, p *domain.Property){
	categoryRestaurants: func(cp *domain.CityPrices, p *domain.Property) {
		cp.Restaurants = append(cp.Restaurants, p)
	},
	categoryMarkets: func(cp *domain.CityPrices, p *domain.Property) {
		cp.Markets = append(cp.Markets, p)
	},
	categoryTransportation: func(cp *domain.CityPrices, p *domain.Property) {
		cp.Transportation = append(cp.Transportation, p)
	},
	categoryUtilities: func(cp *domain.CityPrices, p *domain.Property) {
		cp.Utilities = append(cp.Utilities, p)
	},
	categorySports: func(cp *domain.CityPrices, p *domain.Property) {
		cp.Sports = append(cp.Sports, p)
	},
	categoryChildcare: func(cp *domain.CityPrices, p *domain.Property) {
		cp.Childcare = append(cp.Childcare, p)
	},
	categoryClothing: func(cp *domain.CityPrices, p *domain.Property) {
		cp.Clothing = append(cp.Clothing, p)
	},
	categoryRentApartment: func(cp *domain.CityPrices, p *domain.Property) {
		cp.RentApartment = append(cp.RentApartment, p)
	},
	categoryBuyApartment: func(cp *domain.CityPrices, p *domain.Property) {
		cp.BuyApartment = append(cp.BuyApartment, p)
	},
	categorySalaries: func(cp *domain.CityPrices, p *domain.Property) {
		cp.Salaries = append(cp.Salaries, p)
	},
}

type CityRepository struct {
	db *sql.DB
}

func NewCityRepository(db *sql.DB) *CityRepository {
	return &CityRepository{db}
}

func (cr *CityRepository) FindByName(name string) <-chan *domain.City {
	return lib.Async(func() *domain.City {
		tx, err := cr.db.Begin()
		if err != nil {
			log.Printf("Unexpected error during openning a transaction: name = %s, error = %v\n", name, err)
			return nil
		}
		defer tx.Rollback()

		rows, err := tx.Query(selectCityByName, name)
		if err != nil {
			log.Printf("Unexpected error during selecting city by name: name = %s, error = %v\n", name, err)
			return nil
		}
		defer rows.Close()
		city, err := readCity(rows)
		if err != nil {
			log.Printf("Unexepcted error during read a city from rows: name = %s, error = %v\n", name, err)
			return nil
		}
		return city
	})
}

func (cr *CityRepository) Save(city *domain.City) <-chan bool {
	return lib.Async(func() bool {
		tx, err := cr.db.Begin()
		if err != nil {
			log.Printf("Unexpected error during openning a transaction: city = %+v, error = %v\n", city, err)
			return false
		}
		defer tx.Rollback()
		cityId, err := saveCityName(tx, city.Name)
		if err != nil {
			log.Printf("Unexpected error during saving city name: city = %+v, error = %v\n", city, err)
			return false
		}

		if err := savePrices(tx, cityId, city); err != nil {
			log.Printf("Unexpected error during updating categories: city = %+v, error = %v\n", city, err)
			return false
		}

		if err := tx.Commit(); err != nil {
			log.Printf("Unexpected error dirung commit a trasnaction: city = %+v, error = %v\n", city, err)
			return false
		}
		return true
	})
}

func saveCityName(tx *sql.Tx, name string) (int64, error) {
	var id int64
	err := tx.QueryRow(insertCity, name).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func savePrices(tx *sql.Tx, cityId int64, city *domain.City) error {
	now := time.Now().In(time.UTC)
	if city.Id != 0 {
		cityId = city.Id
	}

	for category, getter := range categoryGetters {
		err := createCategoryPrice(tx, now, cityId, category, getter(city.Prices))
		if err != nil {
			return err
		}
	}
	return nil
}

func createCategoryPrice(tx *sql.Tx, now time.Time, cityId int64, name string, props []*domain.Property) error {
	for _, p := range props {
		_, err := tx.Exec(insertPrice, cityId, name, p.Name, p.Avg, p.Min, p.Max, now)
		if err != nil {
			return err
		}
	}
	return nil
}

func readCity(r *sql.Rows) (*domain.City, error) {
	city := domain.City{
		Prices: &domain.CityPrices{},
	}
	for r.Next() {
		var category string
		var property domain.Property
		err := r.Scan(&city.Id, &city.Name, &category, &property.Name, &property.Avg, &property.Min, &property.Max)
		if err != nil {
			return nil, err
		}
		if setter, ok := categorySetters[category]; ok {
			setter(city.Prices, &property)
		} else {
			log.Printf("Missed setter for property: category = %s, property = %v, city = %s\n", category, property, city.Name)
		}
	}
	return &city, nil
}
