package factories

import (
	"unnamed-api/app/models/admin_user"
	"unnamed-api/pkg/helpers"

	"github.com/bxcodec/faker/v3"
)

func MakeAdminUsers(count int) []admin_user.AdminUser {

	var objs []admin_user.AdminUser

	// 设置唯一性，如 AdminUser 模型的某个字段需要唯一，即可取消注释
	// faker.SetGenerateUniqueValues(true)

	for i := 0; i < count; i++ {
		adminUserModel := admin_user.AdminUser{
			Name:     faker.Username(),
			Email:    faker.Email(),
			Phone:    helpers.RandomNumber(11),
			Password: "$2a$14$oPzVkIdwJ8KqY0erYAYQxOuAAlbI/sFIsH0C0R4MPc.3JbWWSuaUe",
		}
		objs = append(objs, adminUserModel)
	}

	return objs
}
