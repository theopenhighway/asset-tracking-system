package entity

import "time"

type Request struct {
	Id             int
	UserId         int
	AssetType      string
	Reason         string
	ApproverUserId int
	CreatedAt      time.Time
	RequestStatus  string
	UpdatedAt      time.Time
}
