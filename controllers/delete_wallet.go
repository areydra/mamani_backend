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

	var walletInformation models.Wallets

	// Check if the wallet exists before attempting to delete
	if err := database.DB.Where("user_id = ? AND id = ?", userId, input.ID).First(&walletInformation).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Wallet not found"})
		return
	}

	// Save the wallet balance before deletion
	balance := walletInformation.InitialBalance

	// Delete the wallet
	if err := database.DB.Delete(&walletInformation).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update total wallet information
	isSuccessUpdateTotalWallet := UpdateTotalWalletsInformation(userId, false, balance, true)

	if !isSuccessUpdateTotalWallet {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Failed to update wallet information"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    2006,
		"message": "Berhasil hapus wallet!",
	})
}
