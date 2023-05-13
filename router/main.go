package router

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"jyksServer/common"
	"jyksServer/controller"
)

// SetRouter 绑定路由
func SetRouter(router *gin.Engine) {
	store := cookie.NewStore([]byte(common.SessionSecret))
	// 使用Cookies
	router.Use(sessions.Sessions("jyksServer", store))

	setWebRouter(router)
	setApiRouterV1(router)
	router.NoRoute(controller.Get404Page)
}
