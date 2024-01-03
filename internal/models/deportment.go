package models

import "time"

type Deportment struct {
	ID        int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAT time.Time
}
