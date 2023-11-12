package models

import "time"

type Users struct {
	ID              uint      `json:"id" gorm:"primaryKey"`
	Name            string    `json:"name"`
	Email           string    `json:"email"`
	Password        string    `json:"password"`
	IsPhoneVerified int       `json:"is_phone_verified" gorm:"default:0"`
	CreatedAt       time.Time `json:"created_at"`
}
