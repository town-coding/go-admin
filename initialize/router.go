package initialize

import (
	"github.com/gin-gonic/gin"
	"go-admin/global"
	"go-admin/router"
	"net/http"
)

// Routers 初始化路由
func Routers() *gin.Engine {
	Router := gin.New()
	// gin.Recovery()：Gin 自带的中间件，用于从 panic 中恢复，防止程序崩溃，并返回 HTTP 500 错误
	Router.Use(gin.Recovery())
	if gin.Mode() == gin.DebugMode {
		Router.Use(gin.Logger())
	}

	systemRouter := router.RouterGroup

	// todo swagger

	PublicGroup := Router.Group(global.GA_CONFIG.System.RouterPrefix)
	//PrivateGroup := Router.Group(global.GA_CONFIG.System.RouterPrefix)
	{
		// 健康监测
		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, "ok")
		})
		// 注册验证码
		systemRouter.InitCaptchaRouter(PublicGroup)
	}
	return Router
}
