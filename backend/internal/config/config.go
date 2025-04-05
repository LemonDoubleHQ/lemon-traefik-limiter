package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	GlobalConfig *Config
)

type Config struct {
	Environment string
	Redis       RedisConfig
	RateLimit   RateLimitConfig
}

type RedisConfig struct {
	Addr     string
	Password string
	DB       int
}

type RateLimitConfig struct {
	MaxRequestInSecondPerIp int64
}

func LoadEnvConfig() error {
	_ = godotenv.Load(".env")
	GlobalConfig = &Config{
		Environment: getStringEnv("ENVIRONMENT", true),
		Redis: RedisConfig{
			Addr: getStringEnv("REDIS_ADDR", true),
			Password: getStringEnv("REDIS_PASSWORD", true),
			DB : getIntEnv("REDIS_DB", true),
		},
		RateLimit: RateLimitConfig{
			MaxRequestInSecondPerIp: int64(getIntEnv("MAX_REQUEST_IN_SECOND_PER_IP", true)),
		},
	}
	return nil
}


func getStringEnv(key string, mustExist bool) string {
	value, exist := os.LookupEnv(key)
	if mustExist && !exist {
		log.Fatalf("%s environment variable is required", key)
	}
	return value
}

func getIntEnv(key string, mustExist bool) int {
	value, exist := os.LookupEnv(key)
	if mustExist && !exist {
		log.Fatalf("%s environment variable is required", key)
	}

	intValue, err := strconv.Atoi(value)
	if err != nil {
		log.Fatalf("%s environment variable is not an integer: %v", key, err)
	}
	return intValue
}