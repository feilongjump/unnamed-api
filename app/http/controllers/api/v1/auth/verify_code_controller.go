package auth

import (
	v1 "unnamed-api/app/http/controllers/api/v1"
	"unnamed-api/app/requests"
	"unnamed-api/pkg/captcha"
	"unnamed-api/pkg/logger"
	"unnamed-api/pkg/response"
	"unnamed-api/pkg/verifycode"

	"github.com/gin-gonic/gin"
)

// VerifyCodeController 用户控制器
type VerifyCodeController struct {
	v1.BaseAPIController
}

// ShowCaptcha 显示图片验证码
func (vc *VerifyCodeController) ShowCaptcha(ctx *gin.Context) {
	// 生成验证码
	id, b64s, err := captcha.NewCaptcha().GenerateCaptcha()

	// 记录错误日志，因为验证码是用户的入口，出错时应该记 error 等级的日志
	logger.LogIf(err)

	// 返回给用户
	response.JSON(ctx, gin.H{
		"captcha_id":    id,
		"captcha_image": b64s,
	})
}

// SendUsingEmail 发送 Email 验证码
func (vc *VerifyCodeController) SendUsingEmail(ctx *gin.Context) {
	// 验证表单
	request := requests.VerifyCodeEmailRequest{}
	if ok := requests.Validate(ctx, &request, requests.VerifyCodeEmail); !ok {
		return
	}

	// 发送 Email
	if err := verifycode.NewVerifyCode().SendEmail(request.Email); err != nil {
		response.Abort500(ctx, "发送 Email 验证码失败！")
	} else {
		response.Success(ctx)
	}
}
