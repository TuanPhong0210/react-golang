// ===========================================
// Package controllers - User Controller
// ===========================================
// File này chứa các handler cho user management endpoints
// Bao gồm: CRUD users, change password
// ===========================================

package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hanbiro/hrms-server/models"
	"github.com/hanbiro/hrms-server/services"
	"github.com/hanbiro/hrms-server/utils"
)

// UserController chứa các handler cho user routes
type UserController struct {
	userService services.UserService
}

// NewUserController tạo instance mới của UserController
//
// Input:
//   - userService: User service instance
//
// Output:
//   - *UserController
func NewUserController(userService services.UserService) *UserController {
	return &UserController{userService: userService}
}

// GetAllUsers lấy danh sách tất cả users
// @route GET /api/users
// @access Admin, Manager
//
// Query params:
//   - page: Số trang (default: 1)
//   - limit: Số item mỗi trang (default: 10)
//
// Output:
//   - 200: Danh sách users với pagination
func (c *UserController) GetAllUsers(ctx *gin.Context) {
	// Parse query params
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "10"))

	// Gọi service
	users, total, err := c.userService.GetAllUsers(page, limit)
	if err != nil {
		utils.InternalServerErrorResponse(ctx, "Không thể lấy danh sách users", err.Error())
		return
	}

	// Chuyển đổi sang response format
	var userResponses []models.UserResponse
	for _, user := range users {
		userResponses = append(userResponses, user.ToResponse())
	}

	// Trả về với pagination
	utils.PaginatedSuccessResponse(ctx, "Lấy danh sách users thành công", userResponses, page, limit, total)
}

// GetUserByID lấy thông tin user theo ID
// @route GET /api/users/:id
// @access All authenticated users
//
// Params:
//   - id: ID của user
//
// Output:
//   - 200: Thông tin user
//   - 404: Không tìm thấy user
func (c *UserController) GetUserByID(ctx *gin.Context) {
	// Parse ID từ URL
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		utils.BadRequestResponse(ctx, "ID không hợp lệ")
		return
	}

	// Gọi service
	user, err := c.userService.GetUserByID(uint(id))
	if err != nil {
		utils.NotFoundResponse(ctx, err.Error())
		return
	}

	utils.SuccessResponse(ctx, "Lấy thông tin user thành công", user.ToResponse())
}

// CreateUser tạo user mới
// @route POST /api/users
// @access Admin only
//
// Input (JSON body):
//   - email, password, full_name, phone, department, position, role_id
//
// Output:
//   - 201: User đã tạo
//   - 400: Dữ liệu không hợp lệ
func (c *UserController) CreateUser(ctx *gin.Context) {
	// Parse input
	var input models.CreateUserInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		utils.BadRequestResponse(ctx, "Dữ liệu không hợp lệ: "+err.Error())
		return
	}

	// Gọi service
	user, err := c.userService.CreateUser(input)
	if err != nil {
		utils.BadRequestResponse(ctx, err.Error())
		return
	}

	utils.CreatedResponse(ctx, "Tạo user thành công", user.ToResponse())
}

// UpdateUser cập nhật thông tin user
// @route PUT /api/users/:id
// @access Admin, Manager
//
// Params:
//   - id: ID của user cần cập nhật
//
// Input (JSON body):
//   - full_name, phone, department, position, role_id, is_active
//
// Output:
//   - 200: User sau khi cập nhật
//   - 404: Không tìm thấy user
func (c *UserController) UpdateUser(ctx *gin.Context) {
	// Parse ID từ URL
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		utils.BadRequestResponse(ctx, "ID không hợp lệ")
		return
	}

	// Parse input
	var input models.UpdateUserInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		utils.BadRequestResponse(ctx, "Dữ liệu không hợp lệ")
		return
	}

	// Gọi service
	user, err := c.userService.UpdateUser(uint(id), input)
	if err != nil {
		utils.NotFoundResponse(ctx, err.Error())
		return
	}

	utils.SuccessResponse(ctx, "Cập nhật user thành công", user.ToResponse())
}

// DeleteUser xóa user
// @route DELETE /api/users/:id
// @access Admin only
//
// Params:
//   - id: ID của user cần xóa
//
// Output:
//   - 200: Xóa thành công
//   - 404: Không tìm thấy user
func (c *UserController) DeleteUser(ctx *gin.Context) {
	// Parse ID từ URL
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		utils.BadRequestResponse(ctx, "ID không hợp lệ")
		return
	}

	// Gọi service
	err = c.userService.DeleteUser(uint(id))
	if err != nil {
		utils.NotFoundResponse(ctx, err.Error())
		return
	}

	utils.SuccessResponse(ctx, "Xóa user thành công", nil)
}

// ChangePassword đổi mật khẩu
// @route PUT /api/users/change-password
// @access All authenticated users (chỉ đổi password của chính mình)
//
// Input (JSON body):
//   - old_password: Mật khẩu cũ
//   - new_password: Mật khẩu mới
//
// Output:
//   - 200: Đổi mật khẩu thành công
//   - 400: Mật khẩu cũ không đúng
func (c *UserController) ChangePassword(ctx *gin.Context) {
	// Lấy userID từ context
	userID, _ := ctx.Get("userID")

	// Parse input
	var input struct {
		OldPassword string `json:"old_password" binding:"required"`
		NewPassword string `json:"new_password" binding:"required,min=6"`
	}
	if err := ctx.ShouldBindJSON(&input); err != nil {
		utils.BadRequestResponse(ctx, "Vui lòng nhập mật khẩu cũ và mật khẩu mới (tối thiểu 6 ký tự)")
		return
	}

	// Gọi service
	err := c.userService.ChangePassword(userID.(uint), input.OldPassword, input.NewPassword)
	if err != nil {
		utils.BadRequestResponse(ctx, err.Error())
		return
	}

	utils.SuccessResponse(ctx, "Đổi mật khẩu thành công", nil)
}
