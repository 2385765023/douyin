package dal

import "gorm.io/gorm"

//	表名users
type User struct {
	gorm.Model
	ID       uint `json:"user_id"`
	Pwd      string
	NickName string `json:"Name"`
}

func QueryUser(userId string) (*User, error) {
	var res User
	if err := DB.First(&res, userId).Error; err != nil {
		return nil, err
	}
	return &res, nil
}
