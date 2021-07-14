package database

import (
	"time"
)

type Url struct {
	ID        uint
	Url       string
	ExpiredAt time.Time
}

var UrlColumns = []string{"id", "url", "expired_at"}