package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/karmek-k/go-monitor/routes"
	"github.com/karmek-k/go-monitor/utils"
)

func main() {
	accounts, err := utils.GetAccounts()
	if err != nil {
		log.Fatalf("An error has occurred while fetching accounts: %v\n", err)
	}

	gin.SetMode(gin.ReleaseMode)

	r := gin.New()

	r.Use(gin.BasicAuth(*accounts))

	r.GET("/memory", routes.MemoryRoute)
	r.GET("/cpu", routes.CpuRoute)

	port := "8000"
	if envPort := os.Getenv("PORT"); envPort != "" {
		port = envPort
	}

	log.Printf("Running on port %v\n", port)
	r.Run(":" + port)
}
