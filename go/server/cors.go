package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Cors CORS設定
func corsSet() gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	return cors.New(config)
}
