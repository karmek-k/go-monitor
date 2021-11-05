package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/karmek-k/go-monitor/routes"
	"github.com/karmek-k/go-monitor/utils"
)

func main() {
	accounts, err := utils.GetAccounts()
	if err != nil {
		log.Fatalf("An error has occurred while fetching accounts: %v\n", err)
	}

	r := gin.New()

	r.Use(gin.BasicAuth(*accounts))

	r.GET("/memory", routes.MemoryRoute)
	r.GET("/cpu", routes.CpuRoute)

	r.Run(":8000")
}
