package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *AuthHandler) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
