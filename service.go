package main

import (
	"encoding/json"
	"io"
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

func processReceipt(c *gin.Context) {

	var receipt Receipt

	body, _ := io.ReadAll(c.Request.Body)

	jsonMarshErr := json.Unmarshal(body, &receipt)

	if jsonMarshErr != nil {
		log.Println(jsonMarshErr)
		c.JSON(http.StatusBadRequest, gin.H{"error": "unable to marshall json into a receipt object"})
		return
	}

	receiptId, receiptIdErr := GenerateHash()

	if receiptIdErr != nil {
		log.Println(receiptIdErr)
		c.JSON(http.StatusBadRequest, gin.H{"error": "unable to generate uuid"})
		return
	}


	receipt.Id = receiptId

	receipts[receipt.Id] = receipt


// 	Example Response:
// ```json
// { "id": "7fb1377b-b223-49d9-a31a-5a02701dd310" }

	c.JSON(http.StatusOK, gin.H{"id": receipt.Id})
}
