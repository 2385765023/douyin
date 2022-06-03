package dal

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//	全局 db 连接
var DB *gorm.DB

func Init() (err error) {
	DB, err = gorm.Open(mysql.Open(MySQLDSN))
	if err != nil {
		log.Fatal("初始化数据库失败")
		panic(err)
	}
	log.Println("初始化数据库成功")
	DB.AutoMigrate(&User{})
	return
}
