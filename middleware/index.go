package middleware

import "github.com/gin-gonic/gin"

func InitMiddleware(ctx *gin.Engine) {
	// 设置跨域
	ctx.Use(CorsMiddleware)
}
