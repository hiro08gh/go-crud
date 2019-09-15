package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"

	"github.com/joho/godotenv"
	"log"
	"os"
)

var db = NewDBConn()

func NewDBConn() *gorm.DB {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}
	connectionStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("USER_NAME"),
		os.Getenv("PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	db, err := gorm.Open("postgres", connectionStr)
	if err != nil {
		panic("filaed to conenct database")
	}

	db.LogMode(true)

	return db
}

func GetDBConn() *gorm.DB {
	return db
}
