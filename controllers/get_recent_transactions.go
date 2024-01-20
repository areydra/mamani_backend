package controllers

import (
	"net/http"

	"areydra.com/mamani/api/database"
	"areydra.com/mamani/api/utils"
	"github.com/gin-gonic/gin"
)

func GetRecentTransactions(c *gin.Context) {
	userId, err := utils.VerifyToken(c.Request.Header["Authorization"])

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	var transactions []Transaction
	database.DB.Table("transactions").
		Select("id, note, type, amount, user_id, date_time, wallet_id, created_at, category_id").
		Where("user_id = ?", userId).
		Limit(5).
		Order("created_at desc").
		Scan(&transactions)

	finalResponse := gin.H{
		"code": 2005,
		"data": transactions,
	}

	c.JSON(http.StatusOK, finalResponse)
}
