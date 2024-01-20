package models

import "time"

type TotalWalletsInformation struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	UserId    uint      `json:"user_id"`
	Income    uint      `json:"income"`
	Outcome   uint      `json:"outcome"`
	CreatedAt time.Time `json:"created_at"`
}

func (TotalWalletsInformation) TableName() string {
	return "total_wallets_information"
}
