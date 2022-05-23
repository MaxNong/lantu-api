package controller

import (
	"lantu/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context) {
	userName := ctx.Query("userName")
	password := ctx.Query("password")

	_, err := models.QueryUserByUsernameAndPassword(userName, password)

	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"success": false,
			"msg":     "用户名或密码错误",
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"success": true,
			"msg":     "登录成功",
		})
	}

}
