package model

import (
	"time"
)

type Link struct {
	ID          string
	ShortCode   string
	OriginalURL string
	Clicks      int
	CreatedAt   time.Time
}
