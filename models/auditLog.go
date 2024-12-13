package models

import "time"

type AuditLog struct {
	LogID     uint      `gorm:"primaryKey"`
	Action    string    `gorm:"not null"`
	UserID    uint      `gorm:"not null"`
	User      User      `gorm:"foreignKey:UserID"`
	Timestamp time.Time `gorm:"autoCreateTime"`
}
