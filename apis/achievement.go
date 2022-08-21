package apis

import (
	"ToriBackend/entity/response"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type AchievementApi interface {
	GetAllAchievements(c *gin.Context)
	GetUnlockedAchievements(c *gin.Context)
	CanUnlockNewAchievements(c *gin.Context)
}

type achievementApi struct {
}

func (a achievementApi) GetAllAchievements(c *gin.Context) {
	achievements, ok := achievementService.GetAllAchievements()
	if ok {
		c.JSON(http.StatusOK, response.Response{
			Code: http.StatusOK,
			Data: achievements,
		})
	} else {
		c.JSON(http.StatusOK, response.Response{
			Code:    http.StatusInternalServerError,
			Message: "fail to get all achievements",
		})
	}
}

func (a achievementApi) GetUnlockedAchievements(c *gin.Context) {
	userId := c.Query("user_id")
	id, err := strconv.Atoi(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			Code:    http.StatusBadRequest,
			Message: "user id needs to be uint",
		})
		return
	}
	if achievements, ok := achievementService.GetUnlockedAchievements(uint(id)); ok {
		c.JSON(http.StatusOK, response.Response{
			Code: http.StatusOK,
			Data: achievements,
		})
	} else {
		c.JSON(http.StatusOK, response.Response{
			Code:    http.StatusInternalServerError,
			Message: "fail to get achievements",
		})
	}
}

func (a achievementApi) CanUnlockNewAchievements(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}
