package service

import "github.com/gin-gonic/gin"

type Service interface {
	GetReceiptPoints(c *gin.Context)
	ProcessReceipt(c *gin.Context)
}
