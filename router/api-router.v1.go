package router

import (
	"github.com/gin-gonic/gin"
	"jyksServer/controller"
)

func setApiRouterMy(router *gin.Engine) {
	// 自己写的路由
	router.GET("/myjson", controller.MyJson)
}
