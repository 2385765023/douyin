package middleware

import (
	"time"

	"github.com/RaymondCode/simple-demo/service"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

const (
	SecretKey   = "SecretKeySecretKeySecretKeySecretKeySecretKey"
	IdentityKey = "id"
)

func JwtInit() (*jwt.GinJWTMiddleware, error) {
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Key:        []byte(SecretKey),
		Timeout:    time.Hour,
		MaxRefresh: time.Hour,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(int64); ok {
				return jwt.MapClaims{
					IdentityKey: v,
				}
			}
			return jwt.MapClaims{}
		},
		//	主要负责检查用户密码是否正确
		Authenticator: func(c *gin.Context) (interface{}, error) {
			username := c.Query("username")
			password := c.Query("password")

			if len(username) == 0 || len(password) == 0 {
				return "", jwt.ErrMissingLoginValues
			}
			return service.CheckUser(username, password)
		},
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	})
	return authMiddleware, err
}
