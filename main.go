package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	// new 一个 Gin 实例
	r := gin.New()

	// 注册中间件
	r.Use(gin.Logger(), gin.Recovery())

	// 注册一个路由
	r.GET("/", func(c *gin.Context) {

		// 以 JSON 格式响应
		c.JSON(http.StatusOK, gin.H{
			"Hello": "World!",
		})
	})

	r.NoRoute(func(ctx *gin.Context) {
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

	// 运行服务
	r.Run()
}
