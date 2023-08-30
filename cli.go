package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"mumu/app/cmd"
	"mumu/bootstrap"
	"mumu/config"
	"mumu/pkg/console"
	"os"
)

func init() {
	config.Initialize()
}

func main() {
	var rootCmd = &cobra.Command{
		Use:   "XX",
		Short: "Cli",
		Long:  "Cli   asasasas",

		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			// 初始化 Logger
			bootstrap.SetupLogger()
			// 初始化 SQL
			bootstrap.SetupDB()

			// 初始化 Redis
			bootstrap.SetupRedis()

			// 初始化缓存
			bootstrap.SetupCache()
		},
	}

	rootCmd.AddCommand(
		cmd.CmdCache,
		cmd.CmdKey,
		cmd.CmdPlay,
		cmd.CmdSearch,
	)

	cmd.RegisterGlobalFlags(rootCmd)

	// 执行主命令
	if err := rootCmd.Execute(); err != nil {
		console.Exit(fmt.Sprintf("Failed to run app with %v: %s", os.Args, err.Error()))
	}
}
