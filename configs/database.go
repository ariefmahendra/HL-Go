package configs

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dbName = "hackintoshlovers"
var dbHost = "127.0.0.1"
var dbPort = "3306"
var dbUser = "root"
var dbPass = ""

func ConnectDB() *gorm.DB {
	dsn := dbUser + ":" + dbPass + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	return db
}
