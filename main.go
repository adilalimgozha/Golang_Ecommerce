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

	// User Adress routes
	router.GET("/users/address", controllers.GetUserAddress)
	router.POST("/users/address", controllers.CreateUserAddress)
	router.GET("/users/address/:id", controllers.GetUserAddressByID)

	// Product routes
	router.GET("/products", controllers.GetProducts)
	router.POST("/products", controllers.CreateProduct)
	router.GET("/products/:id", controllers.GetProductByID)

	// Product Images routes
	router.GET("/products/images", controllers.GetProductImages)
	router.POST("/products/images", controllers.CreateProductImage)
	router.GET("/products/images/:id", controllers.GetProductImageByID)

	// Product Reviews routes
	router.GET("/products/reviews", controllers.GetProductReviews)
	router.POST("/products/reviews", controllers.CreateProductReview)
	router.GET("/products/reviews/:id", controllers.GetProductReviewByID)
	router.PUT("/products/reviews/:id", controllers.UpdateProductReview)
	router.DELETE("/products/reviews/:id", controllers.DeleteProductReview)

	// Shopping Cart routes
	router.GET("/shopping_cart", controllers.GetShoppingCarts)
	router.POST("/shopping_cart", controllers.CreateShoppingCart)
	router.GET("/shopping_cart/:id", controllers.GetShoppingCartByID)

	// Cart Items routes
	router.GET("/shopping_cart/items", controllers.GetCartItems)
	router.POST("/shopping_cart/items", controllers.CreateCartItem)
	router.GET("/shopping_cart/items/:id", controllers.GetCartItemByID)
	router.DELETE("/shopping_cart/items/:id", controllers.DeleteCartItem)

	// Orders routes
	router.GET("/orders", controllers.GetOrders)
	router.POST("/orders", controllers.CreateOrder)
	router.GET("/orders/:id", controllers.GetOrderByID)

	// Order Items routes
	router.GET("/orders/items", controllers.GetOrderItems)
	router.POST("/orders/items", controllers.CreateCartItem)
	router.GET("/orders/items/:id", controllers.GetCartItemByID)
	router.DELETE("/orders/items/:id", controllers.DeleteCartItem)

	// Run the server
	router.Run(":8080")
}
