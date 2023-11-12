package models

import "time"

type OneTimePassword struct {
	ID              uint      `json:"id" gorm:"primaryKey"`
	UserId          uint      `json:"user_id"`
	OneTimePassword string    `json:"one_time_password"`
	IsVerified      int       `json:"is_verified" gorm:"default:0"`
	CreatedAt       time.Time `json:"created_at"`
}

func (OneTimePassword) TableName() string {
	return "one_time_password"
}
