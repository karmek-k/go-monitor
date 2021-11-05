package routes

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/karmek-k/go-monitor/resources"
)

// CpuRoute returns CPU stats
func CpuRoute(c *gin.Context) {
	stats, err := resources.GetCpuStats(time.Second, true)
	if err != nil {
		c.JSON(500, gin.H{
			"error": "Could not read CPU stats",
		})
		return
	}

	c.JSON(200, gin.H{
		"error": nil,
		"cpu": *stats,
	})
}