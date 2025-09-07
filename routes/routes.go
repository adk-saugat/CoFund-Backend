package routes

import (
	"net/http"

	"github.com/adk-saugat/cofund/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine){
	// Dummy server tester
	server.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "Hello, World!"})
	})

	server.POST("/register", controllers.RegisterUser)
}