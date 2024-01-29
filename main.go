package main

import (
	// "encoding/json"
	"github.com/gin-gonic/gin"
)

func main() {

	ginEngine := gin.Default()

	RegisterRoutes(ginEngine)

	ginEngine.Run(":80")

}


