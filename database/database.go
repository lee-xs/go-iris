package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var (
	DB = New()
)

func New() *gorm.DB{
	db, err := gorm.Open("mysql", "root:123456@/iris?charset=utf8&parseTime=true&loc=Local")

	if err != nil {
		log.Println("数据库连接出现错误", err)
	}

	return db
}