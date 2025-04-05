package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AdminApiController struct{}

func (a *AdminApiController) AdminApi(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "AdminApi",
	})
}
