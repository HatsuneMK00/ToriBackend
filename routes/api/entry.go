package api

import "github.com/gin-gonic/gin"

type RouterGroup struct {
	userRouter
	cardRouter
}

func (r RouterGroup) AddApiRoutes(rg *gin.RouterGroup) {
	apiGroup := rg.Group("/api")
	r.AddUserRoutes(apiGroup)
	r.AddCardRoutes(apiGroup)
}
