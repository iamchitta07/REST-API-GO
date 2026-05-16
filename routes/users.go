package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iamchitta07/models"
)

func signup(ctx *gin.Context) {
	var user models.User
	err := ctx.ShouldBindBodyWithJSON(&user)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse requested data."})
		return
	}

	err = user.Save()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update event."})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "User created successfully!"})
}

func login(ctx *gin.Context) {
	var user models.User
	err := ctx.ShouldBindBodyWithJSON(&user)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse requested data."})
		return
	}
	err = user.ValidateCredentials()
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Login Successfully!"})
}
