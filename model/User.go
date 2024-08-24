package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        uint           `json:"id" gorm:"primaryKey;autoIncrement"`
	Name      string         `json:"name" gorm:"not null;unique;index"`
	Password  string         `json:"password" gorm:"not null"`
	Signature string         `json:"signature"`
	Avatar    string         `json:"avatar"`
	Role      uint           `json:"role" gorm:"not null;default:1"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"update_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
