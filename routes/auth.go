package routes

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

type AuthRouter struct{}

func (r AuthRouter) AddAuthRoutes(rg *gin.RouterGroup, authJWT *jwt.GinJWTMiddleware) {
	rg.POST("/login", authJWT.LoginHandler)
	rg.GET("/refresh_token", authJWT.RefreshHandler)
}
