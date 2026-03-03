// ===========================================
// Package controllers - Approval Controller
// ===========================================
// File này chứa các handler cho approval endpoints
// Bao gồm: tạo đơn, xem đơn, duyệt/từ chối đơn
// ===========================================

package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hanbiro/hrms-server/models"
	"github.com/hanbiro/hrms-server/services"
	"github.com/hanbiro/hrms-server/utils"
)

// ApprovalController chứa các handler cho approval routes
type ApprovalController struct {
	approvalService services.ApprovalService
}

// NewApprovalController tạo instance mới của ApprovalController
//
// Input:
//   - approvalService: Approval service instance
//
// Output:
//   - *ApprovalController
func NewApprovalController(approvalService services.ApprovalService) *ApprovalController {
	return &ApprovalController{approvalService: approvalService}
}

// CreateApproval tạo đơn xin mới
// @route POST /api/approvals
// @access All authenticated users
//
// Input (JSON body):
//   - type: Loại đơn ("leave" hoặc "ot")
//   - start_date: Ngày bắt đầu (YYYY-MM-DD)
//   - end_date: Ngày kết thúc (YYYY-MM-DD)
//   - reason: Lý do
//
// Output:
//   - 201: Đơn đã tạo
//   - 400: Dữ liệu không hợp lệ
func (c *ApprovalController) CreateApproval(ctx *gin.Context) {
	// Lấy userID từ context
	userID, _ := ctx.Get("userID")

	// Parse input
	var input models.CreateApprovalInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		utils.BadRequestResponse(ctx, "Dữ liệu không hợp lệ: "+err.Error())
		return
	}

	// Gọi service
	approval, err := c.approvalService.CreateApproval(userID.(uint), input)
	if err != nil {
		utils.BadRequestResponse(ctx, err.Error())
		return
	}

	utils.CreatedResponse(ctx, "Tạo đơn thành công", approval.ToResponse())
}

// GetMyApprovals lấy danh sách đơn của user hiện tại
// @route GET /api/approvals/my
// @access All authenticated users
//
// Output:
//   - 200: Danh sách đơn
func (c *ApprovalController) GetMyApprovals(ctx *gin.Context) {
	// Lấy userID từ context
	userID, _ := ctx.Get("userID")

	// Gọi service
	approvals, err := c.approvalService.GetMyApprovals(userID.(uint))
	if err != nil {
		utils.InternalServerErrorResponse(ctx, "Không thể lấy danh sách đơn", err.Error())
		return
	}

	// Chuyển đổi sang response format
	var responses []models.ApprovalResponse
	for _, app := range approvals {
		responses = append(responses, app.ToResponse())
	}

	utils.SuccessResponse(ctx, "Lấy danh sách đơn thành công", responses)
}

// GetApprovalByID lấy thông tin đơn theo ID
// @route GET /api/approvals/:id
// @access All authenticated users
//
// Output:
//   - 200: Thông tin đơn
//   - 404: Không tìm thấy đơn
func (c *ApprovalController) GetApprovalByID(ctx *gin.Context) {
	// Parse ID từ URL
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		utils.BadRequestResponse(ctx, "ID không hợp lệ")
		return
	}

	// Gọi service
	approval, err := c.approvalService.GetApprovalByID(uint(id))
	if err != nil {
		utils.NotFoundResponse(ctx, err.Error())
		return
	}

	utils.SuccessResponse(ctx, "Lấy thông tin đơn thành công", approval.ToResponse())
}

// GetAllApprovals lấy danh sách tất cả đơn
// @route GET /api/approvals
// @access Admin, Manager
//
// Query params:
//   - user_id: Filter theo user (optional)
//   - type: Filter theo loại (optional)
//   - status: Filter theo trạng thái (optional)
//   - page: Số trang (default: 1)
//   - limit: Số item mỗi trang (default: 10)
//
// Output:
//   - 200: Danh sách đơn với pagination
func (c *ApprovalController) GetAllApprovals(ctx *gin.Context) {
	// Parse filter
	var filter models.ApprovalFilter
	ctx.ShouldBindQuery(&filter)

	// Parse pagination
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "10"))

	// Gọi service
	approvals, total, err := c.approvalService.GetAllApprovals(filter, page, limit)
	if err != nil {
		utils.InternalServerErrorResponse(ctx, "Không thể lấy danh sách đơn", err.Error())
		return
	}

	// Chuyển đổi sang response format
	var responses []models.ApprovalResponse
	for _, app := range approvals {
		responses = append(responses, app.ToResponse())
	}

	utils.PaginatedSuccessResponse(ctx, "Lấy danh sách đơn thành công", responses, page, limit, total)
}

// GetPendingApprovals lấy danh sách đơn đang chờ duyệt
// @route GET /api/approvals/pending
// @access Admin, Manager
//
// Query params:
//   - page: Số trang (default: 1)
//   - limit: Số item mỗi trang (default: 10)
//
// Output:
//   - 200: Danh sách đơn pending
func (c *ApprovalController) GetPendingApprovals(ctx *gin.Context) {
	// Parse pagination
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "10"))

	// Gọi service
	approvals, total, err := c.approvalService.GetPendingApprovals(page, limit)
	if err != nil {
		utils.InternalServerErrorResponse(ctx, "Không thể lấy danh sách đơn", err.Error())
		return
	}

	// Chuyển đổi sang response format
	var responses []models.ApprovalResponse
	for _, app := range approvals {
		responses = append(responses, app.ToResponse())
	}

	utils.PaginatedSuccessResponse(ctx, "Lấy danh sách đơn chờ duyệt thành công", responses, page, limit, total)
}

// ApproveRequest duyệt đơn
// @route PUT /api/approvals/:id/approve
// @access Admin, Manager
//
// Output:
//   - 200: Đơn đã duyệt
//   - 400: Đơn đã được xử lý
//   - 404: Không tìm thấy đơn
func (c *ApprovalController) ApproveRequest(ctx *gin.Context) {
	// Parse ID từ URL
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		utils.BadRequestResponse(ctx, "ID không hợp lệ")
		return
	}

	// Lấy approverID từ context
	approverID, _ := ctx.Get("userID")

	// Gọi service
	approval, err := c.approvalService.ApproveRequest(uint(id), approverID.(uint))
	if err != nil {
		if err.Error() == "không tìm thấy đơn" {
			utils.NotFoundResponse(ctx, err.Error())
		} else {
			utils.BadRequestResponse(ctx, err.Error())
		}
		return
	}

	utils.SuccessResponse(ctx, "Duyệt đơn thành công", approval.ToResponse())
}

// RejectRequest từ chối đơn
// @route PUT /api/approvals/:id/reject
// @access Admin, Manager
//
// Output:
//   - 200: Đơn đã từ chối
//   - 400: Đơn đã được xử lý
//   - 404: Không tìm thấy đơn
func (c *ApprovalController) RejectRequest(ctx *gin.Context) {
	// Parse ID từ URL
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		utils.BadRequestResponse(ctx, "ID không hợp lệ")
		return
	}

	// Lấy approverID từ context
	approverID, _ := ctx.Get("userID")

	// Gọi service
	approval, err := c.approvalService.RejectRequest(uint(id), approverID.(uint))
	if err != nil {
		if err.Error() == "không tìm thấy đơn" {
			utils.NotFoundResponse(ctx, err.Error())
		} else {
			utils.BadRequestResponse(ctx, err.Error())
		}
		return
	}

	utils.SuccessResponse(ctx, "Từ chối đơn thành công", approval.ToResponse())
}
