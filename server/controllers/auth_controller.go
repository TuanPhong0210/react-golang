// ===========================================
// Package controllers - Auth Controller
// ===========================================
// File này chứa các handler cho authentication endpoints
// Bao gồm: login, logout, get current user
// ===========================================

package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/hanbiro/hrms-server/services"
	"github.com/hanbiro/hrms-server/utils"
)

// AuthController chứa các handler cho auth routes
type AuthController struct {
	authService services.AuthService
}

// NewAuthController tạo instance mới của AuthController
//
// Input:
//   - authService: Auth service instance
//
// Output:
//   - *AuthController
func NewAuthController(authService services.AuthService) *AuthController {
	return &AuthController{authService: authService}
}

// Login xử lý đăng nhập
// @route POST /api/auth/login
//
// Input (JSON body):
//   - email: Email đăng nhập
//   - password: Mật khẩu
//
// Output:
//   - 200: Token và thông tin user
//   - 400: Email hoặc password không đúng
func (c *AuthController) Login(ctx *gin.Context) {
	// Parse input
	var input services.LoginInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		utils.BadRequestResponse(ctx, "Vui lòng nhập email và mật khẩu")
		return
	}

	// Gọi service để login
	user, token, err := c.authService.Login(input.Email, input.Password)
	if err != nil {
		utils.UnauthorizedResponse(ctx, err.Error())
		return
	}

	// Trả về response
	utils.SuccessResponse(ctx, "Đăng nhập thành công", services.LoginResponse{
		Token: token,
		User:  user.ToResponse(),
	})
}

// Logout xử lý đăng xuất
// @route POST /api/auth/logout
//
// Trong JWT-based auth, logout chỉ cần xóa token ở client
// Server-side chỉ trả về success message
//
// Output:
//   - 200: Logout thành công
func (c *AuthController) Logout(ctx *gin.Context) {
	// JWT stateless nên không cần xử lý gì ở server
	// Client sẽ xóa token
	utils.SuccessResponse(ctx, "Đăng xuất thành công", nil)
}

// GetCurrentUser lấy thông tin user hiện tại
// @route GET /api/auth/me
//
// Yêu cầu JWT token trong header: Authorization: Bearer <token>
//
// Output:
//   - 200: Thông tin user
//   - 401: Token không hợp lệ
func (c *AuthController) GetCurrentUser(ctx *gin.Context) {
	// Lấy userID từ context (được set bởi auth middleware)
	userID, exists := ctx.Get("userID")
	if !exists {
		utils.UnauthorizedResponse(ctx, "Không tìm thấy thông tin xác thực")
		return
	}

	// Gọi service để lấy user
	user, err := c.authService.GetCurrentUser(userID.(uint))
	if err != nil {
		utils.UnauthorizedResponse(ctx, err.Error())
		return
	}

	// Trả về response
	utils.SuccessResponse(ctx, "Lấy thông tin user thành công", user.ToResponse())
}
