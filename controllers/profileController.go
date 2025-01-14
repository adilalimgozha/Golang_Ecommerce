package controllers

import (
	"net/http"

	"github.com/adilalimgozha/Golang_Ecommerce/config"
	"github.com/adilalimgozha/Golang_Ecommerce/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// GetProfile handles the request to the protected route.
func GetProfile(c *gin.Context) {
	// Extract the JWT token from the Authorization header
	tokenString := c.GetHeader("Authorization")[7:] // Remove the "Bearer " prefix

	// Parse and validate the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Ensure the token's signing method is valid
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.NewValidationError("Unexpected signing method", jwt.ValidationErrorSignatureInvalid)
		}
		// Return the secret key used to sign the token
		return []byte("your-secret-key"), nil // Use the same secret key as when generating the JWT
	})

	// Handle errors with token parsing/validation
	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	// Extract claims from the token (username or userID)
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Could not parse token claims"})
		return
	}

	// Get the username or userID from the claims (make sure to use the correct claim name)
	username := claims["username"].(string)

	// Fetch the user's profile data from the database
	var user models.User
	if err := config.DB.Where("username = ?", username).Preload("Role").First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Respond with the user's profile data
	c.JSON(http.StatusOK, gin.H{
		"userID":    user.UserID,
		"username":  user.Username,
		"email":     user.Email,
		"createdAt": user.CreatedAt,
		"role":      user.Role.RoleName,
	})
}
