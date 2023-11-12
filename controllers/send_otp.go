package controllers

import (
	"net/http"

	"areydra.com/mamani/api/database"
	"areydra.com/mamani/api/models"
	"areydra.com/mamani/api/utils"
	"github.com/gin-gonic/gin"
)

type SuccessSendOTP struct {
	Code    int    `json:"code" binding:"required"`
	Message string `json:"message" binding:"required"`
}

func SendOTP(c *gin.Context) {
	userId, err := utils.VerifyToken(c.Request.Header["Authorization"])

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	otp := models.OneTimePassword{UserId: userId, OneTimePassword: utils.GenerateOTP(6)}
	database.DB.Create(&otp)

	c.JSON(http.StatusOK, gin.H{"data": SuccessSendOTP{2001, "OTP berhasil dikirim!"}})
}
