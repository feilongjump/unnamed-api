package requests

import (
	"unnamed-api/app/requests/validators"

	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type LoginByEmailRequest struct {
	Email      string `json:"email,omitempty" valid:"email"`
	VerifyCode string `json:"verify_code,omitempty" valid:"verify_code"`
}

// LoginByEmail 验证表单，返回长度等于零即通过
func LoginByEmail(data interface{}, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		"email":       []string{"required", "min:4", "max:30", "email"},
		"verify_code": []string{"required", "digits:6"},
	}

	messages := govalidator.MapData{
		"email": []string{
			"required:Email 为必填项",
			"min:Email 长度需大于 4",
			"max:Email 长度需小于 30",
			"email:Email 格式不正确，请提供有效的邮箱地址",
		},
		"verify_code": []string{
			"required:验证码答案必填",
			"digits:验证码长度必须为 6 位的数字",
		},
	}

	errs := validate(data, rules, messages)

	_data := data.(*LoginByEmailRequest)
	errs = validators.ValidateVerifyCode(_data.Email, _data.VerifyCode, errs)

	return errs
}
