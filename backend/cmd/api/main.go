package main

import (
	"lemon-traefik-limiter-backend/internal/api"
	"lemon-traefik-limiter-backend/internal/config"
	"lemon-traefik-limiter-backend/internal/db/postgres"
	"lemon-traefik-limiter-backend/internal/db/redis"

	"github.com/gin-gonic/gin"
)

func main(){
	config.LoadEnvConfig()
	redis.InitRedis()
	postgres.InitPostgres()

	if config.GlobalConfig.Environment == "local" {
		postgres.RunDbMigration()
	}
	
	r := gin.Default()
	api.InitRoutes(r)
	r.Run()
}