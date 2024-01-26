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
	// private := ginEngine.Group("/api/")

	unAuthedEngine.GET("health", healthCheck)
}

func RegisterMainRoutes(mainEngine *gin.RouterGroup){
	//region gets
	// routerGroup.GET("/emailToken/:userEmail", getEmailToken)		
	//endregion gets

	//region patches

	// routerGroup.PATCH("/patchEmailTokenStandardContactsNextSyncToken", patchEmailTokenStandardContactsNextSyncToken)

	//endregion patches

	//region posts

	// routerGroup.POST("/emailToken", postEmailToken)
	//endregion posts

}

/* ---------------------------  Private routes  --------------------------- */

// private := ginEngine.Group("/api/")

// private.GET("health", healthCheck)