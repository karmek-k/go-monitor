package utils

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

// GetAccount fetches credentials from GO_MONITOR_LOGIN and GO_MONITOR_PASSWORD
// environment variables and returns them as `gin.Accounts`.
func GetAccount() gin.Accounts {
	login := os.Getenv("GO_MONITOR_LOGIN")
	if login == "" {
		log.Fatalln("No login found! Please set the GO_MONITOR_LOGIN environment variable.")
	}

	pass := os.Getenv("GO_MONITOR_PASSWORD")
	if pass == "" {
		log.Fatalln("No password found! Please set the GO_MONITOR_PASSWORD environment variable.")
	}

	return gin.Accounts{
		login: pass,
	}
}