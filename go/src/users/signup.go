package users

import (
	"codegen/go/models"
	"codegen/go/mysql"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// ReqData リクエストのデータ
type ReqData struct {
	Phone    string `json:"phone" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func HandleSignup(c *gin.Context) {
	var reqData ReqData

	// エラーハンドリング
	if err := c.Bind(&reqData); err != nil {
		if reqData.Phone == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "電話番号を入力してください",
			})
			return
		} else if reqData.Password == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "パスワードを入力してください",
			})
			return
		}
		return
	}

	// userの作成
	if err := createUser(reqData.Phone, reqData.Password); err != "true" {
		log.Fatal(err)
	}
	c.JSON(http.StatusOK, gin.H{"message": "user create!！"})
}

func createUser(phone string, password string) string {
	mysql.ConnectDB()
	defer mysql.CloseDB()
	db := mysql.DB
	tx := db.Begin()
	hashPwd, _ := passwordHash(password)
	user := models.User{Phone: phone, Password: hashPwd, UpdatedAt: time.Now(), CreatedAt: time.Now()}
	if err := tx.Create(&user).Error; err != nil {
		tx.Rollback()
		return "db"
	}
	tx.Commit()
	return "true"
}
