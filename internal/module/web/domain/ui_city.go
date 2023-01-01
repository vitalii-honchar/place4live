package domain

type UiCity struct {
	Id             int64
	Order          int64
	Name           string
	Restaurants    []*Property
	Markets        []*Property
	Transportation []*Property
	Utilities      []*Property
	Sports         []*Property
	Childcare      []*Property
	Clothing       []*Property
	RentApartment  []*Property
	BuyApartment   []*Property
	Salaries       []*Property
}

type Property struct {
	Name  string `json:"name"`
	Price `json:"price"`
}

type Price struct {
	Avg float64 `json:"avg"`
	Min float64 `json:"min"`
	Max float64 `json:"max"`
}
