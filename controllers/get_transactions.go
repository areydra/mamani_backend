package controllers

import (
	"net/http"

	"time"

	"areydra.com/mamani/api/database"
	"areydra.com/mamani/api/utils"
	"github.com/gin-gonic/gin"
)

type Transaction struct {
	ID         uint   `json:"id"`
	Note       string `json:"note"`
	Type       uint8  `json:"type"`
	Amount     uint   `json:"amount"`
	UserID     uint   `json:"user_id"`
	DateTime   int64  `json:"date_time"`
	WalletID   uint   `json:"wallet_id"`
	CreatedAt  string `json:"created_at"`
	CategoryID uint   `json:"category_id"`
}

type TransactionData struct {
	Date         string        `json:"date"`
	TotalOutcome uint          `json:"total_outcome"`
	TotalIncome  uint          `json:"total_income"`
	Transactions []Transaction `json:"transactions"`
}

func GetTransactions(c *gin.Context) {
	userId, err := utils.VerifyToken(c.Request.Header["Authorization"])

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	var transactions []Transaction
	database.DB.Table("transactions").
		Select("id, note, type, amount, user_id, date_time, wallet_id, created_at, category_id").
		Where("user_id = ?", userId).
		Order("date_time desc").
		Scan(&transactions)

	mappedData := make(map[string]*TransactionData)

	for _, transaction := range transactions {
		date := time.Unix(transaction.DateTime, 0)
		date = date.UTC()
		dateString := date.Format("2006-01-02")

		// Check if date already exists in mappedData
		if existingData, ok := mappedData[dateString]; ok {
			// Update existing entry in mappedData
			existingData.Transactions = append(existingData.Transactions, transaction)
			if transaction.Type == 1 {
				existingData.TotalOutcome += transaction.Amount
			} else {
				existingData.TotalIncome += transaction.Amount
			}
		} else {
			// If date not found, create a new entry in mappedData
			newEntry := &TransactionData{
				Date:         dateString,
				Transactions: []Transaction{transaction},
			}

			if transaction.Type == 1 {
				newEntry.TotalOutcome += transaction.Amount
			} else {
				newEntry.TotalIncome += transaction.Amount
			}

			mappedData[dateString] = newEntry
		}
	}

	// Convert map values to a slice for JSON serialization
	var result []TransactionData
	for _, entry := range mappedData {
		result = append(result, *entry)
	}

	finalResponse := gin.H{
		"code": 2005,
		"data": result,
	}

	c.JSON(http.StatusOK, finalResponse)
}
