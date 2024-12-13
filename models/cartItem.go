package models

type CartItem struct {
	CartItemID uint         `gorm:"primaryKey"`
	CartID     uint         `gorm:"not null"`
	Cart       ShoppingCart `gorm:"foreignKey:CartID"`
	ProductID  uint         `gorm:"not null"`
	Product    Product      `gorm:"foreignKey:ProductID"`
	Quantity   int          `gorm:"not null"`
}
