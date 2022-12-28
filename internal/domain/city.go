package domain

type City struct {
	Name   string      `json:"name"`
	Prices *CityPrices `json:"prices"`
}

type CityPrices struct {
	Restaurants    `json:"restaurants"`
	Markets        `json:"markets"`
	Transportation `json:"transportation"`
	Utilities      `json:"utilities"`
	Sports         `json:"sports"`
	Childcare      `json:"childcare"`
	Clothing       `json:"clothing"`
	RentApartment  `json:"rentApartment"`
	BuyApartment   `json:"buyApartment"`
	Salaries       `json:"salaries"`
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
	Avg float64 `json:"avg"`
	Min float64 `json:"min"`
	Max float64 `json:"max"`
}

type Property struct {
	Name  string `json:"name"`
	Price `json:"price"`
}
