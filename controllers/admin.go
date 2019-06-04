package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func AdminView(ctx *gin.Context)  {
	ctx.HTML(http.StatusOK,"admin/index.html",nil)
}
