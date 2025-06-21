package entity

import "time"

type AssetAssignment struct {
	Id         int
	AssetId    int
	UserId     int
	AssignedAt time.Time
	ReturnedAt time.Time
	AssignedBy int
}
