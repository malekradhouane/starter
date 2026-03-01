package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetStatus service status
// @Summary Service status
// @Description Health check endpoint returning service status
// @Tags status
// @Produce json
// @Success 200 {object} map[string]string
// @Router /status [get]
func GetStatus(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
