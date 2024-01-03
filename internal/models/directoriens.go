package models

import "time"

type Directorion struct {
	ID           int
	DeportmentID int
	Name         string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAT    time.Time
}
