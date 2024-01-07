package controllers

import (
	"net/http"

	"areydra.com/mamani/api/database"
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

	var transactions []map[string]interface{}
	database.DB.Table("transactions").
		Select("DATE(created_at) as date, SUM(amount) as total_amount, JSONB_AGG(ROW_TO_JSON(transactions.*)) as transactions").
		Where("user_id = ?", userId).
		Group("date").
		Order("date").
		Scan(&transactions)

	c.JSON(http.StatusOK, gin.H{
		"code": 2003,
		"data": transactions,
	})
}
