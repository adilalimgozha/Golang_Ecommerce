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
	result := config.DB.Preload("User").Find(&ordersItems)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to retrieve ordersItems"})
		return
	}
	c.JSON(http.StatusOK, ordersItems)
}
