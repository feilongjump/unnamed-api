package user

import (
	"unnamed-api/pkg/app"
	"unnamed-api/pkg/database"
	"unnamed-api/pkg/paginator"

	"github.com/gin-gonic/gin"
)

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

// Get 通过 ID 获取用户
func Get(idStr string) (userModel User) {
	database.DB.Where("id", idStr).First(&userModel)
	return
}

// All 获取所有用户数据
func All() (users []User) {
	database.DB.Find(&users)
	return
}

// Paginate 分页内容
func Paginate(ctx *gin.Context, perPage int) (users []User, paging paginator.Paging) {
	paging = paginator.Paginate(
		ctx,
		database.DB.Model(User{}),
		&users,
		app.V1URL(database.TableName(&User{})),
		perPage,
	)

	return
}
