package models

type UserAddress struct {
	AddressID uint   `gorm:"primaryKey"`
	UserID    uint   `gorm:"not null"`
	User      User   `gorm:"foreignKey:UserID"`
	Street    string `gorm:"not null"`
	City      string `gorm:"not null"`
	State     string `gorm:"not null"`
	ZipCode   string `gorm:"not null"`
}
