package entity

import "time"

type User struct {
	Id        int
	Name      string
	Email     string
	PwHash    string
	Role      string
	DeptId    int
	IsActive  bool
	CreatedAt time.Time
	UpdatedAt time.Time
}
