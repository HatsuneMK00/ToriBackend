package services

import (
	"ToriBackend/entity"
	"ToriBackend/global"
)

type CardService interface {
	GetAllCardPacks() ([]entity.CardPack, bool)
	GetCards(cardPackId uint) ([]entity.Card, bool)
}

type cardService struct{}

func (s cardService) GetAllCardPacks() ([]entity.CardPack, bool) {
	var cardPacks []entity.CardPack
	result := global.MysqlDB.Find(&cardPacks)
	ok := true
	if result.Error != nil {
		global.Logger.Errorf("%v", result.Error)
		ok = false
	}
	return cardPacks, ok
}

func (s cardService) GetCards(cardPackId uint) ([]entity.Card, bool) {
	var cards []entity.Card
	var cardPack entity.CardPack
	cardPack.ID = cardPackId
	ok := true
	err := global.MysqlDB.Model(&cardPack).Association("Cards").Find(&cards)
	if err != nil {
		global.Logger.Errorf("%v", err)
		ok = false
	}
	return cards, ok
}
