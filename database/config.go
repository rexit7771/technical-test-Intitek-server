package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectDB() {
	// Check both naming conventions for flexibility
	dbUser := getEnvWithFallback("MYSQLUSER", "DB_USER")
	dbPass := getEnvWithFallback("MYSQLPASSWORD", "DB_PASSWORD")
	dbHost := getEnvWithFallback("MYSQLHOST", "DB_HOST")
	dbPort := getEnvWithFallback("MYSQLPORT", "DB_PORT")
	dbName := getEnvWithFallback("MYSQLDATABASE", "DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPass, dbHost, dbPort, dbName)

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database: ", err)
	}
	fmt.Println("Database Connected")
}

// Helper function to check multiple environment variable names
func getEnvWithFallback(key, fallbackKey string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return os.Getenv(fallbackKey)
}
