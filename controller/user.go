package controller

import (
	"fmt"
	"lantu/dao"
	"lantu/logic"
	"lantu/models"
	"math/rand"
	"net/http"
	"strings"
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
	// 将 邮箱-验证码 存入redis
	dao.RedisSet(registerEmail+"emailCode", registerEmail+code)
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

func Register(ctx *gin.Context) {
	registerEmail := ctx.PostForm("email")
	registerCode := ctx.PostForm("code")
	password := ctx.PostForm("password")

	emailCode := dao.RedisGet(registerEmail + "emailCode")
	fmt.Printf("emailCode: %v\n", emailCode)

	if emailCode != registerEmail+registerCode {
		ctx.JSON(http.StatusOK, gin.H{
			"success": false,
			"msg":     "验证码错误",
		})
		return
	}

	umAccount := strings.Split(registerEmail, "@")
	isSuccess := models.InsertUser(umAccount[0], password)
	if isSuccess {
		ctx.JSON(http.StatusOK, gin.H{
			"success": true,
			"msg":     "注册成功",
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"success": false,
			"msg":     "注册成功",
		})
	}
}
