package main

import (
	"github.com/gin-gonic/gin"
)

func setRouter() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")
	{
		api.GET("/test", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"message": "test successful",
			})
		})
		api.POST("/chat/message")
	}
	return router
}

func main() {
	router := setRouter()
	router.Run(":8080")
}
