package domain

type Dashboard struct {
	Id     int64
	Cities map[string]*DashboardCity
}

type DashboardCity struct {
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

func (d *Dashboard) Add(c *DashboardCity) {
	if _, ok := d.Cities[c.Name]; !ok {
		d.Cities[c.Name] = c
	}
}

func (d *Dashboard) Remove(c *DashboardCity) {
	delete(d.Cities, c.Name)
}
