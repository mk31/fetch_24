package service

import "github.com/gin-gonic/gin"

type ReceiptService interface {
	GetReceiptPoints(c *gin.Context)
	ProcessReceipt(c *gin.Context)
}
