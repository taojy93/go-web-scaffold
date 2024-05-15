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
	// go run main.go server1
	rootCmd.AddCommand(server1Cmd)
	// go run main.go server2
	rootCmd.AddCommand(server2Cmd)
}
