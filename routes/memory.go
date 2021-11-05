package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/karmek-k/go-monitor/resources"
)

// MemoryRoute returns memory stats
func MemoryRoute(c *gin.Context) {
	stats, err := resources.GetMemoryStats()
	if err != nil {
		c.JSON(500, gin.H{
			"error": "Could not read memory stats",
		})
		return
	}

	c.JSON(200, gin.H{
		"error": nil,
		"memory": stats,
	})
}