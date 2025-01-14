package main

import (
	"log"

	"github.com/adilalimgozha/Golang_Ecommerce/config"
	"github.com/adilalimgozha/Golang_Ecommerce/controllers"
	"github.com/adilalimgozha/Golang_Ecommerce/middleware"
	"github.com/adilalimgozha/Golang_Ecommerce/models"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	config.ConnectDatabase()
	// Initialize the Gin router
	router := gin.Default()

	// CORS middleware
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Content-Type"},
		AllowCredentials: true,
	}))

	err := config.DB.AutoMigrate(&models.Role{})
	if err != nil {
		log.Fatalf("Failed to migrate Role: %v", err)
	}

	err = config.DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalf("Failed to migrate User: %v", err)
	}

	err = config.DB.AutoMigrate(&models.Category{})
	if err != nil {
		log.Fatalf("Failed to migrate Category: %v", err)
	}

	err = config.DB.AutoMigrate(&models.Product{})
	if err != nil {
		log.Fatalf("Failed to migrate Product: %v", err)
	}

	err = config.DB.AutoMigrate(&models.ShoppingCart{})
	if err != nil {
		log.Fatalf("Failed to migrate ShoppingCart: %v", err)
	}

	err = config.DB.AutoMigrate(&models.CartItem{})
	if err != nil {
		log.Fatalf("Failed to migrate CartItem: %v", err)
	}

	err = config.DB.AutoMigrate(&models.Order{})
	if err != nil {
		log.Fatalf("Failed to migrate Order: %v", err)
	}

	err = config.DB.AutoMigrate(&models.OrderItem{})
	if err != nil {
		log.Fatalf("Failed to migrate OrderItem: %v", err)
	}

	err = config.DB.AutoMigrate(&models.Payment{})
	if err != nil {
		log.Fatalf("Failed to migrate Payment: %v", err)
	}

	err = config.DB.AutoMigrate(&models.Review{})
	if err != nil {
		log.Fatalf("Failed to migrate Review: %v", err)
	}

	err = config.DB.AutoMigrate(&models.Session{})
	if err != nil {
		log.Fatalf("Failed to migrate Session: %v", err)
	}

	err = config.DB.AutoMigrate(&models.UserAddress{})
	if err != nil {
		log.Fatalf("Failed to migrate UserAddress: %v", err)
	}

	err = config.DB.AutoMigrate(&models.ProductImage{})
	if err != nil {
		log.Fatalf("Failed to migrate ProductImage: %v", err)
	}

	err = config.DB.AutoMigrate(&models.AuditLog{})
	if err != nil {
		log.Fatalf("Failed to migrate AuditLog: %v", err)
	}

	err = config.DB.AutoMigrate(&models.Cache{})
	if err != nil {
		log.Fatalf("Failed to migrate Cache: %v", err)
	}

	// User routes
	router.GET("/users", controllers.GetUsers)
	router.POST("/register", controllers.CreateUser)
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
	router.POST("/orders/items", controllers.CreateOrderItem)
	router.GET("/orders/items/:id", controllers.GetOrderItemByID)
	router.DELETE("/orders/items/:id", controllers.DeleteOrderItem)

	// Public Routes (No JWT required)
	router.POST("/login", controllers.Login) // Login route

	// Protected Routes (JWT required)
	protected := router.Group("/protected")
	protected.Use(middleware.JWTAuthMiddleware()) // Apply JWT middleware to this group
	{
		protected.GET("/profile", controllers.GetProfile)
	}

	// Run the server
	router.Run(":8080")
}
