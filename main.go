package main

import (
	"github.com/adk-saugat/cofund/db"
	"github.com/adk-saugat/cofund/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main(){
	db.InitDB()
	
	err := godotenv.Load()
	if err != nil{
		panic("couldnot load env")
	}

	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080")
}