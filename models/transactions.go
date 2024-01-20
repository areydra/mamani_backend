package models

import "time"

type Transactions struct {
	ID         uint      `json:"id" gorm:"primaryKey"`
	WalletId   uint      `json:"wallet_id"`
	UserId     uint      `json:"user_id"`
	Amount     uint      `json:"amount"`
	CategoryId uint      `json:"category_id"`
	Type       uint8     `json:"type"`
	Note       string    `json:"note"`
	DateTime   int64     `json:"date_time"`
	CreatedAt  time.Time `json:"created_at"`
}
