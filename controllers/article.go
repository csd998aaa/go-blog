package controllers

import (
	. "blog/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetArticleByPage(ctx *gin.Context)  {
	page := ctx.Query("page")

	if page == "" {
		page = "0"
	}
	fmt.Println(page)
	user := &User{UserName:"admin"}
	DB.Find(user).Scan(user)
	ctx.JSON(http.StatusOK,&user)
}
