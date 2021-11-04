package main

import (
	"github.com/gin-gonic/gin"

	"github.com/karmek-k/go-monitor/routes"
	"github.com/karmek-k/go-monitor/utils"
)

func main() {
	r := gin.New()

	r.Use(gin.BasicAuth(utils.GetAccounts()))

	r.GET("/", routes.IndexRoute)

	r.Run(":8000")
}
