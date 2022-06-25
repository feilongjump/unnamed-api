package validators

import "unnamed-api/pkg/captcha"

// ValidateCaptcha 自定义规则，验证『图片验证码』
func ValidateCaptcha(captchaId, captchaAnswer string, errs map[string][]string) map[string][]string {
	if ok := captcha.NewCaptcha().VerifyCaptcha(captchaId, captchaAnswer); !ok {
		errs["captcha_answer"] = append(errs["captcha_answer"], "图片验证码错误")
	}

	return errs
}
