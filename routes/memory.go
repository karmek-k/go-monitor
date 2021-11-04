package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/v3/mem"
)

// MemoryRoute returns memory stats
func MemoryRoute(c *gin.Context) {
	mem, err := mem.VirtualMemory()
	if err != nil {
		c.JSON(500, gin.H{
			"error": "Could not read memory stats",
		})
		return
	}

	c.JSON(200, gin.H{
		"error": nil,
		"memory": gin.H{
			"total":       mem.Total,
			"used":        mem.Used,
			"usedPercent": mem.UsedPercent,
		},
	})
}