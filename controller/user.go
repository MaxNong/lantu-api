package controller

import (
	"fmt"
	"lantu/logic"
	"lantu/models"
	"math/rand"
	"net/http"
	"time"

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

func SendRegisterEmail(ctx *gin.Context) {
	registerEmail := ctx.PostForm("email")

	if registerEmail == "undefined.pingan.com.cn" {
		ctx.JSON(http.StatusOK, gin.H{
			"success": false,
			"msg":     "邮箱地址不能为空",
		})
		fmt.Println("Send fail! - ", "邮箱地址不能为空")
		return
	}

	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	code := fmt.Sprintf("%06v", rnd.Int31n(1000000))
	// 邮件接收方
	mailTo := []string{
		registerEmail,
	}

	subject := "注册蓝兔!"          // 邮件主题
	body := "【蓝兔】本次验证码为" + code // 邮件正文

	emailError := logic.SendMail(mailTo, subject, body, "B端产品协作平台", "")
	if emailError != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"success": false,
			"msg":     emailError,
		})
		fmt.Println("Send fail! - ", emailError)
		return
	} else {
		fmt.Println("Send successfully!")
		ctx.JSON(http.StatusOK, gin.H{
			"success": true,
			"msg":     "发送成功",
		})
	}
}
