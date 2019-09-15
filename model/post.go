package model

import (
	"time"
)

type Post struct {
	ID        int    `gorm:"primary_key"`
	Title     string `json:"username"`
	CreatedAt time.Time
	UpdatedAT time.Time
	DeletedAt time.Time
}
