package cmd

import (
	"go-web-scaffold/internal/config"
	"go-web-scaffold/internal/database"
	"go-web-scaffold/internal/handlers"
	"go-web-scaffold/internal/kafka"
	"go-web-scaffold/internal/logging"
	"go-web-scaffold/internal/repository"
	"go-web-scaffold/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var server1Cmd = &cobra.Command{
	Use:   "server1",
	Short: "Start server 1",
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.LoadConfig()
		database.InitDB(cfg.MySQLDSN)
		database.InitRedis(cfg.RedisAddr, cfg.RedisPassword, cfg.RedisDB)
		kafka.InitKafka(cfg.KafkaBrokers)

		userRepository := repository.NewUserRepository(database.DB)
		userService := service.NewUserService(userRepository)
		userHandler := handlers.NewUserHandler(userService)

		router := gin.Default()

		router.POST("/users", userHandler.CreateUser)
		router.GET("/users/:id", userHandler.GetUserByID)
		router.GET("/users", userHandler.GetAllUsers)
		router.PUT("/users/:id", userHandler.UpdateUser)
		router.DELETE("/users/:id", userHandler.DeleteUser)

		logging.Logger.Info("Starting server 1", zap.String("address", cfg.Server1Addr))
		if err := router.Run(cfg.Server1Addr); err != nil {
			logging.Logger.Fatal("Server 1 failed to start", zap.Error(err))
		}
	},
}
