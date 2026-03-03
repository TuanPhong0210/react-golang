// ===========================================
// Package middlewares - Auth Middleware
// ===========================================
// File này chứa middleware xác thực JWT token
// Được sử dụng để bảo vệ các routes cần đăng nhập
// ===========================================

package middlewares

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/hanbiro/hrms-server/utils"
)

// AuthMiddleware kiểm tra JWT token trong header
// Nếu token hợp lệ, lưu thông tin user vào context
//
// Input: không có
// Output: gin.HandlerFunc
//
// Usage:
//
//	router.Use(middlewares.AuthMiddleware())
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Lấy Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			utils.UnauthorizedResponse(c, "Vui lòng đăng nhập")
			c.Abort()
			return
		}

		// Kiểm tra format: Bearer <token>
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			utils.UnauthorizedResponse(c, "Token không hợp lệ")
			c.Abort()
			return
		}

		tokenString := parts[1]

		// Parse và validate token
		claims, err := utils.ParseToken(tokenString)
		if err != nil {
			utils.UnauthorizedResponse(c, "Token không hợp lệ hoặc đã hết hạn")
			c.Abort()
			return
		}

		// Lưu thông tin user vào context để sử dụng ở các handler sau
		c.Set("userID", claims.UserID)
		c.Set("email", claims.Email)
		c.Set("roleID", claims.RoleID)
		c.Set("roleName", claims.RoleName)

		// Tiếp tục xử lý request
		c.Next()
	}
}

// OptionalAuthMiddleware kiểm tra token nếu có
// Không bắt buộc phải có token
//
// Input: không có
// Output: gin.HandlerFunc
func OptionalAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.Next()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.Next()
			return
		}

		tokenString := parts[1]
		claims, err := utils.ParseToken(tokenString)
		if err != nil {
			c.Next()
			return
		}

		// Lưu thông tin user nếu token hợp lệ
		c.Set("userID", claims.UserID)
		c.Set("email", claims.Email)
		c.Set("roleID", claims.RoleID)
		c.Set("roleName", claims.RoleName)

		c.Next()
	}
}
