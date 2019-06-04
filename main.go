package main

import (
	"blog/conf"
	"blog/models"
	"blog/router"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	app := gin.Default()

	// 设置视图
	app.LoadHTMLGlob("./views/**/*")

	// 注册路由
	router.Route(app)

	// 初始化数据库
	db,err := models.InitDB()
	if err != nil {
		log.Fatal("ERR:数据库连接出错！REASON：",err)
	}
	defer db.Close()

	// 启动服务
	app.Run(conf.GetConfig().Server.Port)
}
