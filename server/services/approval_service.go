// ===========================================
// Package services - Approval Service
// ===========================================
// File này chứa logic xử lý nghiệp vụ cho Approval
// Bao gồm: tạo đơn, duyệt đơn, từ chối đơn
// ===========================================

package services

import (
	"errors"
	"time"

	"github.com/hanbiro/hrms-server/models"
	"github.com/hanbiro/hrms-server/repositories"
)

// ApprovalService là interface định nghĩa các phương thức
// cho Approval service
type ApprovalService interface {
	CreateApproval(userID uint, input models.CreateApprovalInput) (*models.Approval, error)
	GetApprovalByID(id uint) (*models.Approval, error)
	GetMyApprovals(userID uint) ([]models.Approval, error)
	GetAllApprovals(filter models.ApprovalFilter, page, limit int) ([]models.Approval, int64, error)
	GetPendingApprovals(page, limit int) ([]models.Approval, int64, error)
	ApproveRequest(id, approverID uint) (*models.Approval, error)
	RejectRequest(id, approverID uint) (*models.Approval, error)
}

// approvalService là implementation của ApprovalService
type approvalService struct {
	approvalRepo repositories.ApprovalRepository
}

// NewApprovalService tạo instance mới của ApprovalService
//
// Input:
//   - approvalRepo: Approval repository instance
//
// Output:
//   - ApprovalService interface
func NewApprovalService(approvalRepo repositories.ApprovalRepository) ApprovalService {
	return &approvalService{approvalRepo: approvalRepo}
}

// CreateApproval tạo đơn xin mới (nghỉ phép, OT)
//
// Input:
//   - userID: ID của user tạo đơn
//   - input: Thông tin đơn xin
//
// Output:
//   - *models.Approval: Đơn đã tạo
//   - error: Lỗi nếu có (vd: ngày không hợp lệ)
func (s *approvalService) CreateApproval(userID uint, input models.CreateApprovalInput) (*models.Approval, error) {
	// Parse dates
	startDate, err := time.Parse("2006-01-02", input.StartDate)
	if err != nil {
		return nil, errors.New("ngày bắt đầu không hợp lệ")
	}

	endDate, err := time.Parse("2006-01-02", input.EndDate)
	if err != nil {
		return nil, errors.New("ngày kết thúc không hợp lệ")
	}

	// Validate dates
	if endDate.Before(startDate) {
		return nil, errors.New("ngày kết thúc phải sau ngày bắt đầu")
	}

	// Ngày bắt đầu phải từ hôm nay trở đi (cho đơn mới)
	today := time.Now().Truncate(24 * time.Hour)
	if startDate.Before(today) {
		return nil, errors.New("ngày bắt đầu phải từ hôm nay trở đi")
	}

	// Tạo đơn mới
	approval := &models.Approval{
		UserID:    userID,
		Type:      input.Type,
		StartDate: startDate,
		EndDate:   endDate,
		Reason:    input.Reason,
		Status:    models.ApprovalStatusPending,
	}

	err = s.approvalRepo.Create(approval)
	if err != nil {
		return nil, errors.New("không thể tạo đơn")
	}

	// Lấy lại với thông tin user
	approval, _ = s.approvalRepo.FindByID(approval.ID)

	return approval, nil
}

// GetApprovalByID lấy thông tin đơn theo ID
//
// Input:
//   - id: ID của đơn
//
// Output:
//   - *models.Approval: Đơn nếu tìm thấy
//   - error: Lỗi nếu không tìm thấy
func (s *approvalService) GetApprovalByID(id uint) (*models.Approval, error) {
	approval, err := s.approvalRepo.FindByID(id)
	if err != nil {
		return nil, errors.New("không tìm thấy đơn")
	}
	return approval, nil
}

// GetMyApprovals lấy danh sách đơn của user hiện tại
//
// Input:
//   - userID: ID của user
//
// Output:
//   - []models.Approval: Danh sách đơn
//   - error: Lỗi nếu có
func (s *approvalService) GetMyApprovals(userID uint) ([]models.Approval, error) {
	return s.approvalRepo.FindByUserID(userID)
}

// GetAllApprovals lấy danh sách tất cả đơn với filter
// Dùng cho Admin/Manager
//
// Input:
//   - filter: Điều kiện lọc
//   - page: Số trang
//   - limit: Số item mỗi trang
//
// Output:
//   - []models.Approval: Danh sách đơn
//   - int64: Tổng số records
//   - error: Lỗi nếu có
func (s *approvalService) GetAllApprovals(filter models.ApprovalFilter, page, limit int) ([]models.Approval, int64, error) {
	// Validate pagination params
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	return s.approvalRepo.FindAll(filter, page, limit)
}

// GetPendingApprovals lấy danh sách đơn đang chờ duyệt
// Dùng cho Admin/Manager
//
// Input:
//   - page: Số trang
//   - limit: Số item mỗi trang
//
// Output:
//   - []models.Approval: Danh sách đơn pending
//   - int64: Tổng số records
//   - error: Lỗi nếu có
func (s *approvalService) GetPendingApprovals(page, limit int) ([]models.Approval, int64, error) {
	// Validate pagination params
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	return s.approvalRepo.FindPending(page, limit)
}

// ApproveRequest duyệt đơn
//
// Input:
//   - id: ID của đơn cần duyệt
//   - approverID: ID của người duyệt (Admin/Manager)
//
// Output:
//   - *models.Approval: Đơn sau khi duyệt
//   - error: Lỗi nếu đơn không ở trạng thái pending
func (s *approvalService) ApproveRequest(id, approverID uint) (*models.Approval, error) {
	// Tìm đơn
	approval, err := s.approvalRepo.FindByID(id)
	if err != nil {
		return nil, errors.New("không tìm thấy đơn")
	}

	// Kiểm tra trạng thái
	if approval.Status != models.ApprovalStatusPending {
		return nil, errors.New("đơn này đã được xử lý")
	}

	// Cập nhật trạng thái
	now := time.Now()
	approval.Status = models.ApprovalStatusApproved
	approval.ApprovedBy = &approverID
	approval.ApprovedAt = &now

	err = s.approvalRepo.Update(approval)
	if err != nil {
		return nil, errors.New("không thể cập nhật đơn")
	}

	// Lấy lại với thông tin đầy đủ
	approval, _ = s.approvalRepo.FindByID(approval.ID)

	return approval, nil
}

// RejectRequest từ chối đơn
//
// Input:
//   - id: ID của đơn cần từ chối
//   - approverID: ID của người từ chối (Admin/Manager)
//
// Output:
//   - *models.Approval: Đơn sau khi từ chối
//   - error: Lỗi nếu đơn không ở trạng thái pending
func (s *approvalService) RejectRequest(id, approverID uint) (*models.Approval, error) {
	// Tìm đơn
	approval, err := s.approvalRepo.FindByID(id)
	if err != nil {
		return nil, errors.New("không tìm thấy đơn")
	}

	// Kiểm tra trạng thái
	if approval.Status != models.ApprovalStatusPending {
		return nil, errors.New("đơn này đã được xử lý")
	}

	// Cập nhật trạng thái
	now := time.Now()
	approval.Status = models.ApprovalStatusRejected
	approval.ApprovedBy = &approverID
	approval.ApprovedAt = &now

	err = s.approvalRepo.Update(approval)
	if err != nil {
		return nil, errors.New("không thể cập nhật đơn")
	}

	// Lấy lại với thông tin đầy đủ
	approval, _ = s.approvalRepo.FindByID(approval.ID)

	return approval, nil
}
