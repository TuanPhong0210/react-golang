// ===========================================
// HR Management System - Main Entry Point
// ===========================================
// File này là điểm khởi động của ứng dụng backend
// Khởi tạo config, database, routes và start server
// ===========================================

package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/hanbiro/hrms-server/config"
	"github.com/hanbiro/hrms-server/database"
	"github.com/hanbiro/hrms-server/routes"
)

// main là entry point của ứng dụng
// Thực hiện các bước:
// 1. Load cấu hình từ .env
// 2. Kết nối database
// 3. Khởi tạo Gin router
// 4. Cấu hình routes
// 5. Start HTTP server
func main() {
	log.Println("🚀 Starting HR Management System Server...")

	// =========================================
	// Step 1: Load Configuration
	// =========================================
	log.Println("📝 Loading configuration...")
	cfg := config.LoadConfig()

	// =========================================
	// Step 2: Connect to Database
	// =========================================
	log.Println("🔌 Connecting to database...")
	db, err := database.ConnectDB(cfg)
	if err != nil {
		log.Fatalf("❌ Failed to connect to database: %v", err)
	}

	// Lấy underlying sql.DB để đóng connection khi cần
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("❌ Failed to get database instance: %v", err)
	}
	defer sqlDB.Close()

	// =========================================
	// Step 3: Initialize Gin Router
	// =========================================
	log.Println("🌐 Initializing HTTP router...")

	// Set Gin mode từ config
	gin.SetMode(cfg.GinMode)

	// Tạo router với default middleware (Logger, Recovery)
	router := gin.Default()

	// =========================================
	// Step 4: Setup Routes
	// =========================================
	log.Println("🛤️  Setting up routes...")
	routes.SetupRoutes(router, db)

	// =========================================
	// Step 5: Start Server
	// =========================================
	serverAddr := ":" + cfg.ServerPort
	log.Printf("✅ Server is running on http://localhost%s", serverAddr)
	log.Println("📚 API Documentation:")
	log.Println("   - POST   /api/auth/login      - Đăng nhập")
	log.Println("   - GET    /api/auth/me         - Thông tin user")
	log.Println("   - GET    /api/users           - Danh sách users")
	log.Println("   - POST   /api/attendance/check-in  - Check in")
	log.Println("   - GET    /api/approvals       - Danh sách đơn")
	log.Println("   - GET    /health              - Health check")
	log.Println("")
	log.Println("🔐 Default Admin Account:")
	log.Println("   - Email: admin@company.com")
	log.Println("   - Password: admin123")
	log.Println("")

	if err := router.Run(serverAddr); err != nil {
		log.Fatalf("❌ Failed to start server: %v", err)
	}
}
