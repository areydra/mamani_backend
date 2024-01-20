package controllers

import (
	"net/http"

	"areydra.com/mamani/api/database"
	"areydra.com/mamani/api/models"
	"areydra.com/mamani/api/utils"
	"github.com/gin-gonic/gin"
)

type WalletInput struct {
	Name         string `json:"name" binding:"required"`
	Location     string `json:"location" binding:"required"`
	Emoji        string `json:"emoji" binding:"required"`
	Balance      uint   `json:"balance" binding:"required"`
	MonthlyLimit uint   `json:"monthly_limit" binding:"required"`
}

type WalletSuccess struct {
	ID             uint   `json:"id"`
	Name           string `json:"name"`
	Location       string `json:"location"`
	Emoji          string `json:"emoji"`
	Balance        uint   `json:"balance"`
	InitialBalance uint   `json:"initial_balance"`
	MonthlyLimit   uint   `json:"monthly_limit"`
}

func CreateWallet(c *gin.Context) {
	userId, errVerifyToken := utils.VerifyToken(c.Request.Header["Authorization"])

	if errVerifyToken != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": errVerifyToken.Error()})
		return
	}

	var input WalletInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	wallet := models.Wallets{
		UserId:         userId,
		Name:           input.Name,
		Location:       input.Location,
		Emoji:          input.Emoji,
		Balance:        input.Balance,
		InitialBalance: input.Balance,
		MonthlyLimit:   input.MonthlyLimit,
	}

	err := database.DB.Create(&wallet).Error

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	isSuccessUpdateTotalWallet := UpdateTotalWalletsInformation(userId, true, input.Balance, false)

	if !isSuccessUpdateTotalWallet {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Failed to update wallet information"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 2004,
		"data": WalletSuccess{
			wallet.ID,
			wallet.Name,
			wallet.Location,
			wallet.Emoji,
			wallet.Balance,
			wallet.MonthlyLimit,
			wallet.InitialBalance,
		},
		"message": "Berhasil membuat wallet!",
	})
}
