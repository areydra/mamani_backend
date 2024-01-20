package controllers

import (
	"net/http"

	"areydra.com/mamani/api/database"
	"areydra.com/mamani/api/models"
	"areydra.com/mamani/api/utils"
	"github.com/gin-gonic/gin"
)

type DeleteWalletInput struct {
	ID uint `json:"id" binding:"required"`
}

func DeleteWallet(c *gin.Context) {
	userId, errVerifyToken := utils.VerifyToken(c.Request.Header["Authorization"])

	if errVerifyToken != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": errVerifyToken.Error()})
		return
	}

	var input DeleteWalletInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	wallet := models.Wallets{
		UserId: userId,
		ID:     input.ID,
	}

	database.DB.Delete(&wallet)

	c.JSON(http.StatusOK, gin.H{
		"code":    2006,
		"message": "Berhasil hapus wallet!",
	})
}
