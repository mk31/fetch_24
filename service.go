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