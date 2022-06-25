package verifycode

import (
	"strings"
	"sync"
	"unnamed-api/pkg/app"
	"unnamed-api/pkg/config"
	"unnamed-api/pkg/helpers"
	"unnamed-api/pkg/logger"
	"unnamed-api/pkg/redis"
)

type VerifyCode struct {
	Store Store
}

var once sync.Once
var internalVerifyCode *VerifyCode

// NewVerifyCode 单例模式获取
func NewVerifyCode() *VerifyCode {
	once.Do(func() {
		internalVerifyCode = &VerifyCode{
			Store: &RedisStore{
				RedisClient: redis.Redis,
				// 增加前缀保持数据库整洁
				KeyPrefix: config.GetString("app.name") + ":verifycode:",
			},
		}
	})

	return internalVerifyCode
}

// CheckAnswer 检查用户提交的验证码是否正确
func (vc *VerifyCode) CheckAnswer(key, answer string) bool {
	logger.DebugJSON("验证码", "检查验证码", map[string]string{key: answer})

	// 方便开发，在非生产环境下，具备特殊前缀的手机号码和 Email 后缀，会直接验证成功
	if !app.IsProduction() &&
		(strings.HasSuffix(key, config.GetString("verifycode.debug_email_suffix")) ||
			strings.HasPrefix(key, config.GetString("verifycode.debug_phone_prefix"))) {
		return true
	}

	return vc.Store.Verify(key, answer, false)
}

// generateVerifyCode 生成验证码，并放置于 Redis 中
func (vc *VerifyCode) generateVerifyCode(key string) string {
	// 生成随机码
	code := helpers.RandomNumber(config.GetInt("verifycode.code_length"))

	// 方便开发，本地环境使用固定验证码
	if app.IsLocal() {
		code = config.GetString("verifycode.debug_code")
	}

	logger.DebugJSON("验证码", "生成验证码", map[string]string{key: code})

	// 将验证码及 KEY （邮箱或手机号码）存放到 Redis 中并设置过期时间
	vc.Store.Set(key, code)
	return code
}
