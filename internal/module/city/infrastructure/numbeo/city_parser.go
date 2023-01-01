package numbeo

import (
	"github.com/PuerkitoBio/goquery"
	"io"
	"log"
	"place4live/internal/module/city/domain"
	"strconv"
	"strings"
)

const tableRows = ".data_wide_table tr"
const classFirstCurrency = ".first_currency"
const classBarTextLeft = ".barTextLeft"
const classBarTextRight = ".barTextRight"
const categorySelector = ".category_title"
const td = "td"

var priceUpdaters = map[string]func(*domain.CityPrices, *domain.Property){
	"Restaurants": func(c *domain.CityPrices, p *domain.Property) {
		c.Restaurants = append(c.Restaurants, p)
	},
	"Markets": func(c *domain.CityPrices, p *domain.Property) {
		c.Markets = append(c.Markets, p)
	},
	"Transportation": func(c *domain.CityPrices, p *domain.Property) {
		c.Transportation = append(c.Transportation, p)
	},
	"Utilities (Monthly)": func(c *domain.CityPrices, p *domain.Property) {
		c.Utilities = append(c.Utilities, p)
	},
	"Sports And Leisure": func(c *domain.CityPrices, p *domain.Property) {
		c.Sports = append(c.Sports, p)
	},
	"Childcare": func(c *domain.CityPrices, p *domain.Property) {
		c.Childcare = append(c.Childcare, p)
	},
	"Clothing And Shoes": func(c *domain.CityPrices, p *domain.Property) {
		c.Clothing = append(c.Clothing, p)
	},
	"Rent Per Month": func(c *domain.CityPrices, p *domain.Property) {
		c.RentApartment = append(c.RentApartment, p)
	},
	"Buy Apartment Price": func(c *domain.CityPrices, p *domain.Property) {
		c.BuyApartment = append(c.BuyApartment, p)
	},
	"Salaries And Financing": func(c *domain.CityPrices, p *domain.Property) {
		c.Salaries = append(c.Salaries, p)
	},
}

func parseCity(name string, rc io.ReadCloser) *domain.City {
	doc, err := goquery.NewDocumentFromReader(rc)
	if err != nil {
		log.Printf("Unexpected error during read of document: err = %v\n", err)
		return nil
	}

	return &domain.City{
		Name:   name,
		Prices: parsePrices(doc),
	}
}

func parsePrices(d *goquery.Document) *domain.CityPrices {
	rows := d.Find(tableRows)
	if rows.Size() > 0 {
		result := &domain.CityPrices{}
		var category string

		rows.Each(func(i int, s *goquery.Selection) {
			if ctg := s.Find(categorySelector); ctg != nil && ctg.Text() != "" {
				category = ctg.Text()
			}
			if p := parseProperty(s); p != nil {
				if update, ok := priceUpdaters[category]; ok {
					update(result, p)
				} else {
					log.Printf("Can't add property to empty category: category = %s, property = %+v\n", category, p)
				}
			}
		})
		return result
	}
	return nil
}

func getAmount(s string) float64 {
	s = strings.ReplaceAll(s, ",", "")
	if s == "" {
		return 0
	}
	r, err := strconv.ParseFloat(s, 10)
	if err != nil {
		log.Printf("Unexpected error during convert string to float: string = %s, err = %v\n", s, err)
	}
	return r
}

func parseProperty(s *goquery.Selection) *domain.Property {
	if td := s.Find(td).First(); td != nil && td.Text() != "" {
		var p domain.Property
		p.Name = strings.TrimSpace(td.Text())
		p.Avg = parseAvg(s)
		p.Min = parseMin(s)
		p.Max = parseMax(s)
		return &p
	}
	return nil
}

func parseAvg(s *goquery.Selection) float64 {
	if pv := s.Find(classFirstCurrency).First(); pv != nil {
		s := strings.Split(strings.TrimSpace(pv.Text()), "\u00A0")[0]
		return getAmount(s)
	}
	return 0
}

func parseMin(s *goquery.Selection) float64 {
	if pl := s.Find(classBarTextLeft).First(); pl != nil {
		return getAmount(strings.TrimSpace(pl.Text()))
	}
	return 0
}

func parseMax(s *goquery.Selection) float64 {
	if ph := s.Find(classBarTextRight).First(); ph != nil {
		return getAmount(strings.TrimSpace(ph.Text()))
	}
	return 0
}
