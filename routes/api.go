package routes

import (
	"unnamed-api/app/http/controllers/api/v1/auth"

	"github.com/gin-gonic/gin"
)

func RegisterApiRoutes(router *gin.Engine) {
	v1 := router.Group("v1")
	{
		authGroup := v1.Group("/auth")
		{
			src := new(auth.SignupController)
			authGroup.POST("/signup/phone/exist", src.IsPhoneExist)
			authGroup.POST("/signup/email/exist", src.IsEmailExist)
		}
	}
}
