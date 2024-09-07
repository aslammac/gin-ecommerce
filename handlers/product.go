package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Product represents a simple product structure
type Product struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

// In-memory product storage (replace with database in production)
var products = []Product{
	{ID: "1", Name: "Sample Product 1", Price: 19.99},
	{ID: "2", Name: "Sample Product 2", Price: 29.99},
}

// GetProducts handles the GET request to retrieve all products
func GetProducts(c *gin.Context) {
	c.JSON(http.StatusOK, products)
}

// CreateProduct handles the POST request to add a new product
func CreateProduct(c *gin.Context) {
	var newProduct Product
	if err := c.ShouldBindJSON(&newProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Generate a new ID (in production, use a proper ID generation method)
	newProduct.ID = fmt.Sprintf("%d", len(products)+1)

	// Add the new product to the slice
	products = append(products, newProduct)

	c.JSON(http.StatusCreated, newProduct)
}
