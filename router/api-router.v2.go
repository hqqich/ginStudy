package router

import (
	"ginStudy/controller"
	"ginStudy/middleware"
	"github.com/gin-gonic/gin"
)

func setApiRouterV2(router *gin.Engine) {
	// 使用api限流
	router.Use(middleware.GlobalAPIRateLimit())

	tokenAuth := router.Group("/apiv2")
	tokenAuth.Use(middleware.TokenAuth())
	{
		tokenAuth.GET("/myJson", controller.MyJson)
	}

}
