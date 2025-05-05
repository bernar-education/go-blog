package model

import (
	"time"

	"gorm.io/gorm"
)

// Article struct
type Article struct {
	gorm.Model
	Title       string    `gorm:"not null" json:"title"`
	Description string    `gorm:"not null" json:"description"`
	PublishedAt time.Time `gorm:"not null" json:"published_at"`
	IsActive    bool      `gorm:"default=false"`
}
