package auth

import (
	v1 "unnamed-api/app/http/controllers/api/v1"
	"unnamed-api/app/models/admin_user"
	"unnamed-api/app/requests"
	"unnamed-api/pkg/jwt"
	"unnamed-api/pkg/response"

	"github.com/gin-gonic/gin"
)

// SignupController 注册控制器
type SignupController struct {
	v1.BaseAPIController
}

// IsEmailExist 检测 Email 是否被注册
func (sc *SignupController) IsEmailExist(ctx *gin.Context) {

	request := requests.SignupEmailExistRequest{}
	if ok := requests.Validate(ctx, &request, requests.SignupEmailExist); !ok {
		return
	}

	// 检查数据库并返回响应
	response.JSON(ctx, gin.H{
		"exist": admin_user.IsExist("email", request.Email),
	})
}

// SignupUsingEmail 使用 Email + 验证码进行注册
func (sc *SignupController) SignupUsingEmail(ctx *gin.Context) {
	// 验证表单
	request := requests.SignupUsingEmailRequest{}
	if ok := requests.Validate(ctx, &request, requests.SignupUsingEmail); !ok {
		return
	}

	// 验证成功，创建数据
	adminUserModel := admin_user.AdminUser{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	}
	adminUserModel.Create()

	if adminUserModel.ID > 0 {
		token := jwt.NewJWT().IssueToken(adminUserModel.GetStringID(), adminUserModel.Name)
		response.CreatedJSON(ctx, gin.H{
			"token": token,
			"data":  adminUserModel,
		})
	} else {
		response.Abort500(ctx, "创建用户失败，请稍后尝试~")
	}
}
