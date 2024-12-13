package models


// Budget struct
type Budget struct {
	ID     int     `json:"id"`
	Amount float64 `json:"amount"`
	Name   string  `json:"name"`
}
