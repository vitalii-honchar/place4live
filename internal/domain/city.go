package domain

type City struct {
	Name       string      `json:"name"`
	Properties []*Property `json:"properties"`
}

type Price struct {
	Avg float64 `json:"avg"`
	Min float64 `json:"min"`
	Max float64 `json:"max"`
}

type Property struct {
	Name  string `json:"name"`
	Price `json:"price"`
}
