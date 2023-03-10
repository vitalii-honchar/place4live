package rest

import (
	"place4live/internal/module/web/domain"
	"sort"
)

type uiDashboard struct {
	Id     int64    `json:"id"`
	Cities []uiCity `json:"cities"`
}

type uiCity struct {
	Order int64 `json:"order"`
	domain.UiCity
}

func newUiDashboard(d *domain.Dashboard) *uiDashboard {
	res := &uiDashboard{Id: d.Id}
	for _, city := range d.Cities {
		res.Cities = append(res.Cities, uiCity{Order: city.Order, UiCity: city})
	}
	sort.Slice(res.Cities, func(i, j int) bool {
		return res.Cities[i].Order < res.Cities[j].Order
	})
	return res
}
