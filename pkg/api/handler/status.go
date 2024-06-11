package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *handler) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
