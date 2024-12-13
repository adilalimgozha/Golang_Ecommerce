package models

import "time"

type Product struct {
	ProductID   uint      `gorm:"primaryKey"`
	Name        string    `gorm:"not null"`
	Description string    `gorm:"type:text"`
	Price       float64   `gorm:"not null"`
	Stock       int       `gorm:"not null"`
	CategoryID  uint      `gorm:"not null"`
	Category    Category  `gorm:"foreignKey:CategoryID"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
}
