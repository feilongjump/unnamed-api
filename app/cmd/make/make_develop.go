package make

import (
	"github.com/spf13/cobra"
)

var CmdMakeDevelop = &cobra.Command{
	Use:   "develop",
	Short: "develop",
	Run:   runMakeDevelop,
	Args:  cobra.ExactArgs(1), // 只允许且必须传 1 个参数
}

func runMakeDevelop(cmd *cobra.Command, args []string) {

	// 创建 migrator
	migratorArgs := []string{
		"add_" + args[0] + "_table",
	}
	runMakeMigration(cmd, migratorArgs)

	// 创建 factory
	runMakeFactory(cmd, args)

	// 创建 seeder
	runMakeSeeder(cmd, args)

	// 创建 model
	runMakeModel(cmd, args)

	// 创建 policy
	runMakePolicy(cmd, args)

	// 创建 request
	runMakeRequest(cmd, args)

	// 创建 controller
	controllerArgs := []string{
		"v1/" + args[0],
	}
	runMakeAPIController(cmd, controllerArgs)
}
