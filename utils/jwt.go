package utils

import (
	"errors"
	"os"
	"time"

	"areydra.com/mamani/api/models"
	"github.com/golang-jwt/jwt/v5"
)

type MyCustomClaims struct {
	Expired int64 `json:"expired"`
	Id      uint  `json:"id"`
	jwt.RegisteredClaims
}

func VerifyToken(authorization []string) (uint, error) {
	secretKey := []byte(os.Getenv("JWT_SECRET_KEY"))

	if len(authorization) == 0 {
		return 0, errors.New("Invalid token.")
	}

	tokenString := authorization[0][len("Bearer "):]

	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(t *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil || !token.Valid {
		return 0, errors.New("Invalid token.")
	}

	claims, ok := token.Claims.(*MyCustomClaims)

	if ok && claims.Expired < time.Now().Unix() {
		return 0, errors.New("Token has expired.")
	}

	return claims.Id, nil
}

func GenerateToken(user models.Users) (string, error) {
	secretKey := []byte(os.Getenv("JWT_SECRET_KEY"))

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":      user.ID,
		"email":   user.Email,
		"expired": GenerateExpiredDateInUnix(30),
	})

	tokenString, err := token.SignedString(secretKey)

	if err != nil {
		return tokenString, errors.New("Error generate token!")
	}

	return tokenString, nil
}
