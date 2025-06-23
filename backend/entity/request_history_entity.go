package entity

import "time"

type RequestHistory struct {
	Id             int
	RequestId      int
	ApprovalStatus string
	RequestStatus  string
	UpdatedBy      int
	UpdatedAt      time.Time
}
