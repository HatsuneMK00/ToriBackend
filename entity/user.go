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
}
