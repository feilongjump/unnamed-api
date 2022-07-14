package routes

import (
	controllers "unnamed-api/app/http/controllers/api/v1"
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
		authRouter(authGroup)

		adminUserRouter(v1)

		customerRouter(v1)
	}
}

// auth
func authRouter(routerGroup *gin.RouterGroup) {
	routerGroup.Use(middlewares.LimitIP("1000-H"))
	{
		suc := new(auth.SignupController)
		routerGroup.POST("/signup/email/exist", middlewares.GuestJWT(), suc.IsEmailExist)
		routerGroup.POST("/signup/using-email", middlewares.GuestJWT(), middlewares.LimitPerRoute("60-H"), suc.SignupUsingEmail)

		// 验证码
		vcc := new(auth.VerifyCodeController)
		// 图片验证码，需要加限流
		routerGroup.POST("/verify-codes/captcha", middlewares.LimitPerRoute("50-H"), vcc.ShowCaptcha)
		// 发送 Email 验证码
		routerGroup.POST("/verify-codes/email", middlewares.LimitPerRoute("20-H"), vcc.SendUsingEmail)

		// 登录
		lgc := new(auth.LoginController)
		// 使用邮箱和验证码进行登录
		routerGroup.POST("/login/using-email", middlewares.GuestJWT(), lgc.LoginByEmail)
		// 使用邮箱、手机号、用户名 + 密码进行登录
		routerGroup.POST("/login/using-password", middlewares.GuestJWT(), lgc.LoginByPassword)

		// 重置密码
		pwc := new(auth.PasswordController)
		routerGroup.POST("/password-reset/using-email", middlewares.GuestJWT(), pwc.ResetByEmail)
	}
}

// admin_user
func adminUserRouter(routerGroup *gin.RouterGroup) {
	auc := new(controllers.AdminUsersController)
	// 获取当前管理员用户
	routerGroup.GET("/admin_user", middlewares.AuthJWT(), auc.CurrentAdminUser)

	adminUsersGroup := routerGroup.Group("/admin_users")
	{
		adminUsersGroup.GET("", auc.Index)
	}
}

// customer
func customerRouter(routerGroup *gin.RouterGroup) {
	cc := new(controllers.CustomersController)

	customersGroup := routerGroup.Group("/customers")
	{
		customersGroup.GET("", middlewares.AuthJWT(), cc.Index)
		customersGroup.GET("/:customer", middlewares.AuthJWT(), cc.Show)
		customersGroup.POST("", middlewares.AuthJWT(), cc.Store)
		customersGroup.PUT("/:customer", middlewares.AuthJWT(), cc.Update)
		customersGroup.DELETE("/:customer", middlewares.AuthJWT(), cc.Delete)
	}

	ccc := new(controllers.CustomerContactsController)
	customerContactsGroup := customersGroup.Group("/:customer/contacts")
	{
		customerContactsGroup.GET("", middlewares.AuthJWT(), ccc.Index)
		customerContactsGroup.GET("/:contact", middlewares.AuthJWT(), ccc.Show)
		customerContactsGroup.POST("", middlewares.AuthJWT(), ccc.Store)
		customerContactsGroup.PUT("/:contact", middlewares.AuthJWT(), ccc.Update)
		customerContactsGroup.DELETE("/:contact", middlewares.AuthJWT(), ccc.Delete)
	}

}
