package api

import (
	"ToriBackend/apis"
	"github.com/gin-gonic/gin"
)

type achievementRouter struct {
}

func (r achievementRouter) AddAchievementRoutes(rg *gin.RouterGroup) {
	achievement := rg.Group("/achievement")
	{
		achievement.GET("/all", apis.ApiGroupApp.Achievement.GetAllAchievements)
		achievement.GET("/unlocked", apis.ApiGroupApp.Achievement.GetUnlockedAchievements)
		achievement.GET("/can_unlock", apis.ApiGroupApp.Achievement.CanUnlockNewAchievements)
	}
}
