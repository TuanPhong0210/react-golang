// ===========================================
// Package middlewares - Permission Middleware
// ===========================================
// File này chứa middleware phân quyền theo role
// Kiểm tra user có đủ quyền truy cập resource không
// ===========================================

package middlewares

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hanbiro/hrms-server/utils"
)

// RequireAdmin yêu cầu user phải có role Admin
//
// Input: không có
// Output: gin.HandlerFunc
//
// Usage:
//
//	router.POST("/users", middlewares.RequireAdmin(), userController.CreateUser)
func RequireAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		roleName, exists := c.Get("roleName")
		if !exists {
			utils.ForbiddenResponse(c, "Không có quyền truy cập")
			c.Abort()
			return
		}

		if roleName.(string) != "Admin" {
			utils.ForbiddenResponse(c, "Chỉ Admin mới có quyền thực hiện thao tác này")
			c.Abort()
			return
		}

		c.Next()
	}
}

// RequireManager yêu cầu user phải có role Admin hoặc Manager
//
// Input: không có
// Output: gin.HandlerFunc
//
// Usage:
//
//	router.GET("/users", middlewares.RequireManager(), userController.GetAllUsers)
func RequireManager() gin.HandlerFunc {
	return func(c *gin.Context) {
		roleName, exists := c.Get("roleName")
		if !exists {
			utils.ForbiddenResponse(c, "Không có quyền truy cập")
			c.Abort()
			return
		}

		role := roleName.(string)
		if role != "Admin" && role != "Manager" {
			utils.ForbiddenResponse(c, "Chỉ Admin hoặc Manager mới có quyền thực hiện thao tác này")
			c.Abort()
			return
		}

		c.Next()
	}
}

// RequireRoles yêu cầu user phải có một trong các roles được chỉ định
//
// Input:
//   - roles: Danh sách roles được phép truy cập
//
// Output: gin.HandlerFunc
//
// Usage:
//
//	router.GET("/data", middlewares.RequireRoles("Admin", "Manager"), handler)
func RequireRoles(roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		roleName, exists := c.Get("roleName")
		if !exists {
			utils.ForbiddenResponse(c, "Không có quyền truy cập")
			c.Abort()
			return
		}

		userRole := roleName.(string)
		for _, role := range roles {
			if userRole == role {
				c.Next()
				return
			}
		}

		utils.ForbiddenResponse(c, "Bạn không có quyền thực hiện thao tác này")
		c.Abort()
	}
}

// RequireSameUserOrAdmin yêu cầu user phải là chính mình hoặc Admin
// Dùng cho các endpoint mà user chỉ được thao tác với data của mình
//
// Input:
//   - userIDParam: Tên param chứa userID trong URL (vd: "id")
//
// Output: gin.HandlerFunc
//
// Usage:
//
//	router.GET("/users/:id", middlewares.RequireSameUserOrAdmin("id"), handler)
func RequireSameUserOrAdmin(userIDParam string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Lấy userID từ context (người đang đăng nhập)
		currentUserID, exists := c.Get("userID")
		if !exists {
			utils.ForbiddenResponse(c, "Không có quyền truy cập")
			c.Abort()
			return
		}

		// Lấy role
		roleName, _ := c.Get("roleName")

		// Admin được quyền truy cập tất cả
		if roleName.(string) == "Admin" {
			c.Next()
			return
		}

		// Lấy userID từ URL param
		targetUserID := c.Param(userIDParam)

		// So sánh (convert để so sánh)
		if targetUserID != "" {
			// Kiểm tra có phải đang truy cập resource của chính mình
			targetUID, err := strconv.ParseUint(targetUserID, 10, 32)
			if err == nil && uint(targetUID) == currentUserID.(uint) {
				c.Next()
				return
			}
		} else {
			// Nếu không có param (vd route không có ID), cho phép truy cập
			// (Giả sử các routes này đã được bảo vệ bởi middleware khác hoặc không cần check ID)
			c.Next()
			return
		}

		utils.ForbiddenResponse(c, "Bạn không có quyền thực hiện thao tác này")
		c.Abort()
	}
}
