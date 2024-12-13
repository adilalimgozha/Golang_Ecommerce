package main

import (
	"log"

	"github.com/adilalimgozha/Golang_Ecommerce/config"
	"github.com/adilalimgozha/Golang_Ecommerce/controllers"
	"github.com/adilalimgozha/Golang_Ecommerce/models"
	"github.com/gin-gonic/gin"
)

func main() {

	config.ConnectDatabase()
	// Initialize the Gin router
	router := gin.Default()

	// Миграция таблицы Role
	err := config.DB.AutoMigrate(&models.Role{})
	if err != nil {
		log.Fatalf("Failed to migrate Role: %v", err)
	}

	// Миграция таблицы User
	err = config.DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalf("Failed to migrate User: %v", err)
	}

	// Миграция таблицы Category
	err = config.DB.AutoMigrate(&models.Category{})
	if err != nil {
		log.Fatalf("Failed to migrate Category: %v", err)
	}

	// Миграция таблицы Product
	err = config.DB.AutoMigrate(&models.Product{})
	if err != nil {
		log.Fatalf("Failed to migrate Product: %v", err)
	}

	// Миграция таблицы ShoppingCart
	err = config.DB.AutoMigrate(&models.ShoppingCart{})
	if err != nil {
		log.Fatalf("Failed to migrate ShoppingCart: %v", err)
	}

	// Миграция таблицы CartItem
	err = config.DB.AutoMigrate(&models.CartItem{})
	if err != nil {
		log.Fatalf("Failed to migrate CartItem: %v", err)
	}

	// Миграция таблицы Order
	err = config.DB.AutoMigrate(&models.Order{})
	if err != nil {
		log.Fatalf("Failed to migrate Order: %v", err)
	}

	// Миграция таблицы OrderItem
	err = config.DB.AutoMigrate(&models.OrderItem{})
	if err != nil {
		log.Fatalf("Failed to migrate OrderItem: %v", err)
	}

	// Миграция таблицы Payment
	err = config.DB.AutoMigrate(&models.Payment{})
	if err != nil {
		log.Fatalf("Failed to migrate Payment: %v", err)
	}

	// Миграция таблицы Review
	err = config.DB.AutoMigrate(&models.Review{})
	if err != nil {
		log.Fatalf("Failed to migrate Review: %v", err)
	}

	// Миграция таблицы Session
	err = config.DB.AutoMigrate(&models.Session{})
	if err != nil {
		log.Fatalf("Failed to migrate Session: %v", err)
	}

	// Миграция таблицы UserAddress
	err = config.DB.AutoMigrate(&models.UserAddress{})
	if err != nil {
		log.Fatalf("Failed to migrate UserAddress: %v", err)
	}

	// Миграция таблицы ProductImage
	err = config.DB.AutoMigrate(&models.ProductImage{})
	if err != nil {
		log.Fatalf("Failed to migrate ProductImage: %v", err)
	}

	// Миграция таблицы AuditLog
	err = config.DB.AutoMigrate(&models.AuditLog{})
	if err != nil {
		log.Fatalf("Failed to migrate AuditLog: %v", err)
	}

	// Миграция таблицы Cache
	err = config.DB.AutoMigrate(&models.Cache{})
	if err != nil {
		log.Fatalf("Failed to migrate Cache: %v", err)
	}

	// User routes
	router.GET("/users", controllers.GetUsers)
	router.POST("/users", controllers.CreateUser)
	router.GET("/users/:id", controllers.GetUserByID)

	// Product routes
	router.GET("/products", controllers.GetProducts)
	router.POST("/products", controllers.CreateProduct)
	router.GET("/products/:id", controllers.GetProductByID)

	// Product Images routes
	router.GET("/productImages", controllers.GetProductImages)
	router.POST("/productImages", controllers.CreateProductImage)
	router.GET("/productImages/:id", controllers.GetProductImageByID)

	// Run the server
	router.Run(":8080")
}
