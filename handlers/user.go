package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// User represents a simple user profile structure
type User struct {
	PhoneNumber string `json:"phone_number"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Address     string `json:"address"`
}

// In-memory user storage (replace with database in production)
var users = map[string]User{}

// GetUserProfile handles the GET request to retrieve the user's profile
func GetUserProfile(c *gin.Context) {
	phoneNumber, exists := c.Get("phone_number")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	user, exists := users[phoneNumber.(string)]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "User profile not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// UpdateUserProfile handles the PUT request to update the user's profile
func UpdateUserProfile(c *gin.Context) {
	phoneNumber, exists := c.Get("phone_number")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	var updatedUser User
	if err := c.ShouldBindJSON(&updatedUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Ensure the phone number matches the authenticated user
	updatedUser.PhoneNumber = phoneNumber.(string)

	// Update the user in the map
	users[phoneNumber.(string)] = updatedUser

	c.JSON(http.StatusOK, updatedUser)
}