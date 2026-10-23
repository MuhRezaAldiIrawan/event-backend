package controllers

import (
	"event-backend/config"
	"event-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type AuthInputRegister struct {
	Name     string `json:"Name" binding:"required"`
	Email    string `json:"Email" binding:"required,email"`
	Password string `json:"Password" binding:"required"`
}

type AuthInputLogin struct {
	Email    string `json:"Email" binding:"required,email"`
	Password string `json:"Password" binding:"required"`
}

func RegisterUser(context *gin.Context) {
	var input AuthInputRegister
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input", "error": err.Error()})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to hash password", "error": err.Error()})
		return
	}

	user := models.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: string(hashedPassword),
	}

	userCreated := config.DB.Create(&user)
	if userCreated.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to create user", "error": userCreated.Error.Error()})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "User created successfully", "user": gin.H{"Name": user.Name, "Email": user.Email}})
}
