package middlewares

import (
	"fmt"
	"unnamed-api/app/models/admin_user"
	"unnamed-api/pkg/config"
	"unnamed-api/pkg/jwt"
	"unnamed-api/pkg/response"

	"github.com/gin-gonic/gin"
)

func AuthJWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		// 从标头 Authorization:Bearer xxxxxx 中获取信息，并验证 JWT 的准确性
		claims, err := jwt.NewJWT().ParserToken(ctx)
		if err != nil {
			response.Unauthorized(ctx, fmt.Sprintf("请查看 %v 相关的接口认证文档", config.GetString("app.name")))
			return
		}

		// JWT 解析成功，设置管理员用户信息
		adminUserModel := admin_user.Get(claims.UserID)
		if adminUserModel.ID == 0 {
			response.Unauthorized(ctx, "找不到对应用户，用户可能已删除")
			return
		}

		// 将管理员用户信息存入 gin.context 里，后续 auth 包将从这里拿到当前管理员用户数据
		ctx.Set("current_admin_user_id", adminUserModel.GetStringID())
		ctx.Set("current_admin_user_name", adminUserModel.Name)
		ctx.Set("current_admin_user", adminUserModel)

		ctx.Next()
	}
}
