package dal

import (
	"errors"

	"github.com/RaymondCode/simple-demo/common/status_code"
	"gorm.io/gorm"
)

//	表名users
type User struct {
	gorm.Model
	//	用户唯一标识
	ID  int64  
	Pwd string
	//	用户昵称，因为要做用户名登录，故也是唯一的
	NickName string  
	Phone    string
	Mail     string
	//	关注数
	FollowCount   int64  
	//	粉丝数
	FollowerCount int64  
}

func QueryUserById(userId int64) (*User, error) {
	var res User
	if err := DB.First(&res, userId).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return &res, nil
}
 
//通过用户昵称检索用户信息
func QueryUserByName(nickName string) (*User, error) { 
	var res User
	if err := DB.Where("nick_name = ?", nickName).First(&res).Error;err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil,status_code.UserNotExistErr
		}
		return nil, err
	}
	return &res, nil
}
