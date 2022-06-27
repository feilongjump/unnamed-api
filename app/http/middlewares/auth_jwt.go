package middlewares

import (
	"fmt"
	"unnamed-api/app/models/user"
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

		// JWT 解析成功，设置用户信息
		userModel := user.Get(claims.UserID)
		if userModel.ID == 0 {
			response.Unauthorized(ctx, "找不到对应用户，用户可能已删除")
			return
		}

		// 将用户信息存入 gin.context 里，后续 auth 包将从这里拿到当前用户数据
		ctx.Set("current_user_id", userModel.GetStringID())
		ctx.Set("current_user_name", userModel.Name)
		ctx.Set("current_user", userModel)

		ctx.Next()
	}
}
