package models_http

import "time"

type UrlRequest struct {
	URL string `json:"url" binding:"required"`
}

type UrlResponse struct {
	ID        uint64     `json:"url"`
	ShortURL  string     `json:"shorturl"`
	OrigURL   string     `json:"origurl"`
	Redirects uint64     `json:"redirects"`
	Revoked   bool       `json:"revoked"`
	CreatedAt time.Time  `json:"createdat"`
	ExpiresAt *time.Time `json:"expiresat"`
}
