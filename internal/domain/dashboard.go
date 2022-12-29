package domain

type Dashboard struct {
	Id     int64
	Cities map[int64]*DashboardCity
}

type DashboardCity struct {
	Order int64
	City
}

func (d *Dashboard) Add(c *DashboardCity) {
	if _, ok := d.Cities[c.Id]; !ok {
		d.Cities[c.Id] = c
	}
}

func (d *Dashboard) Remove(c *DashboardCity) {
	delete(d.Cities, c.Id)
}
