// ===========================================
// Package repositories - Approval Repository
// ===========================================
// File này chứa các hàm thao tác database cho Approval
// Sử dụng GORM để query và manipulate dữ liệu
// ===========================================

package repositories

import (
	"github.com/hanbiro/hrms-server/models"
	"gorm.io/gorm"
)

// ApprovalRepository là interface định nghĩa các phương thức
// cho Approval repository
type ApprovalRepository interface {
	Create(approval *models.Approval) error
	FindByID(id uint) (*models.Approval, error)
	FindByUserID(userID uint) ([]models.Approval, error)
	FindAll(filter models.ApprovalFilter, page, limit int) ([]models.Approval, int64, error)
	FindPending(page, limit int) ([]models.Approval, int64, error)
	Update(approval *models.Approval) error
}

// approvalRepository là implementation của ApprovalRepository
type approvalRepository struct {
	db *gorm.DB
}

// NewApprovalRepository tạo instance mới của ApprovalRepository
//
// Input:
//   - db: GORM database instance
//
// Output:
//   - ApprovalRepository interface
func NewApprovalRepository(db *gorm.DB) ApprovalRepository {
	return &approvalRepository{db: db}
}

// Create tạo đơn xin mới
//
// Input:
//   - approval: Con trỏ đến Approval model
//
// Output:
//   - error: Lỗi nếu có
func (r *approvalRepository) Create(approval *models.Approval) error {
	return r.db.Create(approval).Error
}

// FindByID tìm approval theo ID
//
// Input:
//   - id: ID của approval
//
// Output:
//   - *models.Approval: Approval nếu tìm thấy
//   - error: Lỗi nếu không tìm thấy
func (r *approvalRepository) FindByID(id uint) (*models.Approval, error) {
	var approval models.Approval
	// Preload User và Approver để lấy thông tin
	err := r.db.Preload("User").Preload("Approver").First(&approval, id).Error
	if err != nil {
		return nil, err
	}
	return &approval, nil
}

// FindByUserID tìm tất cả approval của user
//
// Input:
//   - userID: ID của user
//
// Output:
//   - []models.Approval: Danh sách approval
//   - error: Lỗi nếu có
func (r *approvalRepository) FindByUserID(userID uint) ([]models.Approval, error) {
	var approvals []models.Approval
	err := r.db.Preload("User").Preload("Approver").
		Where("user_id = ?", userID).
		Order("created_at DESC").
		Find(&approvals).Error

	if err != nil {
		return nil, err
	}
	return approvals, nil
}

// FindAll lấy danh sách approval với filter và phân trang
//
// Input:
//   - filter: Điều kiện lọc (user_id, type, status)
//   - page: Số trang
//   - limit: Số item mỗi trang
//
// Output:
//   - []models.Approval: Danh sách approval
//   - int64: Tổng số records
//   - error: Lỗi nếu có
func (r *approvalRepository) FindAll(filter models.ApprovalFilter, page, limit int) ([]models.Approval, int64, error) {
	var approvals []models.Approval
	var total int64

	query := r.db.Model(&models.Approval{})

	// Áp dụng filter
	if filter.UserID > 0 {
		query = query.Where("user_id = ?", filter.UserID)
	}
	if filter.Type != "" {
		query = query.Where("type = ?", filter.Type)
	}
	if filter.Status != "" {
		query = query.Where("status = ?", filter.Status)
	}

	// Đếm tổng
	query.Count(&total)

	// Query với pagination
	offset := (page - 1) * limit
	err := query.Preload("User").Preload("Approver").
		Order("created_at DESC").
		Offset(offset).
		Limit(limit).
		Find(&approvals).Error

	if err != nil {
		return nil, 0, err
	}

	return approvals, total, nil
}

// FindPending lấy danh sách đơn đang chờ duyệt
//
// Input:
//   - page: Số trang
//   - limit: Số item mỗi trang
//
// Output:
//   - []models.Approval: Danh sách approval pending
//   - int64: Tổng số records
//   - error: Lỗi nếu có
func (r *approvalRepository) FindPending(page, limit int) ([]models.Approval, int64, error) {
	var approvals []models.Approval
	var total int64

	query := r.db.Model(&models.Approval{}).Where("status = ?", models.ApprovalStatusPending)

	// Đếm tổng
	query.Count(&total)

	// Query với pagination
	offset := (page - 1) * limit
	err := query.Preload("User").Preload("Approver").
		Order("created_at ASC"). // Đơn cũ lên trước
		Offset(offset).
		Limit(limit).
		Find(&approvals).Error

	if err != nil {
		return nil, 0, err
	}

	return approvals, total, nil
}

// Update cập nhật approval (dùng khi approve/reject)
//
// Input:
//   - approval: Con trỏ đến Approval model với data mới
//
// Output:
//   - error: Lỗi nếu có
func (r *approvalRepository) Update(approval *models.Approval) error {
	return r.db.Save(approval).Error
}
