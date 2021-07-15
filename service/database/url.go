package database

import (
	"time"
)

type Url struct {
	ID        uint
	Url       string `gorm:"type:varchar(8183);not null"`
	ExpiredAt time.Time
}

var UrlColumns = []string{"id", "url", "expired_at"}