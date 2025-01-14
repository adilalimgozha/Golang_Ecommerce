package controllers

import (
	"net/http"

	"github.com/adilalimgozha/Golang_Ecommerce/config"
	"github.com/adilalimgozha/Golang_Ecommerce/models"
	"github.com/gin-gonic/gin"
)

// Get all ShoppingCarts
func GetShoppingCarts(c *gin.Context) {
	var shoppingCarts []models.ShoppingCart
	result := config.DB.Preload("User").Find(&shoppingCarts)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to retrieve Shopping Cart"})
		return
	}
	c.JSON(http.StatusOK, shoppingCarts)
}

// Create a new ShoppingCart
func CreateShoppingCart(c *gin.Context) {
	var shoppingCart models.ShoppingCart
	if err := c.ShouldBindJSON(&shoppingCart); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	result := config.DB.Create(&shoppingCart)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to create Shopping Cart"})
		return
	}
	c.JSON(http.StatusCreated, shoppingCart)
}

// Get ShoppingCart by ID
func GetShoppingCartByID(c *gin.Context) {
	var shoppingCart models.ShoppingCart
	id := c.Param("id")
	result := config.DB.Preload("User").First(&shoppingCart, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Shopping Cart not found"})
		return
	}
	c.JSON(http.StatusOK, shoppingCart)
}

// GetShoppingCart - Get all items in the user's shopping cart with product details
func GetCartItems(c *gin.Context) {
	var cartItems []models.CartItem
	var userCartID = 1
	// Preload the Product data along with CartItem data
	if err := config.DB.Preload("Product").Where("cart_id = ?", userCartID).Find(&cartItems).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch shopping cart"})
		return
	}

	// Return the cart items along with product details
	c.JSON(http.StatusOK, cartItems)
}

// CreateCartItem with descreasing amount of stock
func CreateCartItem(c *gin.Context) {
	var cartItem models.CartItem
	if err := c.ShouldBindJSON(&cartItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	result := config.DB.Create(&cartItem)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to create Order"})
		return
	}

	c.JSON(http.StatusCreated, cartItem)
}

// Get ShoppingCart by ID
func GetCartItemByID(c *gin.Context) {
	var cartItem models.CartItem
	id := c.Param("id")
	result := config.DB.Preload("Cart").Preload("Product").First(&cartItem, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Cart Item not found"})
		return
	}
	c.JSON(http.StatusOK, cartItem)
}

// Delete product cartItem by ID
func DeleteCartItem(c *gin.Context) {
	var cartItem models.CartItem
	id := c.Param("id")

	result := config.DB.First(&cartItem, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product cartItem not found"})
		return
	}

	result = config.DB.Delete(&cartItem)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete product cartItem"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product cartItem deleted successfully"})
}
