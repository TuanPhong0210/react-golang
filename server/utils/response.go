// ===========================================
// Package utils - Response Utilities
// ===========================================
// File này chứa các hàm chuẩn hóa response API
// Giúp đảm bảo tất cả API trả về format nhất quán
// ===========================================

package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response là struct chuẩn cho tất cả API response
type Response struct {
	Success bool        `json:"success"`          // true nếu thành công
	Message string      `json:"message"`          // Thông báo
	Data    interface{} `json:"data,omitempty"`   // Dữ liệu trả về (nếu có)
	Error   string      `json:"error,omitempty"`  // Chi tiết lỗi (nếu có)
}

// PaginatedResponse là response có phân trang
type PaginatedResponse struct {
	Success    bool        `json:"success"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
	Pagination Pagination  `json:"pagination"`
}

// Pagination chứa thông tin phân trang
type Pagination struct {
	Page       int   `json:"page"`
	Limit      int   `json:"limit"`
	Total      int64 `json:"total"`
	TotalPages int   `json:"total_pages"`
}

// SuccessResponse trả về response thành công với data
//
// Input:
//   - c: Gin context
//   - message: Thông báo
//   - data: Dữ liệu trả về
//
// Output: JSON response với status 200
func SuccessResponse(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Success: true,
		Message: message,
		Data:    data,
	})
}

// CreatedResponse trả về response khi tạo mới thành công
//
// Input:
//   - c: Gin context
//   - message: Thông báo
//   - data: Dữ liệu trả về
//
// Output: JSON response với status 201
func CreatedResponse(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusCreated, Response{
		Success: true,
		Message: message,
		Data:    data,
	})
}

// ErrorResponse trả về response lỗi
//
// Input:
//   - c: Gin context
//   - statusCode: HTTP status code (400, 401, 403, 404, 500...)
//   - message: Thông báo lỗi
//   - err: Chi tiết lỗi (có thể để trống)
//
// Output: JSON response với status code tương ứng
func ErrorResponse(c *gin.Context, statusCode int, message string, err string) {
	c.JSON(statusCode, Response{
		Success: false,
		Message: message,
		Error:   err,
	})
}

// BadRequestResponse trả về lỗi 400 - Bad Request
//
// Input:
//   - c: Gin context
//   - message: Thông báo lỗi
//
// Output: JSON response với status 400
func BadRequestResponse(c *gin.Context, message string) {
	ErrorResponse(c, http.StatusBadRequest, message, "")
}

// UnauthorizedResponse trả về lỗi 401 - Unauthorized
//
// Input:
//   - c: Gin context
//   - message: Thông báo lỗi
//
// Output: JSON response với status 401
func UnauthorizedResponse(c *gin.Context, message string) {
	ErrorResponse(c, http.StatusUnauthorized, message, "")
}

// ForbiddenResponse trả về lỗi 403 - Forbidden
//
// Input:
//   - c: Gin context
//   - message: Thông báo lỗi
//
// Output: JSON response với status 403
func ForbiddenResponse(c *gin.Context, message string) {
	ErrorResponse(c, http.StatusForbidden, message, "")
}

// NotFoundResponse trả về lỗi 404 - Not Found
//
// Input:
//   - c: Gin context
//   - message: Thông báo lỗi
//
// Output: JSON response với status 404
func NotFoundResponse(c *gin.Context, message string) {
	ErrorResponse(c, http.StatusNotFound, message, "")
}

// InternalServerErrorResponse trả về lỗi 500 - Internal Server Error
//
// Input:
//   - c: Gin context
//   - message: Thông báo lỗi
//   - err: Chi tiết lỗi (để debug)
//
// Output: JSON response với status 500
func InternalServerErrorResponse(c *gin.Context, message string, err string) {
	ErrorResponse(c, http.StatusInternalServerError, message, err)
}

// PaginatedSuccessResponse trả về response thành công với phân trang
//
// Input:
//   - c: Gin context
//   - message: Thông báo
//   - data: Dữ liệu trả về
//   - page: Trang hiện tại
//   - limit: Số item mỗi trang
//   - total: Tổng số item
//
// Output: JSON response với status 200 và thông tin pagination
func PaginatedSuccessResponse(c *gin.Context, message string, data interface{}, page, limit int, total int64) {
	totalPages := int(total) / limit
	if int(total)%limit != 0 {
		totalPages++
	}

	c.JSON(http.StatusOK, PaginatedResponse{
		Success: true,
		Message: message,
		Data:    data,
		Pagination: Pagination{
			Page:       page,
			Limit:      limit,
			Total:      total,
			TotalPages: totalPages,
		},
	})
}
