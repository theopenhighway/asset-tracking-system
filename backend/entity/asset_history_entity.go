package entity

import "time"

type AssetHistory struct {
	Id        int
	AssetId   int
	Status    string
	UpdatedBy int
	UpdatedAt time.Time
}
