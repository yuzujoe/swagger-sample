package route

import (
	"codegen/go/src/users"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func Route() {
	r := gin.Default()

	r.GET("/alive", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "OK")
	})

	rg := r.Group("/api")
	users.AddRoute(rg)

	r.Run(os.Getenv("PORT"))
}
