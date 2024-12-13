package models

import "time"

type Payment struct {
	PaymentID     uint      `gorm:"primaryKey"`
	OrderID       uint      `gorm:"not null"`
	Order         Order     `gorm:"foreignKey:OrderID"`
	Amount        float64   `gorm:"not null"`
	PaymentDate   time.Time `gorm:"autoCreateTime"`
	PaymentMethod string    `gorm:"not null"`
}
