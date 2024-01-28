package main

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(ginEngine *gin.Engine){
	mainEngine := ginEngine.Group("/api/v1")
	RegisterMainRoutes(mainEngine)
	
	unAuthedEngine := ginEngine.Group("/api/")
	RegisterUnAuthedRoutes(unAuthedEngine)
}

func RegisterUnAuthedRoutes(unAuthedEngine *gin.RouterGroup) {
	unAuthedEngine.GET("health", healthCheck)
}

func RegisterMainRoutes(mainEngine *gin.RouterGroup){
	mainEngine.GET("/receipts/:id/points", getReceiptPoints)
	mainEngine.POST("/receipts/process", processReceipt)

}