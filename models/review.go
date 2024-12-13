package models

import "time"

type Review struct {
	ReviewID  uint      `gorm:"primaryKey"`
	ProductID uint      `gorm:"not null"`
	Product   Product   `gorm:"foreignKey:ProductID"`
	UserID    uint      `gorm:"not null"`
	User      User      `gorm:"foreignKey:UserID"`
	Rating    int       `gorm:"not null"`
	Comment   string    `gorm:"type:text"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}
