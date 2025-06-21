package entity

import "time"

type Assets struct {
	Id           int
	Name         string
	AssetType    string
	Status       string
	Detail       string
	PurchaseDate time.Time
	UpdatedAt    time.Time
}
