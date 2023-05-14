package router

import (
	"github.com/gin-gonic/gin"
	"jyksServer/controller"
	"jyksServer/middleware"
)

func setApiRouterV2(router *gin.Engine) {
	router.Use(middleware.GlobalAPIRateLimit())

	tokenAuth := router.Group("/apiv2")
	tokenAuth.Use(middleware.TokenAuth())
	{
		// 自己写的路由
		tokenAuth.GET("/myJson", controller.MyJson)
	}

}
