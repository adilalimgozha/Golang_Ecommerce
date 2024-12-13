package models

type OrderItem struct {
	OrderItemID uint    `gorm:"primaryKey"`
	OrderID     uint    `gorm:"not null"`
	Order       Order   `gorm:"foreignKey:OrderID"`
	ProductID   uint    `gorm:"not null"`
	Product     Product `gorm:"foreignKey:ProductID"`
	Quantity    int     `gorm:"not null"`
	Price       float64 `gorm:"not null"`
}
