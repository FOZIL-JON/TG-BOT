package models

import (
	"time"
)

type Rol struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAT time.Time
}
