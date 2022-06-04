package service

import (
	"crypto/md5"
	"fmt"
	"io"

	"github.com/RaymondCode/simple-demo/common/status_code"
	"github.com/RaymondCode/simple-demo/dal"
)

//	通过用户名和密码进行用户认证
func CheckUser(nickName, password string) (uid int64, err error) {
	h := md5.New()
	if _, err := io.WriteString(h, password); err != nil {
		return 0, err
	}
	pwdHashed := fmt.Sprintf("%x", h.Sum(nil))
	user, err := dal.QueryUserByName(nickName)
	if err != nil {
		return 0, err
	}

	if pwdHashed != user.Pwd {
		return 0, status_code.UserNameOrPwdIncorrectErr
	}
	return user.ID, nil
}

func GetUserInfoById(userId int64) (userInfo *dal.User, err error) {
	return dal.QueryUserById(userId)
}
