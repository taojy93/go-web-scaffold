package cmd

import (
	"go-web-scaffold/internal/config"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "app",
	Short: "A multi-service application",
}

func Execute() {
	// 初始化加载配置
	config.InitConfig()
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(server1Cmd)
	rootCmd.AddCommand(server2Cmd)
}
