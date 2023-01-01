package api

type City struct {
	Id             int64
	Name           string
	Restaurants    []Property
	Markets        []Property
	Transportation []Property
	Utilities      []Property
	Sports         []Property
	Childcare      []Property
	Clothing       []Property
	RentApartment  []Property
	BuyApartment   []Property
	Salaries       []Property
}

type Property struct {
	Name  string `json:"name"`
	Price `json:"price"`
}

type Price struct {
	Avg float64
	Min float64
	Max float64
}

type CityService interface {
	GetCity(name string) <-chan City
}
