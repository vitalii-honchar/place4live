package domain

import "time"

type City struct {
	Id        int64
	Name      string
	Prices    *CityPrices
	UpdatedAt time.Time
}

type CityPrices struct {
	Restaurants
	Markets
	Transportation
	Utilities
	Sports
	Childcare
	Clothing
	RentApartment
	BuyApartment
	Salaries
}

type Restaurants []*Property
type Markets []*Property
type Transportation []*Property
type Utilities []*Property
type Sports []*Property
type Childcare []*Property
type Clothing []*Property
type RentApartment []*Property
type BuyApartment []*Property
type Salaries []*Property

type Price struct {
	Avg float64
	Min float64
	Max float64
}

type Property struct {
	Name  string `json:"name"`
	Price `json:"price"`
}
