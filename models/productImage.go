package models

import "time"

type ProductImage struct {
	ImageID   uint      `gorm:"primaryKey"`
	ProductID uint      `gorm:"not null"`
	Product   Product   `gorm:"foreignKey:ProductID"`
	ImageURL  string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}
