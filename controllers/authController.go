package controllers

import (
	"net/http"
	"time"

	"github.com/adilalimgozha/Golang_Ecommerce/config"
	"github.com/adilalimgozha/Golang_Ecommerce/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// Secret key for JWT signing and validation
var jwtSecretKey = []byte("your-secret-key") // Use a strong secret key

// Generate JWT token
func generateJWT(user models.User) (string, error) {
	// Define the JWT claims
	claims := jwt.MapClaims{
		"userID":   user.UserID,
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 72).Unix(), // Token expiration (72 hours)
	}

	// Create a new token with the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	return token.SignedString(jwtSecretKey)
}

// User Login
func Login(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Find the user in the database
	var existingUser models.User
	if err := config.DB.Where("username = ?", user.Username).First(&existingUser).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Check if the password matches
	if !existingUser.CheckPassword(user.PasswordHash) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Generate JWT token
	token, err := generateJWT(existingUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	// Return the token to the client
	c.JSON(http.StatusOK, gin.H{"token": token})
}
