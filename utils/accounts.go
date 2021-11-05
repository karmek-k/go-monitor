package utils

import (
	"errors"
	"os"

	"github.com/gin-gonic/gin"
)

// GetAccounts fetches credentials from GO_MONITOR_LOGIN and GO_MONITOR_PASSWORD
// environment variables and returns them as `gin.Accounts`.
func GetAccounts() (*gin.Accounts, error) {
	login := os.Getenv("GO_MONITOR_LOGIN")
	pass := os.Getenv("GO_MONITOR_PASSWORD")
	if login == "" || pass == "" {
		return nil, errors.New("please set GO_MONITOR_LOGIN and GO_MONITOR_PASSWORD environment variables")
	}

	return &gin.Accounts{
		login: pass,
	}, nil
}