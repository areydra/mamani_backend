package utils

import "time"

func GenerateExpiredDateInUnix(days int) int64 {
	now := time.Now()
	return now.Add(time.Duration(days) * 24 * time.Hour).Unix()
}
