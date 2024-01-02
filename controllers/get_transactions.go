package controllers

import (
	"net/http"

	"areydra.com/mamani/api/database"
	"areydra.com/mamani/api/models"
	"areydra.com/mamani/api/utils"
	"github.com/gin-gonic/gin"
)

type SuccessGetTransactions struct {
	ID              uint   `json:"id"`
	Name            string `json:"name"`
	Email           string `json:"email"`
	IsPhoneVerified int    `json:"is_phone_verified"`
}

func GetTransactions(c *gin.Context) {
	userId, err := utils.VerifyToken(c.Request.Header["Authorization"])

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	var transactions []models.Transactions

	if err := database.DB.Where("user_id = ?", userId).Find(&transactions).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"error": LoginError{4001, "Internal error!"}})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 2003,
		"data": transactions,
	})
}
