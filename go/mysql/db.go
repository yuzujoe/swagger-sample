package mysql

import (
	"codegen/go/models"
	"fmt"
	"os"
	"strings"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// DB for gormDB
var DB *gorm.DB

var err error

// Init gormでmysqlへの接続データベースの作成
func Init() {

	ConnectDB()
	defer CloseDB()

	DB.LogMode(true)

	err = DB.DB().Ping()
	if err != nil {
		panic(err)
	}
}

// ConnectDB mys qlへの接続
func ConnectDB() *gorm.DB {
	path := strings.Join([]string{os.Getenv("MYSQL_USER"), ":", os.Getenv("MYSQL_PASSWORD"), "@tcp(", os.Getenv("DB_HOST"), ":", os.Getenv("DB_PORT"), ")/", os.Getenv("MYSQL_DATABASE"), "?charset=utf8&parseTime=True&loc=Local"}, "")

	DB, err = gorm.Open("mysql", path)
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected!")

	return DB
}

// CloseDB dbへの接続を閉じる
func CloseDB() {
	DB.Close()
}

func migrate() {
	DB.AutoMigrate(&models.Reservation{}, &models.User{})
}
