package domain

type City struct {
	Name       string      `json:"name"`
	Properties []*Property `json:"properties"`
}

type Price struct {
	Avg string `json:"avg"`
	Min string `json:"min"`
	Max string `json:"max"`
}

type Property struct {
	Name  string `json:"name"`
	Price `json:"price"`
}
