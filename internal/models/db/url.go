package models_db

import "time"

type URL struct {
	ID        uint   `gorm:"primaryKey"`
	ShortURL  string `gorm:"not null;uniqueIndex"`
	UserID    uint64 `gorm:"not null;index"`
	OrigURL   string `gorm:"type:text;not null"`
	CreatedAt time.Time
	ExpiresAt *time.Time
}
