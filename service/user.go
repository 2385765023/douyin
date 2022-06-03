package service

import (
	"crypto/md5"
	"errors"
	"fmt"
	"io"

	"github.com/RaymondCode/simple-demo/dal"
)

func CheckUser(userName, password string) (uid int64, err error) {
	h := md5.New()
	if _, err := io.WriteString(h, password); err != nil {
		return 0, err
	}
	pwdHashed := fmt.Sprintf("%x", h.Sum(nil))
	user, err := dal.QueryUser(userName)
	if err != nil {
		return 0, nil
	}
	if pwdHashed != user.Pwd {
		return 0, errors.New("密码不正确") //todo:	未来须统一错误码
	}
	return int64(user.ID), nil
}

func GetUserInfo(userId string) (userInfo *dal.User, err error) {
	return dal.QueryUser(userId)
}
