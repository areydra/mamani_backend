package controllers

import (
	"log"
	"net/http"

	"areydra.com/mamani/api/database"
	"areydra.com/mamani/api/models"
	"areydra.com/mamani/api/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type SignupInput struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"Email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type SignupSuccess struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func Signup(c *gin.Context) {
	var input SignupInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPasswordBytes, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)

	if err != nil {
		log.Fatalln(err)
	}

	user := models.Users{Name: input.Name, Email: input.Email, Password: string(hashedPasswordBytes)}
	database.DB.Create(&user)

	tokenString, err := utils.GenerateToken(user)

	if err != nil {
		c.AbortWithStatus(502)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    2001,
		"data":    SignupSuccess{user.ID, user.Name, user.Email},
		"message": "Registrasi berhasil!",
		"token":   tokenString,
	})
}
