package service

import "github.com/gin-gonic/gin"

type Service interface {
	getReceiptPoints(c *gin.Context)
	processReceipt(c *gin.Context)
}
