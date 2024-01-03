package models

import "time"

type Tasc struct {
	ID          int
	UserID      int
	Name        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAT   time.Time
}
