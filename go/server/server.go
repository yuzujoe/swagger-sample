package server

import (
	"codegen/go/mysql"
	"codegen/go/route"

	"github.com/gin-gonic/gin"
)

// StartServer サーバーの起動
func StartServer() {

	mysql.Init()

	createServer()
}

func createServer() (r *gin.Engine) {
	r = gin.New()
	gin.SetMode(gin.ReleaseMode)

	r.Use(corsSet())

	route.Route()
	return
}
