package controllers

import (
	"fmt"

	"areydra.com/mamani/api/database"
	"areydra.com/mamani/api/models"
)

type TotalWalletsInformationInput struct {
	UserId   uint `json:"user_id" binding:"required"`
	IsIncome bool `json:"is_income" binding:"required"`
	Amount   uint `json:"amount" binding:"required"`
}

func CreateTotalWalletsInformation(userId uint, isIncome bool, amount uint) bool {
	var totalWalletsInformation models.TotalWalletsInformation

	income := uint(0)
	outcome := uint(0)

	if isIncome {
		income = amount
	} else {
		outcome = amount
	}

	// Create a new record instead of updating the existing one
	newTotalWalletsInformation := models.TotalWalletsInformation{
		UserId:  userId,
		Income:  totalWalletsInformation.Income + income,
		Outcome: totalWalletsInformation.Outcome + outcome,
	}

	if err := database.DB.Create(&newTotalWalletsInformation).Error; err != nil {
		fmt.Println("Error while creating totalWalletsInformation:", err)
		return false
	}

	return true
}

func UpdateTotalWalletsInformation(userId uint, isIncome bool, amount uint, isDeleteWallet bool) bool {
	var totalWalletsInformation models.TotalWalletsInformation

	if err := database.DB.Where("user_id = ?", userId).First(&totalWalletsInformation).Error; err != nil {
		// If the record doesn't exist, create a new one
		CreateTotalWalletsInformation(userId, isIncome, amount)
		return true
	}

	income := totalWalletsInformation.Income
	outcome := totalWalletsInformation.Outcome

	fmt.Println("isDeleteWallet", isDeleteWallet)
	fmt.Println("amount", amount)

	if isIncome {
		income += amount
	} else if isDeleteWallet {
		income -= amount
	} else {
		outcome += amount
	}

	fmt.Println("income", income)

	// Use UpdateColumn to force updating specific columns
	if err := database.DB.Model(&totalWalletsInformation).UpdateColumns(map[string]interface{}{
		"Income":  income,
		"Outcome": outcome,
	}).Error; err != nil {
		fmt.Println("Error while updating totalWalletsInformation:", err)
		return false
	}

	return true
}

func UpdateWalletBalance(userId uint, id uint, isIncome bool, amount uint) bool {
	var wallet models.Wallets

	if err := database.DB.Where("user_id = ? AND id = ?", userId, id).First(&wallet).Error; err != nil {
		fmt.Println("Error while getting wallet", err)
		return false
	}

	balance := uint(0)

	if isIncome {
		balance = wallet.Balance + amount
	} else {
		balance = wallet.Balance - amount
	}

	if err := database.DB.Model(&wallet).Updates(models.Wallets{
		Balance: balance,
	}).Error; err != nil {
		fmt.Println("Error while updating wallet:", err)
		return false
	}

	return true
}
