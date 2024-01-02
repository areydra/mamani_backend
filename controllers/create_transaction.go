package controllers

import (
	"net/http"

	"areydra.com/mamani/api/database"
	"areydra.com/mamani/api/models"
	"areydra.com/mamani/api/utils"
	"github.com/gin-gonic/gin"
)

type TransactionInput struct {
	UserId     uint   `json:"user_id" binding:"required"`
	WalletId   uint   `json:"wallet_id" binding:"required"`
	Amount     uint   `json:"amount" binding:"required"`
	CategoryId uint   `json:"category_id" binding:"required"`
	Note       string `json:"note" binding:"required"`
	DateTime   int64  `json:"date_time" binding:"required"`
}

type TransactionSuccess struct {
	ID         uint   `json:"id"`
	WalletId   uint   `json:"wallet_id"`
	Amount     uint   `json:"amount"`
	CategoryId uint   `json:"category_id"`
	Note       string `json:"note"`
	DateTime   int64  `json:"date_time"`
}

func CreateTransaction(c *gin.Context) {
	userId, errVerifyToken := utils.VerifyToken(c.Request.Header["Authorization"])

	if errVerifyToken != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": errVerifyToken.Error()})
		return
	}

	var input TransactionInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	transaction := models.Transactions{
		UserId:     userId,
		WalletId:   input.WalletId,
		Amount:     input.Amount,
		CategoryId: input.CategoryId,
		Note:       input.Note,
		DateTime:   input.DateTime,
	}

	database.DB.Create(&transaction)

	c.JSON(http.StatusOK, gin.H{
		"code": 2002,
		"data": TransactionSuccess{
			transaction.ID,
			transaction.WalletId,
			transaction.Amount,
			transaction.CategoryId,
			transaction.Note,
			transaction.DateTime,
		},
		"message": "Berhasil membuat transaksi!",
	})
}
