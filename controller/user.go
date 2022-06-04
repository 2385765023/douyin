package controller

import (
	"log"
	"net/http"
	"sync/atomic"

	"github.com/RaymondCode/simple-demo/common/constants"
	"github.com/RaymondCode/simple-demo/common/status_code"
	"github.com/RaymondCode/simple-demo/service"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

// usersLoginInfo use map to store user info, and key is username+password for demo
// user data will be cleared every time the server starts
// test data: username=zhanglei, password=douyin
var usersLoginInfo = map[string]User{
	"zhangleidouyin": {
		Id:            1,
		Name:          "zhanglei",
		FollowCount:   10,
		FollowerCount: 5,
		IsFollow:      true,
	},
}

var userIdSequence = int64(1)

type UserLoginResponse struct {
	status_code.Status
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token,omitempty"`
}

type UserResponse struct {
	status_code.Status
	User User `json:"user"`
}

func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	token := username + password

	if _, exist := usersLoginInfo[token]; exist {
		c.JSON(http.StatusOK, UserLoginResponse{
			Status: status_code.Status{StatusCode: 1, StatusMsg: "User already exist"},
		})
	} else {
		atomic.AddInt64(&userIdSequence, 1)
		newUser := User{
			Id:   userIdSequence,
			Name: username,
		}
		usersLoginInfo[token] = newUser
		c.JSON(http.StatusOK, UserLoginResponse{
			Status: status_code.Status{StatusCode: 0},
			UserId: userIdSequence,
			Token:  username + password,
		})
	}
}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	token := username + password

	if user, exist := usersLoginInfo[token]; exist {
		c.JSON(http.StatusOK, UserLoginResponse{
			Status: status_code.Status{StatusCode: 0},
			UserId: user.Id,
			Token:  token,
		})
	} else {
		c.JSON(http.StatusOK, UserLoginResponse{
			Status: status_code.Status{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	}
}

func UserInfo(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	userId := int64(claims[constants.IdentityKey].(float64)) //坑:这个是 float64
	log.Println("用户id:", userId)

	if user, err := service.GetUserInfoById(userId); err == nil {
		SendResp(c, UserResponse{
			Status: status_code.Success,
			User: User{
				Id:            user.ID,
				Name:          user.NickName,
				FollowCount:   10,
				FollowerCount: 5,
				IsFollow:      true,
			},
		})
	} else {
		SendResp(c, UserResponse{
			Status: status_code.UserNotExistErr,
		})
	}
}
