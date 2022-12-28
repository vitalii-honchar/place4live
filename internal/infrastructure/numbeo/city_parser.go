package numbeo

import (
	"github.com/PuerkitoBio/goquery"
	"io"
	"log"
	"place4live/internal/domain"
	"strings"
)

const tableRows = ".data_wide_table tr"
const classFirstCurrency = ".first_currency"
const classBarTextLeft = ".barTextLeft"
const classBarTextRight = ".barTextRight"
const td = "td"

func parseCity(name string, rc io.ReadCloser) *domain.City {
	doc, err := goquery.NewDocumentFromReader(rc)
	if err != nil {
		log.Printf("Unexpected error during read of document: err = %v\n", err)
		return nil
	}

	var city domain.City
	city.Name = name

	doc.Find(tableRows).Each(func(i int, s *goquery.Selection) {
		if td := s.Find(td).First(); td != nil && td.Text() != "" {
			var p domain.Property
			p.Name = strings.TrimSpace(td.Text())

			if pv := s.Find(classFirstCurrency).First(); pv != nil {
				p.Avg = strings.TrimSpace(pv.Text())
			}
			if pl := s.Find(classBarTextLeft).First(); pl != nil {
				p.Min = strings.TrimSpace(pl.Text())
			}
			if ph := s.Find(classBarTextRight).First(); ph != nil {
				p.Max = strings.TrimSpace(ph.Text())
			}
			city.Properties = append(city.Properties, &p)
		}
	})

	return &city
}
