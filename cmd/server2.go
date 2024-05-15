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

var server2Cmd = &cobra.Command{
	Use:   "server2",
	Short: "Start server 2",
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.LoadConfig()

		// Initialize database
		database.InitDB(cfg.MySQLDSN)
		database.InitRedis(cfg.RedisAddr, cfg.RedisPassword, cfg.RedisDB)
		kafka.InitKafka(cfg.KafkaBrokers)

		// Initialize repositories, services, and handlers
		newsRepo := repository.NewNewsRepository(database.DB)
		newsService := service.NewNewsService(newsRepo)
		newsHandler := handlers.NewNewsHandler(newsService)

		// Initialize Gin engine
		router := gin.Default()

		// Define routes
		router.GET("/", func(c *gin.Context) {
			kafka.SendMessage("topic2", "Hello from Server 2")
			c.String(200, "Hello from Server 2")
		})

		v1 := router.Group("/api/v1")
		{
			newsGroup := v1.Group("/news")
			{
				newsGroup.POST("/", newsHandler.CreateNews)
				newsGroup.GET("/:id", newsHandler.GetNewsByID)
				newsGroup.GET("/", newsHandler.GetAllNews)
				newsGroup.PUT("/", newsHandler.UpdateNews)
				newsGroup.DELETE("/:id", newsHandler.DeleteNews)
			}
		}

		// Start server
		logging.Logger.Info("Starting server 2", zap.String("address", cfg.Server2Addr))
		if err := router.Run(cfg.Server2Addr); err != nil {
			logging.Logger.Fatal("Server 2 failed to start", zap.Error(err))
		}
	},
}

func init() {
	rootCmd.AddCommand(server2Cmd)
}
