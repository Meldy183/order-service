package models

type Order struct {
	ID       ID
	Item     string
	Quantity int
}
type ID = string
