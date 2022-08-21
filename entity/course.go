package entity

import (
	"gorm.io/gorm"
	"time"
)

type Course struct {
	gorm.Model
	Name       string  `json:"name" binding:"required"`
	Duration   int     `json:"duration"`
	Level      string  `json:"level"`
	Content    string  `json:"content"`
	ImageCover string  `json:"image_cover"`
	Users      []*User `json:"users" gorm:"many2many:user_courses;"`
}

type UserCourse struct {
	UserId    uint `gorm:"primaryKey"`
	CourseId  uint `gorm:"primaryKey"`
	CreatedAt time.Time
	DeletedAt gorm.DeletedAt
}
