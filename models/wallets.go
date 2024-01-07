package models

import "time"

type Wallets struct {
	ID           uint      `json:"id" gorm:"primaryKey"`
	UserId       uint      `json:"user_id"`
	Name         string    `json:"name"`
	Location     string    `json:"location"`
	Emoji        string    `json:"emoji"`
	Balance      uint      `json:"balance"`
	MonthlyLimit uint      `json:"monthly_limit"`
	CreatedAt    time.Time `json:"created_at"`
}
