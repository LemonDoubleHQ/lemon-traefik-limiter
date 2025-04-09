package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AdminApiController struct{}

type CreateNewApiKeyRequest struct {
	ApiKey string `json:"api_key" binding:"required"`
	TotalCallLimit int64 `json:"total_call_limit" binding:"required"`
	PerMinuteLimit int `json:"per_minute_limit" binding:"required"`
	Paths []string `json:"paths" binding:"required"`
}

func (a *AdminApiController) CreateNewApiKey(c *gin.Context) {
	var request CreateNewApiKeyRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	

	c.JSON(http.StatusOK, gin.H{
		"message": "AdminApi",
	})
}
