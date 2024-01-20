package controllers

import (
	"net/http"

	"areydra.com/mamani/api/database"
	"areydra.com/mamani/api/models"
	"areydra.com/mamani/api/utils"
	"github.com/gin-gonic/gin"
)

type SuccessVerifyPhoneNumber struct {
	Code    int    `json:"code" binding:"required"`
	Message string `json:"message" binding:"required"`
}

type VerifyPhoneNumberInput struct {
	OTP string `json:"otp" binding:"required"`
}

func VerifyPhoneNumber(c *gin.Context) {
	var input VerifyPhoneNumberInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userId, err := utils.VerifyToken(c.Request.Header["Authorization"])

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	var dataOtp models.OneTimePassword

	if err := database.DB.Where("user_id = ?", userId).First(&dataOtp).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"error": Error{4001, "Kode OTP tidak valid!"}})
		return
	}

	if dataOtp.OneTimePassword != input.OTP || dataOtp.IsVerified == 1 {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"error": Error{4001, "Kode OTP tidak valid!"}})
		return
	}

	database.DB.Model(&models.OneTimePassword{}).Where("user_id = ?", userId).Update("is_verified", 1)
	database.DB.Model(&models.Users{}).Where("id = ?", userId).Update("is_phone_verified", 1)
	c.JSON(http.StatusOK, gin.H{"data": SuccessVerifyPhoneNumber{2001, "Verifikasi berhasil!"}})
}
