package models

import "time"

type User struct {
	ID            int
	RolID         int
	DirectorionID int
	Name          string
	Phone         string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAT     time.Time
}
