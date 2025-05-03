package models_http

import "time"

type UrlRequest struct {
	URL string `json:"url" binding:"required"`
}

// UserID uint64 `json:"userid" binding:"required"`
type UrlResponse struct {
	ID        uint
	UserID    uint64
	ShortURL  string
	OrigURL   string
	Redirects uint64
	Revoked   bool
	CreatedAt time.Time
	ExpiresAt *time.Time
}
