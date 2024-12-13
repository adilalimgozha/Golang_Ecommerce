package models

import "time"

type Cache struct {
	CacheKey       string    `gorm:"primaryKey"`
	CacheValue     string    `gorm:"not null"`
	ExpirationTime time.Time `gorm:"not null"`
}
