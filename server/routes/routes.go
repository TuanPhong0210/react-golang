// ===========================================
// Package routes - API Routes
// ===========================================
// File này định nghĩa tất cả các API routes
// Sử dụng Gin Router để group và protect routes
// ===========================================

package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/hanbiro/hrms-server/controllers"
	"github.com/hanbiro/hrms-server/middlewares"
	"github.com/hanbiro/hrms-server/repositories"
	"github.com/hanbiro/hrms-server/services"
	"gorm.io/gorm"
)

// SetupRoutes cấu hình tất cả routes cho ứng dụng
//
// Input:
//   - router: Gin Engine instance
//   - db: GORM database instance
//
// Output: không có (cấu hình trực tiếp lên router)
//
// Routes structure:
//   - /api/auth: Authentication routes (public)
//   - /api/users: User management routes (protected)
//   - /api/attendance: Attendance routes (protected)
//   - /api/approvals: Approval routes (protected)
func SetupRoutes(router *gin.Engine, db *gorm.DB) {
	// =========================================
	// CORS Configuration
	// =========================================
	// Cho phép frontend gọi API từ domain khác
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// =========================================
	// Initialize Repositories
	// =========================================
	userRepo := repositories.NewUserRepository(db)
	attendanceRepo := repositories.NewAttendanceRepository(db)
	approvalRepo := repositories.NewApprovalRepository(db)

	// =========================================
	// Initialize Services
	// =========================================
	authService := services.NewAuthService(userRepo)
	userService := services.NewUserService(userRepo)
	attendanceService := services.NewAttendanceService(attendanceRepo)
	approvalService := services.NewApprovalService(approvalRepo)

	// =========================================
	// Initialize Controllers
	// =========================================
	authController := controllers.NewAuthController(authService)
	userController := controllers.NewUserController(userService)
	attendanceController := controllers.NewAttendanceController(attendanceService)
	approvalController := controllers.NewApprovalController(approvalService)

	// =========================================
	// API Routes
	// =========================================
	api := router.Group("/api")
	{
		// -----------------------------------------
		// Auth Routes - Public (không cần token)
		// -----------------------------------------
		auth := api.Group("/auth")
		{
			// POST /api/auth/login - Đăng nhập
			auth.POST("/login", authController.Login)

			// POST /api/auth/logout - Đăng xuất (cần token)
			auth.POST("/logout", middlewares.AuthMiddleware(), authController.Logout)

			// GET /api/auth/me - Lấy thông tin user hiện tại
			auth.GET("/me", middlewares.AuthMiddleware(), authController.GetCurrentUser)
		}

		// -----------------------------------------
		// User Routes - Protected
		// -----------------------------------------
		users := api.Group("/users")
		users.Use(middlewares.AuthMiddleware()) // Yêu cầu đăng nhập
		{
			// GET /api/users - Lấy danh sách users (Admin, Manager)
			users.GET("", middlewares.RequireManager(), userController.GetAllUsers)

			// GET /api/users/:id - Lấy thông tin user theo ID
			users.GET("/:id", userController.GetUserByID)

			// POST /api/users - Tạo user mới (Admin only)
			users.POST("", middlewares.RequireAdmin(), userController.CreateUser)

			// PUT /api/users/:id - Cập nhật user (Admin, Manager)
			users.PUT("/:id", middlewares.RequireManager(), userController.UpdateUser)

			// DELETE /api/users/:id - Xóa user (Admin only)
			users.DELETE("/:id", middlewares.RequireAdmin(), userController.DeleteUser)

			// PUT /api/users/change-password - Đổi mật khẩu
			users.PUT("/change-password", userController.ChangePassword)
		}

		// -----------------------------------------
		// Attendance Routes - Protected
		// -----------------------------------------
		attendance := api.Group("/attendance")
		attendance.Use(middlewares.AuthMiddleware())
		{
			// POST /api/attendance/check-in - Check in
			attendance.POST("/check-in", attendanceController.CheckIn)

			// POST /api/attendance/check-out - Check out
			attendance.POST("/check-out", attendanceController.CheckOut)

			// GET /api/attendance/today - Lấy attendance hôm nay
			attendance.GET("/today", attendanceController.GetTodayAttendance)

			// GET /api/attendance/history - Lấy lịch sử của mình
			attendance.GET("/history", attendanceController.GetMyAttendanceHistory)

			// GET /api/attendance - Lấy tất cả (Admin, Manager)
			attendance.GET("", middlewares.RequireManager(), attendanceController.GetAllAttendances)
		}

		// -----------------------------------------
		// Approval Routes - Protected
		// -----------------------------------------
		approvals := api.Group("/approvals")
		approvals.Use(middlewares.AuthMiddleware())
		{
			// POST /api/approvals - Tạo đơn mới
			approvals.POST("", approvalController.CreateApproval)

			// GET /api/approvals/my - Lấy danh sách đơn của mình
			approvals.GET("/my", approvalController.GetMyApprovals)

			// GET /api/approvals/pending - Lấy đơn chờ duyệt (Admin, Manager)
			approvals.GET("/pending", middlewares.RequireManager(), approvalController.GetPendingApprovals)

			// GET /api/approvals/:id - Lấy chi tiết đơn
			approvals.GET("/:id", approvalController.GetApprovalByID)

			// GET /api/approvals - Lấy tất cả đơn (Admin, Manager)
			approvals.GET("", middlewares.RequireManager(), approvalController.GetAllApprovals)

			// PUT /api/approvals/:id/approve - Duyệt đơn (Admin, Manager)
			approvals.PUT("/:id/approve", middlewares.RequireManager(), approvalController.ApproveRequest)

			// PUT /api/approvals/:id/reject - Từ chối đơn (Admin, Manager)
			approvals.PUT("/:id/reject", middlewares.RequireManager(), approvalController.RejectRequest)
		}
	}

	// =========================================
	// Health Check Route
	// =========================================
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "ok",
			"message": "HRMS Server is running",
		})
	})
}
