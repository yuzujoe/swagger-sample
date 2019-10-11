package route

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func Route() {
	r := gin.Default()

	r.GET("/alive", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "OK")
	})
	r.Run(os.Getenv("PORT"))
}
