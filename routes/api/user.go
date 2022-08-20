package api

import (
	"ToriBackend/apis"
	"github.com/gin-gonic/gin"
)

type userRouter struct{}

func (r userRouter) AddUserRoutes(rg *gin.RouterGroup) {
	user := rg.Group("/user")
	{
		user.GET("/:id", apis.ApiGroupApp.User.FindUser)
		user.GET("/offset/:offset", apis.ApiGroupApp.User.FindUsersWithOffset)
		user.POST("/", apis.ApiGroupApp.User.AddUser)
		user.DELETE("/:id", apis.ApiGroupApp.User.DeleteUser)
	}
}
