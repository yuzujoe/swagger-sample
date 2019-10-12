package users

import (
	"github.com/gin-gonic/gin"
)

// AddRoute usersRouteの API設定
func AddRoute(r *gin.RouterGroup) {
	g := r.Group("/users")
	{
		g.GET("/signup", handleSignup)
	}
}
