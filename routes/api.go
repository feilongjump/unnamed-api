package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterApiRoutes(router *gin.Engine) {
	v1 := router.Group("v1")
	{
		v1.GET("/", func(ctx *gin.Context) {
			// 以 JSON 格式响应
			ctx.JSON(http.StatusOK, gin.H{
				"Hello": "World",
			})
		})
	}
}
