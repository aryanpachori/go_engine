package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)



type OrderRequest struct {
	Stock    string  `json:"stock" binding:"required"`
	Type     string  `json:"type" binding:"required,oneof=buy sell"` 
	Price    float64 `json:"price" binding:"required"`
	Quantity int     `json:"quantity" binding:"required"`
}


type OrderResponse struct {
	FilledQty int `json:"filledQty"`
}

func processOrder(c *gin.Context) {
	var order OrderRequest

	
	if err := c.ShouldBindJSON(&order); err != nil {
	
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	
	filledQty := order.Quantity 


	response := OrderResponse{
		FilledQty: filledQty,
	}

	
	c.JSON(http.StatusOK, response)
}

func main() {
	r := gin.Default()

	
	r.POST("/order", processOrder)

	r.Run("127.0.0.1:3000")
}
