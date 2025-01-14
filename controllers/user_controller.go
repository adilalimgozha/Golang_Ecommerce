package controllers

import (
	"net/http"

	"github.com/adilalimgozha/Golang_Ecommerce/config"
	"github.com/adilalimgozha/Golang_Ecommerce/models"
	"github.com/gin-gonic/gin"
)

// Get all users
func GetUsers(c *gin.Context) {
	var users []models.User
	result := config.DB.Preload("Role").Find(&users)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to retrieve users"})
		return
	}
	c.JSON(http.StatusOK, users)
}

// Create a new user
func CreateUser(c *gin.Context) {
	var user models.User

	// Получаем данные из тела запроса
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный запрос"})
		return
	}

	// Проверка на уникальность имени пользователя
	var existingUser models.User
	if err := config.DB.Where("username = ?", user.Username).First(&existingUser).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Имя пользователя уже занято"})
		return
	}

	// Проверка на уникальность email
	if err := config.DB.Where("email = ?", user.Email).First(&existingUser).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Этот email уже используется"})
		return
	}

	// Хэшируем пароль перед сохранением
	if err := user.SetPassword(user.PasswordHash); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при хэшировании пароля"})
		return
	}

	// Сохраняем нового пользователя в базу данных
	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при создании пользователя"})
		return
	}

	// Отправляем успешный ответ
	c.JSON(http.StatusOK, gin.H{
		"message": "Пользователь успешно зарегистрирован",
		"user": gin.H{
			"username": user.Username,
			"email":    user.Email,
		},
	})
}

// Get user by ID
func GetUserByID(c *gin.Context) {
	var user models.User
	id := c.Param("id")
	result := config.DB.First(&user, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

// Get all user addresses
func GetUserAddress(c *gin.Context) {
	var user_addresses []models.UserAddress
	result := config.DB.Preload("User").Find(&user_addresses)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to retrieve user_addresses"})
		return
	}
	c.JSON(http.StatusOK, user_addresses)
}

// Create a new user address
func CreateUserAddress(c *gin.Context) {
	var user_addresse models.UserAddress
	if err := c.ShouldBindJSON(&user_addresse); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	result := config.DB.Create(&user_addresse)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to create user_addresses"})
		return
	}
	c.JSON(http.StatusCreated, user_addresse)
}

// Get user address by ID
func GetUserAddressByID(c *gin.Context) {
	var user_addresse models.UserAddress
	id := c.Param("id")
	result := config.DB.First(&user_addresse, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user_addresse not found"})
		return
	}
	c.JSON(http.StatusOK, user_addresse)
}
