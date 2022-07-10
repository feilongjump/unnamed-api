package seeders

import "unnamed-api/pkg/seed"

func Initialize() {
	// 触发记载本目录下其他文件中的 init 方法

	// 指定优先于同目录下的其他文件运行
	seed.SetRunOrder([]string{
		"SeedAdminUsersTable",
		"SeedCustomersTable",
	})
}
