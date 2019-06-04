package router

import (
	"blog/controllers"
	"github.com/gin-gonic/gin"
)

func Route(router *gin.Engine)  {

	router.GET("/",controllers.IndexView) // 首页视图
	router.GET("/admin",controllers.AdminView) // 后台管理视图

	// apiV1 路由组
	apiV1 := router.Group("/api/v1")
	{
		apiV1.GET("/getAllArticle",controllers.GetArticleByPage) // 分页获取文章
	}

}