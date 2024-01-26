package main

import (
	// "encoding/json"
	"github.com/gin-gonic/gin"


)
func main() {
	// golang_library_mpk.LoadEnv(".env.local")

	// jwtSec := os.Getenv("JWT_SEC")

	// jwtSecByteGbl = []byte(jwtSec)

	// envGbl = os.Getenv("ENVY")

	// setConnUrl()

	// setDbConnection()

	// testPingDb()

	ginEngine := gin.Default()

	RegisterRoutes(ginEngine)
	

	/* ---------------------------  Public routes  --------------------------- */

	// public := ginEngine.Group("/v1")

	// public.Use(authenticate())

	//region gets

	// public.GET("/emailToken/:userEmail", getEmailToken)
	
	// //endregion gets

	// //region patches

	// public.PATCH("/patchEmailTokenStandardContactsNextSyncToken", patchEmailTokenStandardContactsNextSyncToken)

	// //endregion patches

	// //region posts

	// public.POST("/emailToken", postEmailToken)
	// //endregion posts

	// /* ---------------------------  Private routes  --------------------------- */

	// private := ginEngine.Group("/api/")

	// private.GET("health", healthCheck)
	
	// private.POST("login", login)

	ginEngine.Run()

} // close main