package entity

import (
	"database/sql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username      string         `json:"username" binding:"required"`
	Password      string         `json:"password" binding:"required"`
	Email         string         `json:"email"`
	StudentNumber sql.NullString `gorm:"unique;index"`
	Records       []Record       `json:"records"`
}

type Record struct {
	gorm.Model
	UserID      uint
	CardID      uint
	Card        Card
	Description string
}
