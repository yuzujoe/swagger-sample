package db

import (
	"codegen/go/models"
	"fmt"
	"os"
	"strings"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Db for gormDB
var Db *gorm.DB

// Init gormでmysqlへの接続データベースの作成
func Init() *gorm.DB {
	var err error

	path := strings.Join([]string{os.Getenv("MYSQL_USER"), ":", os.Getenv("MYSQL_PASSWORD"), "@tcp(", os.Getenv("DB_HOST"), ":", os.Getenv("DB_PORT"), ")/", os.Getenv("MYSQL_DATABASE"), "?charset=utf8&parseTime=True&loc=Local"}, "")

	Db, err = gorm.Open("mysql", path)
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected!")
	Db.LogMode(true)

	err = Db.DB().Ping()
	if err != nil {
		panic(err)
	}

	Db.AutoMigrate(&models.Reservation{}, &models.User{})

	return Db
}
