package router

import (
	"github.com/gin-gonic/gin"
	"jyksServer/controller"
	"jyksServer/middleware"
)

func setWebRouter(router *gin.Engine) {
	router.Use(middleware.GlobalWebRateLimit())
	// Always available
	router.GET("/", controller.GetIndexPage)
	router.POST("/login", middleware.CriticalRateLimit(), controller.Login)
	router.GET("/logout", controller.Logout)

	/*// Download files
	fileDownloadAuth := router.Group("/")
	fileDownloadAuth.Use(middleware.DownloadRateLimit(), middleware.FileDownloadPermissionCheck())
	{
		fileDownloadAuth.Static("/upload", common.UploadPath)
		fileDownloadAuth.GET("/explorer", controller.GetExplorerPageOrFile)
	}

	imageDownloadAuth := router.Group("/")
	imageDownloadAuth.Use(middleware.DownloadRateLimit(), middleware.ImageDownloadPermissionCheck())
	{
		imageDownloadAuth.Static("/image", common.ImageUploadPath)
	}

	router.GET("/image", controller.GetImagePage)

	router.GET("/video", controller.GetVideoPage)

	basicAuth := router.Group("/")
	basicAuth.Use(middleware.WebAuth())
	{
		basicAuth.GET("/manage", controller.GetManagePage)
	}*/
}
