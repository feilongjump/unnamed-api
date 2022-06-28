package routes

import (
	"unnamed-api/app/http/controllers/api/v1/auth"
	"unnamed-api/app/http/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterApiRoutes(router *gin.Engine) {
	v1 := router.Group("v1")

	// 全局限流中间件：每小时限流。这里是所有 API （根据 IP）请求加起来。
	v1.Use(middlewares.LimitIP("200-H"))
	{
		authGroup := v1.Group("/auth")

		authGroup.Use(middlewares.LimitIP("1000-H"))
		{
			suc := new(auth.SignupController)
			authGroup.POST("/signup/phone/exist", middlewares.GuestJWT(), suc.IsPhoneExist)
			authGroup.POST("/signup/email/exist", middlewares.GuestJWT(), suc.IsEmailExist)
			authGroup.POST("/signup/using-email", middlewares.GuestJWT(), middlewares.LimitPerRoute("60-H"), suc.SignupUsingEmail)

			// 验证码
			vcc := new(auth.VerifyCodeController)
			// 图片验证码，需要加限流
			authGroup.POST("/verify-codes/captcha", middlewares.LimitPerRoute("50-H"), vcc.ShowCaptcha)
			// 发送 Email 验证码
			authGroup.POST("/verify-codes/email", middlewares.LimitPerRoute("20-H"), vcc.SendUsingEmail)

			// 登录
			lgc := new(auth.LoginController)
			// 使用邮箱和验证码进行登录
			authGroup.POST("/login/using-email", middlewares.GuestJWT(), lgc.LoginByEmail)
			// 使用邮箱、手机号、用户名 + 密码进行登录
			authGroup.POST("/login/using-password", middlewares.GuestJWT(), lgc.LoginByPassword)

			// 重置密码
			pwc := new(auth.PasswordController)
			authGroup.POST("/password-reset/using-email", middlewares.GuestJWT(), pwc.ResetByEmail)
		}
	}
}
