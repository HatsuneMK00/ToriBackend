package entity

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Username string    `json:"username" binding:"required"`
	Password string    `json:"password" binding:"required"`
	Email    string    `json:"email"`
	Birthday time.Time `json:"birthday"`
	Records  []Record  `json:"records"`
}

type Record struct {
	gorm.Model
	UserID      uint
	CardID      uint
	Card        Card
	Description string
}
