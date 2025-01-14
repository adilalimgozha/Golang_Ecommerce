package controllers

import (
	"net/http"

	"github.com/adilalimgozha/Golang_Ecommerce/config"
	"github.com/adilalimgozha/Golang_Ecommerce/models"
	"github.com/gin-gonic/gin"
)

// Get all Orders
func GetOrders(c *gin.Context) {
	var orders []models.Order
	result := config.DB.Preload("User").Find(&orders)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to retrieve orders"})
		return
	}
	c.JSON(http.StatusOK, orders)
}

// Create a new Order
func CreateOrder(c *gin.Context) {
	var order models.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	result := config.DB.Create(&order)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to create Order"})
		return
	}
	c.JSON(http.StatusCreated, order)
}

// Get Order by ID
func GetOrderByID(c *gin.Context) {
	var order models.Order
	id := c.Param("id")
	result := config.DB.Preload("User").First(&order, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Shopping Cart not found"})
		return
	}
	c.JSON(http.StatusOK, order)
}

// Get all OrderItems
func GetOrderItems(c *gin.Context) {
	var ordersItems []models.OrderItem
	result := config.DB.Preload("Order").Preload("Product").Find(&ordersItems)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to retrieve ordersItems"})
		return
	}
	c.JSON(http.StatusOK, ordersItems)
}

// Create OrderItem
func CreateOrderItem(c *gin.Context) {
	var orderItem models.OrderItem
	var cartItem models.CartItem

	if err := c.ShouldBindJSON(&orderItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	tx := config.DB.Begin()

	if err := tx.Where("product_id = ?", orderItem.ProductID).First(&cartItem).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusNotFound, gin.H{"error": "Cart item not found"})
		return
	}

	var product models.Product
	if err := tx.Where("product_id = ?", cartItem.ProductID).First(&product).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	if product.Stock < cartItem.Quantity {
		tx.Rollback()
		c.JSON(http.StatusBadRequest, gin.H{"error": "Not enough stock for this product"})
		return
	}

	product.Stock -= cartItem.Quantity
	if err := tx.Save(&product).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update product stock"})
		return
	}

	orderItem.Price = product.Price
	orderItem.Quantity = cartItem.Quantity
	orderItem.ProductID = cartItem.ProductID

	if err := tx.Create(&orderItem).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create Order Item"})
		return
	}

	tx.Commit()

	c.JSON(http.StatusCreated, orderItem)
}

func GetOrderItemByID(c *gin.Context) {
	var orderItem models.OrderItem
	id := c.Param("id")

	if err := config.DB.Preload("Product").First(&orderItem, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order Item not found"})
		return
	}

	c.JSON(http.StatusOK, orderItem)
}

// Detele OrderItem
func DeleteOrderItem(c *gin.Context) {
	var orderItem models.OrderItem
	id := c.Param("id")

	if err := config.DB.First(&orderItem, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order Item not found"})
		return
	}

	var product models.Product
	if err := config.DB.Where("product_id = ?", orderItem.ProductID).First(&product).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	product.Stock += orderItem.Quantity
	if err := config.DB.Save(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update product stock"})
		return
	}

	if err := config.DB.Delete(&orderItem).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete order item"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order item deleted and product stock updated successfully"})
}
