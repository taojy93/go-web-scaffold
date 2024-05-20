// internal/config/config.go
package config

import (
	"go-web-scaffold/internal/logging"
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Server1Addr   string
	Server2Addr   string
	MySQLDSN      string
	RedisAddr     string
	RedisPassword string
	RedisDB       int
	KafkaBrokers  []string
}

func LoadConfig() *Config {
	viper.AutomaticEnv()

	// viper.SetDefault("SERVER1_ADDR", ":4001")
	// viper.SetDefault("SERVER2_ADDR", ":4002")
	// viper.SetDefault("MYSQL_DSN", "root:root@tcp(127.0.0.1:3306)/antiy_license?charset=utf8mb4&parseTime=True&loc=Local")
	// viper.SetDefault("REDIS_ADDR", "localhost:6379")
	// viper.SetDefault("REDIS_PASSWORD", "root")
	// viper.SetDefault("REDIS_DB", 0)
	// viper.SetDefault("KAFKA_BROKERS", "localhost:9092")

	config := &Config{
		Server1Addr:   viper.GetString("SERVER1_ADDR"),
		Server2Addr:   viper.GetString("SERVER2_ADDR"),
		MySQLDSN:      viper.GetString("MYSQL_DSN"),
		RedisAddr:     viper.GetString("REDIS_ADDR"),
		RedisPassword: viper.GetString("REDIS_PASSWORD"),
		RedisDB:       viper.GetInt("REDIS_DB"),
		KafkaBrokers:  viper.GetStringSlice("KAFKA_BROKERS"),
	}

	return config
}

func InitConfig() {
	err := viper.BindEnv("SERVER1_ADDR")
	if err != nil {
		log.Fatalf("Error binding environment variable: %v", err)
	}

	err = viper.BindEnv("SERVER2_ADDR")
	if err != nil {
		log.Fatalf("Error binding environment variable: %v", err)
	}

	err = viper.BindEnv("MYSQL_DSN")
	if err != nil {
		log.Fatalf("Error binding environment variable: %v", err)
	}

	err = viper.BindEnv("REDIS_ADDR")
	if err != nil {
		log.Fatalf("Error binding environment variable: %v", err)
	}

	err = viper.BindEnv("REDIS_PASSWORD")
	if err != nil {
		log.Fatalf("Error binding environment variable: %v", err)
	}

	err = viper.BindEnv("REDIS_DB")
	if err != nil {
		log.Fatalf("Error binding environment variable: %v", err)
	}

	err = viper.BindEnv("KAFKA_BROKERS")
	if err != nil {
		log.Fatalf("Error binding environment variable: %v", err)
	}

	logging.InitLogger()
}
