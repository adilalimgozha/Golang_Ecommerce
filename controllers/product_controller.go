package controllers

import (
	"net/http"

	"github.com/adilalimgozha/Golang_Ecommerce/config"
	"github.com/adilalimgozha/Golang_Ecommerce/models"
	"github.com/gin-gonic/gin"
)

// Get all products
func GetProducts(c *gin.Context) {
	var products []models.Product
	result := config.DB.Preload("Category").Find(&products)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to retrieve products"})
		return
	}
	c.JSON(http.StatusOK, products)
}

// Create a new product
func CreateProduct(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	result := config.DB.Create(&product)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to create product"})
		return
	}
	c.JSON(http.StatusCreated, product)
}

// Get product by ID
func GetProductByID(c *gin.Context) {
	var product models.Product
	id := c.Param("id")
	result := config.DB.First(&product, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}
	c.JSON(http.StatusOK, product)
}

// Get all product images
func GetProductImages(c *gin.Context) {
	var productImages []models.ProductImage
	result := config.DB.Preload("Product").Find(&productImages)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to retrieve product Images"})
		return
	}
	c.JSON(http.StatusOK, productImages)
}

// Create a new product image
func CreateProductImage(c *gin.Context) {
	var productImages models.ProductImage
	if err := c.ShouldBindJSON(&productImages); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	result := config.DB.Create(&productImages)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to create product Image"})
		return
	}
	c.JSON(http.StatusCreated, productImages)
}

// Get product by ID
func GetProductImageByID(c *gin.Context) {
	var productImages models.ProductImage
	id := c.Param("id")
	result := config.DB.First(&productImages, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product Image not found"})
		return
	}
	c.JSON(http.StatusOK, productImages)
}
