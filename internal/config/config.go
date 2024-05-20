// internal/config/config.go
package config

import (
	"go-web-scaffold/internal/logging"
	"log"
	"time"

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

	RateLimitConfig
}

type RateLimitConfig struct {
	WindowsSize  time.Duration
	BucketCount  int
	RequestLimit int
}

func LoadConfig() *Config {

	viper.AutomaticEnv()

	config := &Config{
		Server1Addr:   viper.GetString("SERVER1_ADDR"),
		Server2Addr:   viper.GetString("SERVER2_ADDR"),
		MySQLDSN:      viper.GetString("MYSQL_DSN"),
		RedisAddr:     viper.GetString("REDIS_ADDR"),
		RedisPassword: viper.GetString("REDIS_PASSWORD"),
		RedisDB:       viper.GetInt("REDIS_DB"),
		KafkaBrokers:  viper.GetStringSlice("KAFKA_BROKERS"),
	}
	config.WindowsSize = viper.GetDuration("WINDOWS_SIZE")
	config.BucketCount = viper.GetInt("BUCKET_COUNT")
	config.RequestLimit = viper.GetInt("REQUEST_LIMIT")

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
