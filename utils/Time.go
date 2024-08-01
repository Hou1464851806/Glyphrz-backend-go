package utils

import (
	"fmt"
	"time"
)

func GetTodayFormatTime() string {
	now := time.Now()
	today := fmt.Sprintf("%02d-%02d-%02d", now.Year(), int(now.Month()), now.Day())
	return today
}
