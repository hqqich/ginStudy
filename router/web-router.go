package router

import (
	"ginStudy/controller"
	"ginStudy/middleware"
	"github.com/gin-gonic/gin"
)

// 设置页面路由
func setWebRouter(router *gin.Engine) {

	router.Use(middleware.GlobalWebRateLimit())
	// Always available
	router.GET("/", controller.GetIndexPage)
	router.POST("/login", middleware.CriticalRateLimit(), controller.Login)
	router.POST("/loginByToken", middleware.CriticalRateLimit(), controller.LoginByToken)
	router.GET("/logout", controller.Logout)

}
