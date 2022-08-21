package api

import "github.com/gin-gonic/gin"

type RouterGroup struct {
	userRouter
	cardRouter
	recordRouter
	courseRouter
	achievementRouter
}

func (r RouterGroup) AddApiRoutes(rg *gin.RouterGroup) {
	apiGroup := rg.Group("/api")
	r.AddUserRoutes(apiGroup)
	r.AddCardRoutes(apiGroup)
	r.AddRecordRoutes(apiGroup)
	r.AddCourseRoutes(apiGroup)
	r.AddAchievementRoutes(apiGroup)
}
