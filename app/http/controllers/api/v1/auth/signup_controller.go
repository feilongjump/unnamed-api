package auth

import (
	v1 "unnamed-api/app/http/controllers/api/v1"
	"unnamed-api/app/models/user"
	"unnamed-api/app/requests"
	"unnamed-api/pkg/jwt"
	"unnamed-api/pkg/response"

	"github.com/gin-gonic/gin"
)

// SignupController 注册控制器
type SignupController struct {
	v1.BaseAPIController
}

// IsPhoneExist 检测手机号是否被注册
func (sc *SignupController) IsPhoneExist(ctx *gin.Context) {
	// 获取请求参数，并做表单验证
	request := requests.SignupPhoneExistRequest{}
	if ok := requests.Validate(ctx, &request, requests.SignupPhoneExist); !ok {
		return
	}

	// 检查数据库并返回响应
	response.JSON(ctx, gin.H{
		"exist": user.IsPhoneExist(request.Phone),
	})
}

// IsEmailExist 检测 Email 是否被注册
func (sc *SignupController) IsEmailExist(ctx *gin.Context) {

	request := requests.SignupEmailExistRequest{}
	if ok := requests.Validate(ctx, &request, requests.SignupEmailExist); !ok {
		return
	}

	// 检查数据库并返回响应
	response.JSON(ctx, gin.H{
		"exist": user.IsEmailExist(request.Email),
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
	userModel := user.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	}
	userModel.Create()

	if userModel.ID > 0 {
		token := jwt.NewJWT().IssueToken(userModel.GetStringID(), userModel.Name)
		response.CreatedJSON(ctx, gin.H{
			"token": token,
			"data":  userModel,
		})
	} else {
		response.Abort500(ctx, "创建用户失败，请稍后尝试~")
	}
}
