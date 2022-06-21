package auth

import (
	"net/http"
	v1 "unnamed-api/app/http/controllers/api/v1"
	"unnamed-api/app/models/user"
	"unnamed-api/app/requests"

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
	ctx.JSON(http.StatusOK, gin.H{
		"exist": user.IsPhoneExist(request.Phone),
	})
}

func (sc *SignupController) IsEmailExist(ctx *gin.Context) {

	request := requests.SignupEmailExistRequest{}
	if ok := requests.Validate(ctx, &request, requests.SignupEmailExist); !ok {
		return
	}

	// 检查数据库并返回响应
	ctx.JSON(http.StatusOK, gin.H{
		"exist": user.IsEmailExist(request.Email),
	})
}
