package models

import "time"

type User struct {
	ID            int
	Name          string
	DirectorionID int
	Phone         string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAT     time.Time
}
