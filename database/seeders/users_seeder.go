package seeders

import (
	"fmt"
	"unnamed-api/database/factories"
	"unnamed-api/pkg/console"
	"unnamed-api/pkg/logger"
	"unnamed-api/pkg/seed"

	"gorm.io/gorm"
)

func init() {
	// 添加 Seeder

	seed.Add("SeedUsersTable", func(db *gorm.DB) {
		// 创建 10 个用户对象
		users := factories.MakeUsers(10)

		// 批量创建用户（批量创建不会调用模型钩子）
		result := db.Table("users").Create(&users)
		if err := result.Error; err != nil {
			logger.LogIf(err)
			return
		}

		// 打印运行情况
		console.Success(fmt.Sprintf("Table [%v] %v rows seeded", result.Statement.Table, result.RowsAffected))
	})
}
