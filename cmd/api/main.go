package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	var router *gin.Engine = gin.Default()
	
	router.SetTrustedProxies(nil)

	router.GET("/", func( ctx *gin.Context ) {
		ctx.JSON(200, gin.H{
			"message": "API server is running",
			"status": "SUCCESS !",
		})
	} )

	router.Run(":8080")
}