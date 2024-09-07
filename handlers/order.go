package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Order represents a simple order structure
type Order struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	Products  []Product `json:"products"`
	Total     float64   `json:"total"`
	CreatedAt time.Time `json:"created_at"`
}

// In-memory order storage (replace with database in production)
var orders = []Order{}

// GetOrders handles the GET request to retrieve all orders for a user
func GetOrders(c *gin.Context) {
	userID, exists := c.Get("phone_number")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	// Filter orders for the current user
	userOrders := []Order{}
	for _, order := range orders {
		if order.UserID == userID.(string) {
			userOrders = append(userOrders, order)
		}
	}

	c.JSON(http.StatusOK, userOrders)
}

// CreateOrder handles the POST request to create a new order
func CreateOrder(c *gin.Context) {
	userID, exists := c.Get("phone_number")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	var newOrder Order
	if err := c.ShouldBindJSON(&newOrder); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Set order details
	newOrder.ID = fmt.Sprintf("%d", len(orders)+1)
	newOrder.UserID = userID.(string)
	newOrder.CreatedAt = time.Now()

	// Calculate total (in a real app, you'd verify product prices from the database)
	var total float64
	for _, product := range newOrder.Products {
		total += product.Price
	}
	newOrder.Total = total

	// Add the new order to the slice
	orders = append(orders, newOrder)

	c.JSON(http.StatusCreated, newOrder)
}