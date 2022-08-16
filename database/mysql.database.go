package database

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DatabaseMysql() {
	var err error
	const MYSQL = "root:@tcp(localhost)/golang-fiber-lv1?charset=utf8&parseTime=True&loc=Local"
	dsn := MYSQL
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panicln("failed to connect database")
	}
	log.Println("database connection successful")
}
