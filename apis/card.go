package apis

import (
	"ToriBackend/entity/response"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type CardApi interface {
	GetAllCardPacks(c *gin.Context)
	GetCards(c *gin.Context)
}

type cardApi struct{}

func (api cardApi) GetAllCardPacks(c *gin.Context) {
	cardPacks, ok := cardService.GetAllCardPacks()
	if ok {
		c.JSON(http.StatusOK, response.Response{
			Code: http.StatusOK,
			Data: cardPacks,
		})
	} else {
		c.JSON(http.StatusOK, response.Response{
			Code:    http.StatusInternalServerError,
			Message: "failed to get card packs",
		})
	}
}

func (api cardApi) GetCards(c *gin.Context) {
	id := c.Param("card_pack_id")
	cardPackId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			Code:    http.StatusBadRequest,
			Message: "card pack id must be type int",
		})
		return
	}
	cards, ok := cardService.GetCards(uint(cardPackId))
	if ok {
		c.JSON(http.StatusOK, response.Response{
			Code: http.StatusOK,
			Data: cards,
		})
	} else {
		c.JSON(http.StatusOK, response.Response{
			Code:    http.StatusInternalServerError,
			Message: "failed to get cards",
		})
	}
}
