package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type (
	SignupPhoneExistRequest struct {
		Phone string `json:"phone,omitempty" valid:"phone"`
	}

	SignupEmailExistRequest struct {
		Email string `json:"email,omitempty" valid:"email"`
	}
)

func SignupPhoneExist(data interface{}, ctx *gin.Context) map[string][]string {

	// 自定义验证规则
	rules := govalidator.MapData{
		"phone": []string{"required", "digits:11"},
	}

	// 自定义验证出错时的提示
	message := govalidator.MapData{
		"phone": []string{
			"required:手机号为必填项",
			"digits:手机号长度必须为 11 位的数字",
		},
	}

	return validate(data, rules, message)
}

func SignupEmailExist(data interface{}, ctx *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		"email": []string{"required", "min:4", "max:30", "email"},
	}

	message := govalidator.MapData{
		"email": []string{
			"required:Email 为必填项",
			"min:Email 长度需大于 4",
			"max:Email 长度需小于 30",
			"email:Email 格式不正确，请提供有效的邮箱地址",
		},
	}

	return validate(data, rules, message)
}
