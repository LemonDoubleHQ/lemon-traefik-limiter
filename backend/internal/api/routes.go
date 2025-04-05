package api

import (
	"lemon-traefik-limiter-backend/internal/api/admin"
	"lemon-traefik-limiter-backend/internal/api/forwardauth"
	"lemon-traefik-limiter-backend/internal/api/health"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine) {
	healthGroup := r.Group("/health")
	healthController := new(health.HealthController)
	{
		healthGroup.GET("", healthController.HealthCheck)
	}

	forwardAuthGroup := r.Group("/api/forward-auth")
	forwardAuthController := new(forwardauth.ForwardAuthController)
	{
		forwardAuthGroup.GET("", forwardAuthController.ForwardAuth)
	}

	adminApiGroup := r.Group("/api/admin")
	adminApiController := new(admin.AdminApiController)
	{
		adminApiGroup.GET("", adminApiController.AdminApi)
	}
	
}