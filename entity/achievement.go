package entity

import (
	"gorm.io/gorm"
	"time"
)

type Achievement struct {
	gorm.Model
	Title       string  `json:"title" binding:"required"`
	Icon        string  `json:"icon"`
	Requirement string  `json:"requirement"`
	Users       []*User `json:"users" gorm:"many2many:user_achievements;"`
}

type UserAchievement struct {
	UserId        uint `gorm:"primaryKey"`
	AchievementId uint `gorm:"primaryKey"`
	CreatedAt     time.Time
	DeletedAt     gorm.DeletedAt
}
