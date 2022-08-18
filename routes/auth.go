package routes

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthRouter struct{}

func (r AuthRouter) AddAuthRoutes(rg *gin.RouterGroup, authJWT *jwt.GinJWTMiddleware) {
	rg.POST("/login", authJWT.LoginHandler)
	rg.GET("/refresh_token", authJWT.RefreshHandler)
	rg.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "hello jwt")
	})
}
