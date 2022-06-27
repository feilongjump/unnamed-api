package auth

import (
	"errors"
	"unnamed-api/app/models/user"
	"unnamed-api/pkg/logger"

	"github.com/gin-gonic/gin"
)

// LoginByEmail 登录指定 email 用户
func LoginByEmail(email string) (user.User, error) {
	userModel := user.GetByEmail(email)
	if userModel.ID == 0 {
		return user.User{}, errors.New("邮箱尚未注册")
	}

	return userModel, nil
}

// Attempt 尝试用邮箱、手机号、用户名 + 密码登录
func Attempt(username, password string) (user.User, error) {
	userModel := user.GetByMulti(username)
	if userModel.ID == 0 {
		return user.User{}, errors.New("账号不存在")
	}

	if !userModel.ComparePassword(password) {
		return user.User{}, errors.New("密码错误")
	}

	return userModel, nil
}

// CurrentUser 从 gin.context 中获取当前登录用户
func CurrentUser(ctx *gin.Context) user.User {
	userModel, ok := ctx.MustGet("current_user").(user.User)
	if !ok {
		logger.LogIf(errors.New("无法获取用户"))
		return user.User{}
	}

	// db is now a *DB value
	return userModel
}

// CurrentUID 从 gin.context 中获取当前登录用户 ID
func CurrentUID(ctx *gin.Context) string {
	return ctx.GetString("current_user_id")
}
