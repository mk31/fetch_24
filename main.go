package main

import (
	"fetch_24/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

var receiptService *service.ReceiptServiceImpl

func main() {

	ginEngine := gin.Default()

	RegisterRoutes(ginEngine)

	ginEngine.Run(":80")

}

func registerServices() {
	receiptService = service.NewReceiptServiceImpl()
}

func RegisterRoutes(ginEngine *gin.Engine) {
	mainEngine := ginEngine.Group("/api/v1")
	RegisterMainRoutes(mainEngine)

	unAuthedEngine := ginEngine.Group("/api/")
	RegisterUnAuthedRoutes(unAuthedEngine)
}

func RegisterUnAuthedRoutes(unAuthedEngine *gin.RouterGroup) {
	unAuthedEngine.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "Healthy")
	})
}

func RegisterMainRoutes(mainEngine *gin.RouterGroup) {
	mainEngine.GET("/receipts/:id/points", receiptService.GetReceiptPoints)
	mainEngine.POST("/receipts/process", receiptService.ProcessReceipt)
}
