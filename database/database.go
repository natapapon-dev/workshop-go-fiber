package database

import (
	"fmt"
	"log"
	"os"
	"workshop/go-fiber/models"

	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DBConn *gorm.DB
)

func InitDatabase() {
	log.Printf("START INITIALS DATABASE")
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Local",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_END_POINT"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"))

	log.Printf("DATABASE CONNECT AT %s", dsn)

	var err error
	DBConn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	log.Print("DATABASE CONNECTION SUCCESSFULLY")
	if err != nil {
		log.Fatal(err.Error())
	}

	MigrateTable()
}

func MigrateTable() {
	DBConn.AutoMigrate(&models.Employees{})
}
