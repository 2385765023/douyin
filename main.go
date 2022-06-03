package main

import (
	"github.com/RaymondCode/simple-demo/dal"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	dal.Init() //	初始化数据库连接

	initRouter(r)

	r.Run("0.0.0.0:8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
