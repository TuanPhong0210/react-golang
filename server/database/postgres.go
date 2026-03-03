// ===========================================
// Package database - Kết nối PostgreSQL
// ===========================================
// File này chứa logic kết nối database
// sử dụng GORM ORM
// ===========================================

package database

import (
	"fmt"
	"log"

	"github.com/hanbiro/hrms-server/config"
	"github.com/hanbiro/hrms-server/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// DB là instance global của GORM database
// Được sử dụng trong toàn bộ ứng dụng
var DB *gorm.DB

// ConnectDB thiết lập kết nối đến PostgreSQL
// và thực hiện auto migrate các models
//
// Input: cfg *config.Config - cấu hình database
// Output: *gorm.DB - instance database, error nếu có lỗi
//
// Function này sẽ:
// 1. Tạo connection string từ config
// 2. Kết nối đến PostgreSQL
// 3. Auto migrate các models
// 4. Seed dữ liệu mặc định nếu cần
func ConnectDB(cfg *config.Config) (*gorm.DB, error) {
	// Tạo connection string (DSN - Data Source Name)
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
	)

	// Mở kết nối đến database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // Log SQL queries
	})

	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	log.Println("✅ Connected to PostgreSQL database successfully!")

	// Lưu vào biến global
	DB = db

	// Auto migrate các models
	// GORM sẽ tự động tạo/cập nhật bảng dựa trên struct
	err = autoMigrate(db)
	if err != nil {
		return nil, fmt.Errorf("failed to auto migrate: %v", err)
	}

	// Seed dữ liệu mặc định
	err = seedDefaultData(db)
	if err != nil {
		log.Printf("Warning: failed to seed default data: %v", err)
	}

	return db, nil
}

// autoMigrate tự động tạo/cập nhật các bảng
// dựa trên định nghĩa trong models
//
// Input: db *gorm.DB - instance database
// Output: error nếu có lỗi
func autoMigrate(db *gorm.DB) error {
	log.Println("🔄 Running auto migration...")

	err := db.AutoMigrate(
		&models.Role{},
		&models.User{},
		&models.Attendance{},
		&models.Approval{},
	)

	if err != nil {
		return err
	}

	log.Println("✅ Auto migration completed!")
	return nil
}

// seedDefaultData tạo dữ liệu mặc định
// bao gồm roles và admin user
//
// Input: db *gorm.DB - instance database
// Output: error nếu có lỗi
func seedDefaultData(db *gorm.DB) error {
	// Kiểm tra xem đã có roles chưa
	var roleCount int64
	db.Model(&models.Role{}).Count(&roleCount)

	if roleCount == 0 {
		log.Println("🌱 Seeding default roles...")

		roles := []models.Role{
			{Name: "Admin", Description: "Quản trị viên hệ thống - Full quyền"},
			{Name: "Manager", Description: "Quản lý - Duyệt đơn, quản lý nhân viên"},
			{Name: "Employee", Description: "Nhân viên - Chấm công, tạo đơn"},
		}

		for _, role := range roles {
			db.Create(&role)
		}

		log.Println("✅ Default roles created!")
	}

	// Kiểm tra xem đã có admin user chưa
	var userCount int64
	db.Model(&models.User{}).Count(&userCount)

	if userCount == 0 {
		log.Println("🌱 Seeding default admin user...")

		// Lấy role Admin
		var adminRole models.Role
		db.Where("name = ?", "Admin").First(&adminRole)

		// Tạo admin user với password đã hash
		// Password: admin123
		adminUser := models.User{
			Email:      "admin@company.com",
			Password:   "$2a$10$AqEsVneT/A.vVyWtU5GS8eVXbSRmjrZuW6WZ7jn7O0.7tummgRNQO",
			FullName:   "System Admin",
			Phone:      "0901234567",
			Department: "IT",
			Position:   "System Administrator",
			RoleID:     adminRole.ID,
			IsActive:   true,
		}

		db.Create(&adminUser)

		log.Println("✅ Default admin user created!")
		log.Println("   Email: admin@company.com")
		log.Println("   Password: admin123")
	}

	return nil
}
