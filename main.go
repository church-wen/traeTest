package main

import (
	"fmt"

	"code.byted.org/zhuchaowen/trae/auth"

	// 由于无法导入该包，推测可能需要使用本地替代包或者重新配置模块路径
	// 这里暂时注释掉该导入，你需要确认正确的包路径或者创建相应的模块
	// "code.byted.org/zhuchaowen/trae/config"
	"code.byted.org/zhuchaowen/trae/config"
	"github.com/gin-gonic/gin"
)

func init() {
	var err error
	if err = config.InitLogger(); err != nil {
		panic(err)
	}
	_, err = config.InitRedis()
	_, err = config.InitMySQL()
	if err != nil {
		panic(err)
	}
}

func main() {
	r := gin.Default()

	r.POST("/login", auth.Login)
	r.GET("/protected", auth.AuthMiddleware(), auth.Protected)

	fmt.Println("服务器启动，监听端口 :8080")
	if err := r.Run(":8080"); err != nil {
		fmt.Println("服务器启动失败:", err)
	}
}
