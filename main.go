package main

import (
	"os"

	"areydra.com/mamani/api/controllers"
	"areydra.com/mamani/api/database"
	"areydra.com/mamani/api/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	utils.LoadEnv("staging")

	// uncomment line below before build golang app
	// gin.SetMode(gin.ReleaseMode)

	router := gin.Default()
	database.Connect()

	router.POST("/signup", controllers.Signup)
	router.POST("/login", controllers.Login)
	router.POST("/send/otp", controllers.SendOTP)
	router.POST("/verify/phone-number", controllers.VerifyPhoneNumber)
	router.GET("/user", controllers.GetUser)
	router.POST("/transaction", controllers.CreateTransaction)
	router.GET("/transactions", controllers.GetTransactions)
	router.POST("/wallet", controllers.CreateWallet)

	router.Run(os.Getenv("LISTEN_URL"))
}
