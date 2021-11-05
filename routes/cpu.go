package routes

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/v3/cpu"
)

// CpuRoute returns CPU stats
func CpuRoute(c *gin.Context) {
	usage, err := cpu.Percent(time.Second, true)
	if err != nil {
		c.JSON(500, gin.H{
			"error": "Could not read CPU stats",
		})
		return
	}

	c.JSON(200, gin.H{
		"error": nil,
		"cpu": gin.H{
			"usage": usage,
		},
	})
}