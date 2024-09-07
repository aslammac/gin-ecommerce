package handlers

import (
	"fmt"
	"gin-ecommerce/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SendOTP(c *gin.Context) {
	fmt.Println("Received POST request to /auth/send-otp")
    
    var request struct {
        Phone string `json:"phone"`
    }
    
    if err := c.ShouldBindJSON(&request); err != nil {
        fmt.Println("Error binding JSON:", err)
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    
    fmt.Println("Phone number received:", request.Phone)
    
    // Your OTP sending logic here

	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "OTP sent successfully"})
}

func VerifyOTP(c *gin.Context) {
	var input struct {
		Phone string `json:"phone"`
		OTP   string `json:"otp"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Verify OTP (implement this in utils/otp.go)
	if utils.VerifyOTP(input.Phone, input.OTP) {
		// Generate JWT token
		token, err := utils.GenerateJWT(input.Phone)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status": "success",
			"message": "OTP verified successfully",
			"token": token,
		})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid OTP"})
	}

}

func Signup(c *gin.Context) {
	var input struct {
		Name  string `json:"name" binding:"required"`
		Email string `json:"email" binding:"required,email"`
		Phone string `json:"phone" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Here you would typically:
	// 1. Check if the user already exists
	// 2. Hash the password (if you're using one)
	// 3. Store the user in the database

	// For this example, we'll just return a success message
	// In a real application, you'd want to handle potential errors

	c.JSON(http.StatusCreated, gin.H{
		"status":  "success",
		"message": "User registered successfully",
		"user": gin.H{
			"name":  input.Name,
			"email": input.Email,
			"phone": input.Phone,
		},
	})
}