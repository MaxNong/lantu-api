package main

import (
	"fmt"
	"lantu/dao"
	"lantu/middleware"
	"lantu/routers"
	"lantu/setting"

	"github.com/gin-gonic/gin"
)

func main() {
	// 加载配置文件
	var conf setting.Conf
	conf.GetConf()

	// 初始化数据库
	err := dao.InitMySQL(&conf.MySQLConfig)
	if err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		return
	} else {
		fmt.Println(" init mysql success")
	}
	defer dao.Close() // 程序退出关闭数据库连接

	// 初始化redis
	dao.InitRedis()

	// 初始化gin engine
	r := gin.Default()

	// 初始化中间件
	middleware.InitMiddleware(r)

	// 初始化gin Engine和路由
	routers.SetupRouters(r)

	// 启动服务
	r.Run(":9090")
}
