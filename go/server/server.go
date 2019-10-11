package server

import (
	"codegen/go/db"
	"codegen/go/middleware"
	"codegen/go/route"

	"github.com/gin-gonic/gin"
)

// StartServer サーバーの起動
func StartServer() {

	db.Init()

	createServer()
}

func createServer() (r *gin.Engine) {
	r = gin.New()
	gin.SetMode(gin.ReleaseMode)

	r.Use(middleware.Cors())

	route.Route()
	return
}
