package users

import (
	"codegen/go/db"
	"codegen/go/models"
	"codegen/go/src/util/crypt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// ReqData リクエストのデータ
type ReqData struct {
	Phone    string `form:"phone" json:"phone" binding:"required"`
	Password string `json:"password"`
}

func handleSignup(c *gin.Context) {
	var reqData ReqData

	// json形式なのかチェック
	if err := c.ShouldBindJSON(&reqData); err != nil {
		c.JSON(http.StatusBadRequest, err)
	}

	// 電話番号が空でないか確認
	if reqData.Phone == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "電話番号を入力してください",
		})
		return
	}

	// パスワードが空でないか確認
	if reqData.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "パスワードを入力してください",
		})
		return
	}

	// userの作成
	if err := createUser(reqData.Phone, reqData.Password); err != "true" {
		log.Fatal(err)
	}
	c.JSON(http.StatusOK, gin.H{"message": "user create!！"})
}

func createUser(phone string, password string) string {
	db := db.Db
	tx := db.Begin()
	hashPwd, _ := crypt.PasswordHash(password)
	user := models.User{Phone: phone, Password: hashPwd, UpdatedAt: time.Now(), CreatedAt: time.Now()}
	if err := tx.Create(&user).Error; err != nil {
		tx.Rollback()
		return "db"
	}
	tx.Commit()
	return "true"
}
