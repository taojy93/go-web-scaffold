package cmd

import (
	"go-web-scaffold/internal/config"
	"go-web-scaffold/internal/database"
	"go-web-scaffold/internal/logging"

	"go-web-scaffold/internal/kafka"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var server2Cmd = &cobra.Command{
	Use:   "server2",
	Short: "Start server 2",
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.LoadConfig()
		database.InitDB(cfg.MySQLDSN)
		database.InitRedis(cfg.RedisAddr, cfg.RedisPassword, cfg.RedisDB)
		kafka.InitKafka(cfg.KafkaBrokers)

		router := gin.Default()

		router.GET("/", func(c *gin.Context) {
			kafka.SendMessage("topic2", "Hello from Server 2")
			c.String(200, "Hello from Server 2")
		})

		logging.Logger.Info("Starting server 2", zap.String("address", cfg.Server2Addr))
		if err := router.Run(cfg.Server2Addr); err != nil {
			logging.Logger.Fatal("Server 2 failed to start", zap.Error(err))
		}
	},
}
