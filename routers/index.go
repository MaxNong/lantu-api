package routers

import (
	"lantu/controller"

	"github.com/gin-gonic/gin"
)

func SetupRouters(r *gin.Engine) *gin.Engine {
	// 设置静态目录
	r.Static("/static", "static")

	// 用户模块
	userGroup := r.Group("/user")
	{
		userGroup.GET("/login", controller.Login)
		userGroup.POST("/sendRegisterEmail", controller.SendRegisterEmail)
		userGroup.POST("/register", controller.Register)
	}

	return r
}
