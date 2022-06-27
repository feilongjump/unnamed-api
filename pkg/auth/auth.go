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
