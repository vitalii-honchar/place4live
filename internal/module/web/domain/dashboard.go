package domain

type Dashboard struct {
	Id     int64
	Cities map[int64]UiCity
}

func (d *Dashboard) Add(c UiCity) {
	if _, ok := d.Cities[c.Id]; !ok {
		d.Cities[c.Id] = c
	}
}

func (d *Dashboard) Remove(c UiCity) {
	delete(d.Cities, c.Id)
}
