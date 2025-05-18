package config

import (
	"fmt"
	"go-mvc-demo/models"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := os.Getenv("MySQL")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to DB:", err)
	}

	err = db.AutoMigrate(&models.Post{}, &models.Comment{})
	if err != nil {
		log.Fatal("Migration failed:", err)
	}

	DB = db
	fmt.Println("Database connected and migrated")
}
