package models_db

import "time"

type URL struct {
	ID        uint64 `gorm:"primaryKey"`
	UserID    uint64 `gorm:"not null;index"`
	ShortURL  string `gorm:"not null;uniqueIndex"`
	OrigURL   string `gorm:"type:text;not null"`
	Redirects uint64 `gorm:"default:0"`
	Revoked   bool   `form:"default:false"`
	CreatedAt time.Time
	ExpiresAt *time.Time
}
