package route

import (
	"codegen/go/src/users"

	"github.com/gin-gonic/gin"
)

// AddRoute usersRouteの API設定
func userRoute(r *gin.RouterGroup) {
	g := r.Group("/users")
	{
		g.POST("/signup", users.HandleSignup)
	}
}
