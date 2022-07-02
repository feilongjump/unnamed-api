package requests

import (
	"unnamed-api/app/requests/validators"

	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type (
	SignupEmailExistRequest struct {
		Email string `json:"email,omitempty" valid:"email"`
	}

	// SignupUsingEmailRequest 通过邮箱注册的请求信息
	SignupUsingEmailRequest struct {
		Email           string `json:"email,omitempty" valid:"email"`
		VerifyCode      string `json:"verify_code,omitempty" valid:"verify_code"`
		Name            string `json:"name" valid:"name"`
		Password        string `json:"password,omitempty" valid:"password"`
		PasswordConfirm string `json:"password_confirm,omitempty" valid:"password_confirm"`
	}
)

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

func SignupUsingEmail(data interface{}, ctx *gin.Context) map[string][]string {
	rules := govalidator.MapData{
		"email":            []string{"required", "min:4", "max:30", "email", "not_exists:admin_users,email"},
		"name":             []string{"required", "alpha_num", "between:3,20", "not_exists:admin_users,name"},
		"password":         []string{"required", "min:6"},
		"password_confirm": []string{"required"},
		"verify_code":      []string{"required", "digits:6"},
	}

	messages := govalidator.MapData{
		"email": []string{
			"required:Email 为必填项",
			"min:Email 长度需大于 4",
			"max:Email 长度需小于 30",
			"email:Email 格式不正确，请提供有效的邮箱地址",
			"not_exists:Email 已被占用",
		},
		"name": []string{
			"required:用户名为必填项",
			"alpha_num:用户名格式错误，只允许数字和英文",
			"between:用户名长度需在 3~20 之间",
		},
		"password": []string{
			"required:密码为必填项",
			"min:密码长度需大于 6",
		},
		"password_confirm": []string{
			"required:确认密码框为必填项",
		},
		"verify_code": []string{
			"required:验证码答案必填",
			"digits:验证码长度必须为 6 位的数字",
		},
	}

	errs := validate(data, rules, messages)

	_data := data.(*SignupUsingEmailRequest)
	errs = validators.ValidatePasswordConfirm(_data.Password, _data.PasswordConfirm, errs)
	errs = validators.ValidateVerifyCode(_data.Email, _data.VerifyCode, errs)

	return errs
}
