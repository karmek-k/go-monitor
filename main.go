package main

import (
	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/v3/mem"
)

func main() {
	r := gin.New()

	r.Use(gin.BasicAuth(gin.Accounts{
		"bartek": "12345",
	}))

	// TODO: caching
	r.GET("/", func(c *gin.Context) {
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
				"total": mem.Total,
				"used": mem.Used,
				"usedPercent": mem.UsedPercent,
			},
		})
	})

	r.Run(":8000")
}