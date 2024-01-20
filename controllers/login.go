package controllers

import (
	"net/http"

	"areydra.com/mamani/api/database"
	"areydra.com/mamani/api/models"
	"areydra.com/mamani/api/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type LoginInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"omitempty"`
}

type LoginSuccses struct {
	Token string `json:"token" binding:"required"`
}

type Error struct {
	Code    int    `json:"code" binding:"required"`
	Message string `json:"message" binding:"required"`
}

func Login(c *gin.Context) {
	var input LoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.Users

	if err := database.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"error": Error{4001, "Email atau password salah!"}})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"error": Error{4001, "Email atau password salah!"}})
		return
	}

	tokenString, err := utils.GenerateToken(user)

	if err != nil {
		c.AbortWithStatus(502)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": LoginSuccses{tokenString}})
}
