package controllers

import (
	"net/http"

	"github.com/adk-saugat/cofund/models"
	"github.com/gin-gonic/gin"
)

func LoginUser(ctx *gin.Context){
	var user models.User

	err := ctx.BindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Couldnot parse the request data!"})
		return
	}

	err = user.ValidateCredentials()
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Couldnot authorize!"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Login successful!"})
}

func RegisterUser(ctx *gin.Context){
	var user models.User

	err := ctx.BindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Couldnot parse the request data!"})
		return
	}

	err = user.Save()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Couldnot save user!"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "User created successfully!", "user" : user})
}