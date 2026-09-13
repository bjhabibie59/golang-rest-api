package config

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"go-rest-api/internal/model"
)

func ConnectDB() *gorm.DB {
	host := os.Getenv("DB_HOST")
	if host == "" {
		host = "localhost"
	}
	port := os.Getenv("DB_PORT")
	if port == "" {
		port = "3306"
	}
	user := os.Getenv("DB_USER")
	if user == "" {
		user = "root"
	}
	pass := os.Getenv("DB_PASS")
	name := os.Getenv("DB_NAME")
	if name == "" {
		name = "golang_rest_api"
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pass, host, port, name)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Auto Migrate the user model
	db.AutoMigrate(&model.User{})

	// Seed dummy data if empty
	var count int64
	db.Model(&model.User{}).Count(&count)
	if count == 0 {
		db.Create(&model.User{Name: "Sudais", Email: "sudais@yahoo.com", Password: "password123"})
		db.Create(&model.User{Name: "Novan", Email: "novan@yahoo.com", Password: "password123"})
		log.Println("Dummy data seeded successfully!")
	}

	return db
}
