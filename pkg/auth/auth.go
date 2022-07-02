package auth

import (
	"errors"
	"unnamed-api/app/models/admin_user"
	"unnamed-api/pkg/logger"

	"github.com/gin-gonic/gin"
)

// LoginByEmail 登录指定 email 管理员用户
func LoginByEmail(email string) (admin_user.AdminUser, error) {
	adminUserModel := admin_user.GetBy("email", email)
	if adminUserModel.ID == 0 {
		return admin_user.AdminUser{}, errors.New("邮箱尚未注册")
	}

	return adminUserModel, nil
}

// Attempt 尝试用邮箱、手机号、用户名 + 密码登录
func Attempt(username, password string) (admin_user.AdminUser, error) {
	adminUserModel := admin_user.GetByMulti(username)
	if adminUserModel.ID == 0 {
		return admin_user.AdminUser{}, errors.New("账号不存在")
	}

	if !adminUserModel.ComparePassword(password) {
		return admin_user.AdminUser{}, errors.New("密码错误")
	}

	return adminUserModel, nil
}

// CurrentAdminUser 从 gin.context 中获取当前登录管理员用户
func CurrentAdminUser(ctx *gin.Context) admin_user.AdminUser {
	adminUserModel, ok := ctx.MustGet("current_admin_user").(admin_user.AdminUser)
	if !ok {
		logger.LogIf(errors.New("无法获取用户"))
		return admin_user.AdminUser{}
	}

	// db is now a *DB value
	return adminUserModel
}

// CurrentUID 从 gin.context 中获取当前登录管理员用户 ID
func CurrentUID(ctx *gin.Context) string {
	return ctx.GetString("current_admin_user_id")
}
