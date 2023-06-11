package models

type Product struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Category int     `json:"category"`
	Price    float64 `json:"price"`
	Stock    int     `json:"stock"`
}
