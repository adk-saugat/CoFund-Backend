package controllers

import (
	"net/http"

	"github.com/adk-saugat/cofund/models"
	"github.com/adk-saugat/cofund/utils"
	"github.com/gin-gonic/gin"
)

func SeeProfile(ctx *gin.Context){
	userId := ctx.GetInt64("userId")
	user, err := models.GetProfileById(userId)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Couldnot show user profile!"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"user": gin.H{
		"fullName": user.FirstName + " " + user.LastName,
		"email": user.Email,
		"id": user.ID,
		"createdAt": user.CreatedAt,
	}})
}

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

	token, err := utils.GenerateToken(user.Email, user.ID)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Couldnot generate token!"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Login successful!", "token": token})
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

	ctx.JSON(http.StatusCreated, gin.H{"message": "User created successfully!", "user" : gin.H{
		"userId": user.ID,
		"fullName": user.FirstName + " " + user.LastName,
		"email": user.Email,
	}})
}