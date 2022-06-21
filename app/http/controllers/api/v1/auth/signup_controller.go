package auth

import (
	"fmt"
	"net/http"
	v1 "unnamed-api/app/http/controllers/api/v1"
	"unnamed-api/app/models/user"

	"github.com/gin-gonic/gin"
)

// SignupController 注册控制器
type SignupController struct {
	v1.BaseAPIController
}

// IsPhoneExist 检测手机号是否被注册
func (sc *SignupController) IsPhoneExist(ctx *gin.Context) {

	// 请求对象
	type PhoneExistRequest struct {
		Phone string `json:"phone"`
	}
	request := PhoneExistRequest{}

	// 解析 JSON 请求
	if err := ctx.ShouldBindJSON(&request); err != nil {
		// 解析失败，返回 422 状态码和错误信息
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		// 打印错误信息
		fmt.Println(err.Error())
		// 出错了，中断请求
		return
	}

	// 检查数据库并返回响应
	ctx.JSON(http.StatusOK, gin.H{
		"exist": user.IsPhoneExist(request.Phone),
	})
}
