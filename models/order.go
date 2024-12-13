package models

import "time"

type Order struct {
	OrderID     uint      `gorm:"primaryKey"`
	UserID      uint      `gorm:"not null"`
	User        User      `gorm:"foreignKey:UserID"`
	OrderDate   time.Time `gorm:"autoCreateTime"`
	Status      string    `gorm:"not null"`
	TotalAmount float64   `gorm:"not null"`
}
