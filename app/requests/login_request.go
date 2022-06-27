package requests

import (
	"unnamed-api/app/requests/validators"

	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type (
	LoginByEmailRequest struct {
		Email      string `json:"email,omitempty" valid:"email"`
		VerifyCode string `json:"verify_code,omitempty" valid:"verify_code"`
	}

	LoginByPasswordRequest struct {
		UserName      string `json:"username,omitempty" valid:"username"`
		Password      string `json:"password,omitempty" valid:"password"`
		CaptchaID     string `json:"captcha_id,omitempty" valid:"captcha_id"`
		CaptchaAnswer string `json:"captcha_answer,omitempty" valid:"captcha_answer"`
	}
)

func LoginByEmail(data interface{}, ctx *gin.Context) map[string][]string {

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

func LoginByPassword(data interface{}, ctx *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		"username":       []string{"required", "min:3", "max:30"},
		"password":       []string{"required", "min:6"},
		"captcha_id":     []string{"required"},
		"captcha_answer": []string{"required", "digits:6"},
	}

	messages := govalidator.MapData{
		"username": []string{
			"required:用户名为必填项，支持手机号、邮箱和用户名",
			"min:用户名长度需大于 3",
			"max:用户名长度需小于 30",
		},
		"password": []string{
			"required:密码为必填项",
			"min:密码长度需大于 6",
		},
		"captcha_id": []string{
			"required:图片验证码的 ID 为必填",
		},
		"captcha_answer": []string{
			"required:图片验证码答案必填",
			"digits:图片验证码长度必须为 6 位的数字",
		},
	}

	errs := validate(data, rules, messages)

	_data := data.(*LoginByPasswordRequest)
	errs = validators.ValidateCaptcha(_data.CaptchaID, _data.CaptchaAnswer, errs)

	return errs
}
