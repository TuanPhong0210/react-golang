// ===========================================
// Package config - Cấu hình ứng dụng
// ===========================================
// File này chứa logic load các biến môi trường
// từ file .env vào ứng dụng
// ===========================================

package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config chứa tất cả cấu hình của ứng dụng
// Được load từ file .env hoặc biến môi trường hệ thống
type Config struct {
	// Database config
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string

	// JWT config
	JWTSecret      string
	JWTExpiryHours string

	// Server config
	ServerPort string
	GinMode    string
}

// AppConfig là instance global của Config
// Được sử dụng trong toàn bộ ứng dụng
var AppConfig *Config

// LoadConfig đọc cấu hình từ file .env
// và trả về struct Config
//
// Input: không có
// Output: *Config - con trỏ đến struct Config
//
// Nếu không tìm thấy file .env, sẽ in warning
// và sử dụng biến môi trường hệ thống
func LoadConfig() *Config {
	// Load file .env (nếu có)
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found, using system environment variables")
	}

	// Tạo config từ biến môi trường
	config := &Config{
		DBHost:         getEnv("DB_HOST", "localhost"),
		DBPort:         getEnv("DB_PORT", "5432"),
		DBUser:         getEnv("DB_USER", "postgres"),
		DBPassword:     getEnv("DB_PASSWORD", "postgres"),
		DBName:         getEnv("DB_NAME", "hrms_db"),
		JWTSecret:      getEnv("JWT_SECRET", "default-secret-key"),
		JWTExpiryHours: getEnv("JWT_EXPIRY_HOURS", "24"),
		ServerPort:     getEnv("SERVER_PORT", "8080"),
		GinMode:        getEnv("GIN_MODE", "debug"),
	}

	// Lưu vào biến global
	AppConfig = config

	return config
}

// getEnv lấy giá trị biến môi trường
// Nếu không tồn tại, trả về giá trị mặc định
//
// Input:
//   - key: tên biến môi trường
//   - defaultValue: giá trị mặc định nếu không tìm thấy
//
// Output: string - giá trị của biến môi trường
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
