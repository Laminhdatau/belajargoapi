package config

import (
	"github.com/Laminhdatau/belajargoapi/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDb()(*gorm.DB, error) {
	dsn := "root:hhhhh@tcp(localhost:3306)/db_golanggin?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
	    return nil, err
	}
      
	// AutoMigrate models
	if err := database.AutoMigrate(&models.Product{}, &models.User{}); err != nil {
	    return nil, err
	}
      
	// Assign the database instance to the global variable
	DB = database
      
	return database, nil
}
