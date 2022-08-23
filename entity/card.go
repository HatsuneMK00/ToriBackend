package entity

import "gorm.io/gorm"

type Card struct {
	gorm.Model
	CardPackId  uint   `json:"card_pack_id"`
	Description string `json:"description"`
	ImageCover  string `json:"image_cover"`
}

type CardPack struct {
	gorm.Model
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	ImageCover  string `json:"image_cover"`
	Cards       []Card `json:"cards"`
}
