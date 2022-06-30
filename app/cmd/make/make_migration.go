package make

import (
	"fmt"
	"strings"
	"unnamed-api/pkg/app"
	"unnamed-api/pkg/console"
	"unnamed-api/pkg/str"

	"github.com/spf13/cobra"
)

var CmdMakeMigration = &cobra.Command{
	Use:   "migration",
	Short: "Create a migration file, example: make migration add_users_table",
	Run:   runMakeMigration,
	Args:  cobra.ExactArgs(1), // 只允许且必须传 1 哥参数
}

func runMakeMigration(cmd *cobra.Command, args []string) {
	// 日期格式化
	timeStr := app.TimenowInTimezone().Format("2006_01_02_150405")

	model := makeModelFromString(args[0])
	// 截取中间部分值，替换 model.StructName 值
	// 例如：add_user_table -> User
	argSplit := strings.Split(args[0], "_")
	argSplit = argSplit[1 : len(argSplit)-1]
	model.StructName = str.Singular(str.Camel(strings.Join(argSplit, "_")))

	fileName := timeStr + "_" + model.PackageName
	filePath := fmt.Sprintf("database/migrations/%s.go", fileName)
	createFileFromStub(filePath, "migration", model, map[string]string{"{{FileName}}": fileName})

	console.Success("Migration file created, after modify it, use `migrate up` to migrate database.")
}
