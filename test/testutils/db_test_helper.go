package testutils

import (
	"fmt"
	"golang-api-server-template/configs"
	"golang-api-server-template/internal/model"
	"log"

	"gorm.io/gorm"
)

// TestDB holds the test database connection
var TestDB *gorm.DB

// SetupTestDatabase sets up the test database connection and runs migrations
func SetupTestDatabase() {
	configs.ConnectTestDB()
	TestDB = configs.DB

	// Drop the test database if it exists
	err := TestDB.Exec("DROP DATABASE IF EXISTS test").Error
	if err != nil {
		log.Fatalf("Failed to drop test database test: %v", err)
	}

	// Create the test database
	err = TestDB.Exec(fmt.Sprintf("CREATE DATABASE test")).Error
	if err != nil {
		log.Fatalf("Failed to create test database test: %v", err)
	} else {
		log.Printf("Database test created successfully")
	}
	TestDB.Exec("USE test")

	// Run migrations (you can add your models here)
	TestDB.AutoMigrate(&model.User{})
}

func ResetTestData() {
	TestDB.Exec("TRUNCATE TABLE users")
}

// TearDownTestDatabase closes the database connection
func TearDownTestDatabase() {
	sqlDB, err := TestDB.DB()
	if err != nil {
		log.Fatalf("Failed to close the database: %v", err)
	}
	sqlDB.Close()
}
