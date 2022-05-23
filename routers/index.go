package routers

import (
	"lantu/controller"

	"github.com/gin-gonic/gin"
)

func SetupRouters() *gin.Engine {
	r := gin.Default()

	// 设置静态目录
	r.Static("/static", "static")

	// 用户模块
	userGroup := r.Group("/user")
	{
		userGroup.GET("/login", controller.Login)
	}

	return r
}
