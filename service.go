package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func healthCheck(ginContext *gin.Context) { 
	log.Println("inside healthcheck")

	ginContext.IndentedJSON(http.StatusOK, gin.H{"healthcheck": "great sucess"})
} // close health check

func getReceiptPoints(c *gin.Context) {
    id := c.Param("id")

    // Check if receipt exists
    points, exists := receipts[id]
    if !exists {
        c.JSON(http.StatusNotFound, gin.H{"error": "No receipt found for that id"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"points": points})
}

