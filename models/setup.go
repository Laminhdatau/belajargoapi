package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDb() {
	database, err := gorm.Open(mysql.Open("root:hhhhh@tcp(localhost:3306)/db_golanggin"))
	if err != nil {
		panic(err)
	}

	database.AutoMigrate(&Product{})
	DB = database
}
