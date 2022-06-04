package middleware

import (
	"time"

	"github.com/RaymondCode/simple-demo/common/constants"
	"github.com/RaymondCode/simple-demo/common/status_code"
	"github.com/RaymondCode/simple-demo/controller"
	"github.com/RaymondCode/simple-demo/service"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

const (
	SecretKey = "SecretKeySecretKeySecretKeySecretKeySecretKey"
)

var AuthMiddleware *jwt.GinJWTMiddleware

func JwtInit() (err error) {
	AuthMiddleware, err = jwt.New(&jwt.GinJWTMiddleware{
		Key:        []byte(SecretKey),
		Timeout:    time.Hour,
		MaxRefresh: time.Hour,
		//	Authenticator认证成功以后，往 Payload 中塞东西
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(int64); ok {
				return jwt.MapClaims{
					constants.IdentityKey: v,
				}
			}
			return jwt.MapClaims{}
		},
		//	检查用户名及密码是否正确,登录时候用
		Authenticator: func(c *gin.Context) (interface{}, error) {
			//	以下为用户名(昵称)登录
			username := c.Query("username")
			password := c.Query("password")

			if len(username) == 0 || len(password) == 0 {
				return "", jwt.ErrMissingLoginValues
			}
			userId, err := service.CheckUser(username, password)
			if err != nil {
				controller.SendResp(c, controller.UserLoginResponse{
					Status: status_code.UserNameOrPwdIncorrectErr,
				})
				return 0, nil
			}
			c.Set(constants.IdentityKey, userId)
			return userId, nil //不将 error 向下传，避免库自动发送 401
		},
		//	自定义登录成功以后返回参数
		LoginResponse: func(c *gin.Context, code int, tokenString string, time time.Time) {
			userId, _ := c.Get(constants.IdentityKey)
			if userId == nil {
				return
			}
			controller.SendResp(c, controller.UserLoginResponse{
				Status: status_code.Success,
				UserId: userId.(int64),
				Token:  tokenString,
			})
		},
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	})
	return err
}
