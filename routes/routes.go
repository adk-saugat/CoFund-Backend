package routes

import (
	"net/http"

	"github.com/adk-saugat/cofund/controllers"
	"github.com/adk-saugat/cofund/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine){
	// Dummy server tester
	server.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "Hello, World!"})
	})

	server.POST("/auth/register", controllers.RegisterUser)
	server.POST("/auth/login", controllers.LoginUser)
	
	authenticated :=  server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.GET("/auth/me", controllers.SeeProfile)
}