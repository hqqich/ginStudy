package router

import (
	"ginStudy/controller"
	"ginStudy/middleware"
	"github.com/gin-gonic/gin"
)

func setApiRouterV1(router *gin.Engine) {
	router.Use(middleware.GlobalAPIRateLimit())

	basicAuth := router.Group("/api")
	basicAuth.Use(middleware.ApiAuth())
	{
		basicAuth.GET("/myJson", controller.MyJson)
	}
	adminAuth := router.Group("/api")
	adminAuth.Use(middleware.ApiAdminAuth())
	{
		adminAuth.POST("/user", controller.CreateUser)
		adminAuth.PUT("/manage_user", controller.ManageUser)
	}

}
