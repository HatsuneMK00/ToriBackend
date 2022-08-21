package api

import (
	"ToriBackend/apis"
	"github.com/gin-gonic/gin"
)

type cardRouter struct{}

func (r cardRouter) AddCardRoutes(rg *gin.RouterGroup) {
	card := rg.Group("/card")
	{
		card.GET("/:card_pack_id", apis.ApiGroupApp.Card.GetCards)
		card.GET("/all_card_packs", apis.ApiGroupApp.Card.GetAllCardPacks)
	}
}
