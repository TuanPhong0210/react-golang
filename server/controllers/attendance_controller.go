// ===========================================
// Package controllers - Attendance Controller
// ===========================================
// File này chứa các handler cho attendance endpoints
// Bao gồm: check-in, check-out, xem lịch sử
// ===========================================

package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hanbiro/hrms-server/models"
	"github.com/hanbiro/hrms-server/services"
	"github.com/hanbiro/hrms-server/utils"
)

// AttendanceController chứa các handler cho attendance routes
type AttendanceController struct {
	attendanceService services.AttendanceService
}

// NewAttendanceController tạo instance mới của AttendanceController
//
// Input:
//   - attendanceService: Attendance service instance
//
// Output:
//   - *AttendanceController
func NewAttendanceController(attendanceService services.AttendanceService) *AttendanceController {
	return &AttendanceController{attendanceService: attendanceService}
}

// CheckIn thực hiện check-in
// @route POST /api/attendance/check-in
// @access All authenticated users
//
// Input (JSON body):
//   - note: Ghi chú (optional)
//
// Output:
//   - 201: Bản ghi check-in
//   - 400: Đã check-in trong ngày
func (c *AttendanceController) CheckIn(ctx *gin.Context) {
	// Lấy userID từ context
	userID, _ := ctx.Get("userID")

	// Parse input
	var input models.CheckInInput
	ctx.ShouldBindJSON(&input) // Optional, không cần check error

	// Gọi service
	attendance, err := c.attendanceService.CheckIn(userID.(uint), input.Note)
	if err != nil {
		utils.BadRequestResponse(ctx, err.Error())
		return
	}

	utils.CreatedResponse(ctx, "Check-in thành công", attendance.ToResponse())
}

// CheckOut thực hiện check-out
// @route POST /api/attendance/check-out
// @access All authenticated users
//
// Input (JSON body):
//   - note: Ghi chú bổ sung (optional)
//
// Output:
//   - 200: Bản ghi attendance đã cập nhật
//   - 400: Chưa check-in hoặc đã check-out
func (c *AttendanceController) CheckOut(ctx *gin.Context) {
	// Lấy userID từ context
	userID, _ := ctx.Get("userID")

	// Parse input
	var input models.CheckOutInput
	ctx.ShouldBindJSON(&input) // Optional

	// Gọi service
	attendance, err := c.attendanceService.CheckOut(userID.(uint), input.Note)
	if err != nil {
		utils.BadRequestResponse(ctx, err.Error())
		return
	}

	utils.SuccessResponse(ctx, "Check-out thành công", attendance.ToResponse())
}

// GetTodayAttendance lấy thông tin chấm công ngày hôm nay
// @route GET /api/attendance/today
// @access All authenticated users
//
// Output:
//   - 200: Bản ghi attendance (hoặc null nếu chưa check-in)
func (c *AttendanceController) GetTodayAttendance(ctx *gin.Context) {
	// Lấy userID từ context
	userID, _ := ctx.Get("userID")

	// Gọi service
	attendance, _ := c.attendanceService.GetTodayAttendance(userID.(uint))

	if attendance == nil {
		utils.SuccessResponse(ctx, "Chưa check-in hôm nay", nil)
		return
	}

	utils.SuccessResponse(ctx, "Lấy thông tin chấm công thành công", attendance.ToResponse())
}

// GetMyAttendanceHistory lấy lịch sử chấm công của user hiện tại
// @route GET /api/attendance/history
// @access All authenticated users
//
// Query params:
//   - start_date: Ngày bắt đầu (YYYY-MM-DD)
//   - end_date: Ngày kết thúc (YYYY-MM-DD)
//
// Output:
//   - 200: Danh sách attendance
func (c *AttendanceController) GetMyAttendanceHistory(ctx *gin.Context) {
	// Lấy userID từ context
	userID, _ := ctx.Get("userID")

	// Parse query params
	startDate := ctx.DefaultQuery("start_date", "")
	endDate := ctx.DefaultQuery("end_date", "")

	// Gọi service
	attendances, err := c.attendanceService.GetAttendanceHistory(userID.(uint), startDate, endDate)
	if err != nil {
		utils.InternalServerErrorResponse(ctx, "Không thể lấy lịch sử chấm công", err.Error())
		return
	}

	// Chuyển đổi sang response format
	var responses []models.AttendanceResponse
	for _, att := range attendances {
		responses = append(responses, att.ToResponse())
	}

	utils.SuccessResponse(ctx, "Lấy lịch sử chấm công thành công", responses)
}

// GetAllAttendances lấy danh sách tất cả attendance
// @route GET /api/attendance
// @access Admin, Manager
//
// Query params:
//   - user_id: Filter theo user (optional)
//   - start_date: Ngày bắt đầu (optional)
//   - end_date: Ngày kết thúc (optional)
//   - page: Số trang (default: 1)
//   - limit: Số item mỗi trang (default: 10)
//
// Output:
//   - 200: Danh sách attendance với pagination
func (c *AttendanceController) GetAllAttendances(ctx *gin.Context) {
	// Parse filter
	var filter models.AttendanceFilter
	ctx.ShouldBindQuery(&filter)

	// Parse pagination
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "10"))

	// Gọi service
	attendances, total, err := c.attendanceService.GetAllAttendances(filter, page, limit)
	if err != nil {
		utils.InternalServerErrorResponse(ctx, "Không thể lấy danh sách chấm công", err.Error())
		return
	}

	// Chuyển đổi sang response format
	var responses []models.AttendanceResponse
	for _, att := range attendances {
		responses = append(responses, att.ToResponse())
	}

	utils.PaginatedSuccessResponse(ctx, "Lấy danh sách chấm công thành công", responses, page, limit, total)
}
