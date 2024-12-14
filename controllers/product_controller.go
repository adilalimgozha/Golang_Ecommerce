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

// Get product image by ID
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

// Get all product reviews
func GetProductReviews(c *gin.Context) {
	var productReviews []models.Review
	result := config.DB.Preload("Product").Preload("User").Find(&productReviews)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to retrieve product reviews"})
		return
	}
	c.JSON(http.StatusOK, productReviews)
}

// Create a new product image
func CreateProductReview(c *gin.Context) {
	var productReviews models.Review
	if err := c.ShouldBindJSON(&productReviews); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	result := config.DB.Create(&productReviews)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to create product review"})
		return
	}
	c.JSON(http.StatusCreated, productReviews)
}

// Get product review by ID
func GetProductReviewByID(c *gin.Context) {
	var productReviews models.Review
	id := c.Param("id")
	result := config.DB.First(&productReviews, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product Review not found"})
		return
	}
	c.JSON(http.StatusOK, productReviews)
}

// Update product review by ID
func UpdateProductReview(c *gin.Context) {
	var review models.Review
	id := c.Param("id")

	result := config.DB.First(&review, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product Review not found"})
		return
	}

	if err := c.ShouldBindJSON(&review); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	result = config.DB.Save(&review)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update product review"})
		return
	}

	c.JSON(http.StatusOK, review)
}

// Delete product review by ID
func DeleteProductReview(c *gin.Context) {
	var review models.Review
	id := c.Param("id")

	result := config.DB.First(&review, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product Review not found"})
		return
	}

	result = config.DB.Delete(&review)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete product review"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product review deleted successfully"})
}
