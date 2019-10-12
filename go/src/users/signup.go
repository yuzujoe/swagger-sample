package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ReqData struct {
	Phone    string `form:"phone" json:"email" binding:"required"`
	Password string `json:"password"`
}

func handleSignup(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "hey"})
	// var reqData ReqData

	// if err := c.ShouldBindJSON(&reqData); err != nil {
	// 	c.JSON(http.StatusBadRequest, err)
	// }
}
