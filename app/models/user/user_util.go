package user

import "unnamed-api/pkg/database"

// IsEmailExist 判断 Email 已被注册
func IsEmailExist(email string) bool {
	var count int64
	database.DB.Model(User{}).Where("email = ?", email).Count(&count)
	return count > 0
}

// IsPhoneExist 判断手机号已被注册
func IsPhoneExist(phone string) bool {
	var count int64
	database.DB.Model(User{}).Where("phone = ?", phone).Count(&count)
	return count > 0
}

// GetByEmail 通过 email 获取用户信息
func GetByEmail(email string) (userModel User) {
	database.DB.Where("email = ?", email).First(&userModel)

	return
}

// GetByMulti 通过 email、phone、name 获取用户信息
func GetByMulti(username string) (userModel User) {
	database.DB.
		Where("email = ?", username).
		Or("phone = ?", username).
		Or("name = ?", username).
		First(&userModel)

	return
}
