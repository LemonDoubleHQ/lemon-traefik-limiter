package health

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthController struct {}

func (h HealthController) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}