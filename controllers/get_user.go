package controllers

import (
	"net/http"

	"areydra.com/mamani/api/database"
	"areydra.com/mamani/api/models"
	"areydra.com/mamani/api/utils"
	"github.com/gin-gonic/gin"
)

type SuccessGetUser struct {
	ID              uint   `json:"id"`
	Name            string `json:"name"`
	Email           string `json:"email"`
	IsPhoneVerified int    `json:"is_phone_verified"`
}

func GetUser(c *gin.Context) {
	userId, err := utils.VerifyToken(c.Request.Header["Authorization"])

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	var user models.Users

	if err := database.DB.Where("id = ?", userId).First(&user).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"error": LoginError{4001, "Internal error!"}})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 2001,
		"data": SuccessGetUser{user.ID, user.Name, user.Email, user.IsPhoneVerified},
	})
}
