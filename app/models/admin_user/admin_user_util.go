package admin_user

import (
	"fmt"
	"unnamed-api/pkg/app"
	"unnamed-api/pkg/database"
	"unnamed-api/pkg/paginator"

	"github.com/gin-gonic/gin"
)

func Get(idStr string) (adminUser AdminUser) {
	database.DB.Where("id", idStr).First(&adminUser)
	return
}

func GetBy(field, value string) (adminUser AdminUser) {
	database.DB.Where(fmt.Sprintf("%v = ?", field), value).First(&adminUser)
	return
}

func All() (adminUsers []AdminUser) {
	database.DB.Find(&adminUsers)
	return
}

func IsExist(field, value string) bool {
	var count int64
	database.DB.Model(AdminUser{}).Where(fmt.Sprintf("%v = ?", field), value).Count(&count)
	return count > 0
}

// Paginate 分页内容
func Paginate(ctx *gin.Context, perPage int) (adminUsers []AdminUser, paging paginator.Paging) {
	paging = paginator.Paginate(
		ctx,
		database.DB.Model(AdminUser{}),
		&adminUsers,
		app.V1URL(database.TableName(&AdminUser{})),
		perPage,
	)

	return
}

// GetByMulti 通过 email、phone、name 获取管理员用户信息
func GetByMulti(username string) (adminUser AdminUser) {
	database.DB.
		Where("email = ?", username).
		Or("phone = ?", username).
		Or("name = ?", username).
		First(&adminUser)

	return
}
