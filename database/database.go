package database

import (
	"fiber-cloudinary/model"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() {
	var err error
	DB, err = gorm.Open(mysql.Open("root:@tcp(127.0.0.1:3306)/fiber_cloudinary?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic("cant connect to db")
	}
	fmt.Println("connected to db")
}

func AutoMigration() {
	DB.AutoMigrate(&model.User{})

	log.Println("table migrated")
}
