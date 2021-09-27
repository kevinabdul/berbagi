package config

import (
	//"berbagi/models"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	Db *gorm.DB
)

func InitDb() {
	err1 := godotenv.Load("./.env")
	if err1 != nil {
		log.Fatal("Error loading .env file")
	}

	connectionString := fmt.Sprintf(
		"%s:%s@tcp(%s:%v)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"))

	var err2 error
	Db, err2 = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err2 != nil {
		panic(err2)
	}

	// Db.AutoMigrate(&models.Donator{})
	// Db.AutoMigrate(&models.Volunteer{})
	// Db.AutoMigrate(&models.Yayasan{})
	// Db.AutoMigrate(&models.PersonalRecipient{})
}
