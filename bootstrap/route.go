package bootstrap

import (
	"net/http"
	"strings"
	"unnamed-api/app/http/middlewares"
	"unnamed-api/routes"

	"github.com/gin-gonic/gin"
)

// SetupRoute 路由初始化
func SetupRoute(router *gin.Engine) {

	// 注册全局中间件
	registerGlobalMiddleWare(router)

	// 注册 API 路由
	routes.RegisterApiRoutes(router)

	// 配置 404 路由
	setup404Handle(router)
}

func registerGlobalMiddleWare(router *gin.Engine) {
	router.Use(
		middlewares.Logger(),
		gin.Recovery(),
	)
}

func setup404Handle(router *gin.Engine) {
	router.NoRoute(func(ctx *gin.Context) {
		// 获取标头信息的 Accept 信息
		acceptString := ctx.Request.Header.Get("Accept")

		if strings.Contains(acceptString, "text/html") {
			// 如果是 html 的话
			ctx.String(http.StatusNotFound, "404 页面")
		} else {
			// 默认返回 JSON
			ctx.JSON(http.StatusNotFound, gin.H{
				"error_code":    404,
				"error_message": "路由未定义",
			})
		}
	})
}
