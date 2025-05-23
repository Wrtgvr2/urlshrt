package models_db

import (
	"time"

	"github.com/google/uuid"
)

type RefreshToken struct {
	JTI       uuid.UUID `gorm:"nut null;type:uuid;primaryKey;index"`
	UserID    uint64    `gorm:"not null"`
	ExpiresAt time.Time `gorm:"not null;index"`
	Revoked   bool      `gorm:"default:false"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
