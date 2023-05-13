package router

import (
	"github.com/gin-gonic/gin"
	"jyksServer/controller"
	"jyksServer/middleware"
)

func setApiRouterV1(router *gin.Engine) {
	router.Use(middleware.GlobalAPIRateLimit())

	//router.POST("/api/file", middleware.FileUploadPermissionCheck(), controller.UploadFile)
	//router.POST("/api/image", middleware.ImageUploadPermissionCheck(), controller.UploadImage)

	basicAuth := router.Group("/api")
	basicAuth.Use(middleware.ApiAuth())
	{
		//basicAuth.DELETE("/file", controller.DeleteFile)
		//basicAuth.DELETE("/image", controller.DeleteImage)
		//basicAuth.PUT("/user", controller.UpdateSelf)
		// 自己写的路由
		basicAuth.GET("/myJson", controller.MyJson)
	}
	//adminAuth := router.Group("/api")
	//adminAuth.Use(middleware.ApiAdminAuth())
	//{
	//	adminAuth.POST("/user", controller.CreateUser)
	//	adminAuth.PUT("/manage_user", controller.ManageUser)
	//	adminAuth.GET("/option", controller.GetOptions)
	//	adminAuth.PUT("/option", controller.UpdateOption)
	//}

}
