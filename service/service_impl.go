package service

import (
	"encoding/json"
	"fetch_24/domain"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ReceiptServiceImpl struct {
}

func NewReceiptServiceImpl() *ReceiptServiceImpl {
	return &ReceiptServiceImpl{}
}

var receipts = make(map[string]domain.Receipt)


func (receiptService *ReceiptServiceImpl) GetReceiptPoints(c *gin.Context) {
	id := c.Param("id")

	receipt, exists := receipts[id]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "No receipt found for that id"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"points": receipt.Points})
}

func (receiptService *ReceiptServiceImpl) ProcessReceipt(c *gin.Context) {

	var receipt domain.Receipt

	body, _ := io.ReadAll(c.Request.Body)

	jsonMarshErr := json.Unmarshal(body, &receipt)

	if jsonMarshErr != nil {
		log.Println(jsonMarshErr)
		c.JSON(http.StatusBadRequest, gin.H{"error": "unable to marshall json into a receipt object"})
		return
	}

	receiptId, receiptIdErr := domain.GenerateHash()

	if receiptIdErr != nil {
		log.Println(receiptIdErr)
		c.JSON(http.StatusBadRequest, gin.H{"error": "unable to generate uuid"})
		return
	}

	receipt.CalculatePoints()

	receipt.Id = receiptId

	receipts[receipt.Id] = receipt

	// 	Example Response:
	// ```json
	// { "id": "7fb1377b-b223-49d9-a31a-5a02701dd310" }

	c.JSON(http.StatusOK, gin.H{"id": receipt.Id})
}
