package main

import (
	"fmt"
	"github.com/go-psql-api/model"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"os"
)

func main() {
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

	defer db.Close()

	db.AutoMigrate(&model.Post{})
}
