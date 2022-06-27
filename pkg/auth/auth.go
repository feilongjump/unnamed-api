package auth

import (
	"errors"
	"unnamed-api/app/models/user"
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
