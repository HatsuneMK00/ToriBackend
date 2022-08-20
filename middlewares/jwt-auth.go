package middlewares

import (
	"ToriBackend/entity"
	"ToriBackend/entity/request"
	"ToriBackend/global"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"time"
)

func InitJWTAuth() *jwt.GinJWTMiddleware {
	middleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "GoApp",
		Key:         []byte("this is a secret key"),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: "id",
		SendCookie:  false,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*entity.User); ok {
				return jwt.MapClaims{
					jwt.IdentityKey: v.ID,
					"username":      v.Username,
					"birthday":      v.Birthday,
					"email":         v.Email,
				}
			}
			return jwt.MapClaims{}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var login request.Login
			if err := c.ShouldBindJSON(&login); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			userName := login.Username
			password := login.Password

			if userName == "admin" && password == "123" {
				user := entity.User{
					Model:    gorm.Model{ID: 123},
					Username: "admin",
					Password: "123",
					Email:    "email@email.com",
					Birthday: time.Date(1999, time.January, 9, 0, 0, 0, 0, time.FixedZone("Asia/Shanghai", 8*60*60)),
				}
				return &user, nil
			}

			return nil, jwt.ErrFailedAuthentication
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": "you are not authorized",
			})
		},
		TokenLookup:       "header: Authorization, query: token, cookie: jwt",
		TokenHeadName:     "Bearer",
		TimeFunc:          time.Now,
		SendAuthorization: false,
	})
	if err != nil {
		global.Logger.Errorf("JWT Error:" + err.Error())
	}
	return middleware
}
