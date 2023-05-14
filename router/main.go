package router

import (
	"ginStudy/common"
	"ginStudy/controller"
	"ginStudy/middleware"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

// SetRouter 绑定路由
func SetRouter(router *gin.Engine) {
	store := cookie.NewStore([]byte(common.SessionSecret))
	// 使用Cookies
	router.Use(sessions.Sessions("ginStudy", store))
	// 使用后端跨域
	router.Use(middleware.CORSMiddleware())

	setWebRouter(router)
	setApiRouterV1(router)
	setApiRouterV2(router)
	router.NoRoute(controller.Get404Page)
}
