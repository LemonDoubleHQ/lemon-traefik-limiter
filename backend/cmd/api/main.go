package main

import (
	"lemon-traefik-limiter-backend/internal/api"
	"lemon-traefik-limiter-backend/internal/config"

	"github.com/gin-gonic/gin"
)

func main(){
	config.LoadEnvConfig()
	
	r := gin.Default()
	api.InitRoutes(r)
	r.Run()
}