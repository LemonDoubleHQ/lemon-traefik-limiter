package forwardauth

import (
	"lemon-traefik-limiter-backend/internal/config"
	"lemon-traefik-limiter-backend/internal/db/redis"
	"lemon-traefik-limiter-backend/internal/httperror"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

type ForwardAuthController struct{}

func (f *ForwardAuthController) ForwardAuth(c *gin.Context) {

	ip := c.GetHeader("X-Real-Ip")

	if ip == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}

	redisClient := redis.GetRedisClient()
	redisService := redis.GetRedisService(redisClient)

	var wg sync.WaitGroup
	wg.Add(1)

	var ipError *httperror.HttpError

	go func() {
		defer wg.Done()
		ipCallCount, err := redisService.GetIpCallCountPerSecond(ip)
		if err != nil {
			ipError = &httperror.HttpError{
				StatusCode: http.StatusInternalServerError,
				Message: err.Error(),
			}
			return
		}

		if ipCallCount >= config.GlobalConfig.RateLimit.MaxRequestInSecondPerIp {
			ipError = &httperror.HttpError{
				StatusCode: http.StatusTooManyRequests,
				Message: "Rate Limit Exceeded",
			}
		}
	}()

	wg.Wait()
	
	if ipError != nil {
		c.JSON(ipError.StatusCode, gin.H{
			"message": ipError.Message,
		})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{
		"result": "ok",
	})
}
