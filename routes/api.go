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

		manufacturerRouter(v1)

		materialRouter(v1)
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

	cbc := new(controllers.CustomerBanksController)
	customerBanksGroup := customersGroup.Group("/:customer/banks")
	{
		customerBanksGroup.GET("", middlewares.AuthJWT(), cbc.Index)
		customerBanksGroup.GET("/:bank", middlewares.AuthJWT(), cbc.Show)
		customerBanksGroup.POST("", middlewares.AuthJWT(), cbc.Store)
		customerBanksGroup.PUT("/:bank", middlewares.AuthJWT(), cbc.Update)
		customerBanksGroup.DELETE("/:bank", middlewares.AuthJWT(), cbc.Delete)
	}

}

// manufacturer
func manufacturerRouter(routerGroup *gin.RouterGroup) {
	mc := new(controllers.ManufacturersController)

	manufacturersGroup := routerGroup.Group("/manufacturers")
	{
		manufacturersGroup.GET("", middlewares.AuthJWT(), mc.Index)
		manufacturersGroup.GET("/:manufacturer", middlewares.AuthJWT(), mc.Show)
		manufacturersGroup.POST("", middlewares.AuthJWT(), mc.Store)
		manufacturersGroup.PUT("/:manufacturer", middlewares.AuthJWT(), mc.Update)
		manufacturersGroup.DELETE("/:manufacturer", middlewares.AuthJWT(), mc.Delete)
	}

	mcc := new(controllers.ManufacturerContactsController)
	manufacturerContactsGroup := manufacturersGroup.Group("/:manufacturer/contacts")
	{
		manufacturerContactsGroup.GET("", middlewares.AuthJWT(), mcc.Index)
		manufacturerContactsGroup.GET("/:contact", middlewares.AuthJWT(), mcc.Show)
		manufacturerContactsGroup.POST("", middlewares.AuthJWT(), mcc.Store)
		manufacturerContactsGroup.PUT("/:contact", middlewares.AuthJWT(), mcc.Update)
		manufacturerContactsGroup.DELETE("/:contact", middlewares.AuthJWT(), mcc.Delete)
	}

	mbc := new(controllers.ManufacturerBanksController)
	manufacturerBanksGroup := manufacturersGroup.Group("/:manufacturer/banks")
	{
		manufacturerBanksGroup.GET("", middlewares.AuthJWT(), mbc.Index)
		manufacturerBanksGroup.GET("/:bank", middlewares.AuthJWT(), mbc.Show)
		manufacturerBanksGroup.POST("", middlewares.AuthJWT(), mbc.Store)
		manufacturerBanksGroup.PUT("/:bank", middlewares.AuthJWT(), mbc.Update)
		manufacturerBanksGroup.DELETE("/:bank", middlewares.AuthJWT(), mbc.Delete)
	}

}

// material
func materialRouter(routerGroup *gin.RouterGroup) {
	mc := new(controllers.MaterialsController)

	materialsGroup := routerGroup.Group("/materials")
	{
		materialsGroup.GET("", middlewares.AuthJWT(), mc.Index)
		materialsGroup.GET("/:material", middlewares.AuthJWT(), mc.Show)
		materialsGroup.POST("", middlewares.AuthJWT(), mc.Store)
		materialsGroup.PUT("/:material", middlewares.AuthJWT(), mc.Update)
		materialsGroup.DELETE("/:material", middlewares.AuthJWT(), mc.Delete)
	}
}
