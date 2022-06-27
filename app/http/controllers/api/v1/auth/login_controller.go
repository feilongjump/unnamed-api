package auth

import (
	v1 "unnamed-api/app/http/controllers/api/v1"
	"unnamed-api/app/requests"
	"unnamed-api/pkg/auth"
	"unnamed-api/pkg/jwt"
	"unnamed-api/pkg/response"

	"github.com/gin-gonic/gin"
)

type LoginController struct {
	v1.BaseAPIController
}

// LoginByEmail 邮箱登录
func (lgc *LoginController) LoginByEmail(ctx *gin.Context) {
	// 表单验证
	request := requests.LoginByEmailRequest{}
	if ok := requests.Validate(ctx, &request, requests.LoginByEmail); !ok {
		return
	}

	// 尝试登录
	user, err := auth.LoginByEmail(request.Email)
	if err != nil {
		// 登录失败
		response.Error(ctx, err, "账号不存在或密码错误")
	} else {
		// 登录成功
		token := jwt.NewJWT().IssueToken(user.GetStringID(), user.Name)

		response.JSON(ctx, gin.H{
			"token": token,
		})
	}
}
