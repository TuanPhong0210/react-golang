// ===========================================
// Package models - Model Approval
// ===========================================
// Định nghĩa struct Approval (Phê duyệt)
// Lưu trữ đơn xin nghỉ phép, OT
// ===========================================

package models

import (
	"time"
)

// ApprovalType là loại đơn xin
type ApprovalType string

const (
	ApprovalTypeLeave ApprovalType = "leave" // Nghỉ phép
	ApprovalTypeOT    ApprovalType = "ot"    // Làm thêm giờ
)

// ApprovalStatus là trạng thái đơn
type ApprovalStatus string

const (
	ApprovalStatusPending  ApprovalStatus = "pending"  // Đang chờ duyệt
	ApprovalStatusApproved ApprovalStatus = "approved" // Đã duyệt
	ApprovalStatusRejected ApprovalStatus = "rejected" // Từ chối
)

// Approval đại diện cho một đơn xin phê duyệt
// Có thể là đơn xin nghỉ phép hoặc làm thêm giờ
type Approval struct {
	ID         uint           `gorm:"primaryKey" json:"id"`
	UserID     uint           `gorm:"index;not null" json:"user_id"`
	Type       ApprovalType   `gorm:"size:50;not null" json:"type"`
	StartDate  time.Time      `gorm:"type:date;not null" json:"start_date"`
	EndDate    time.Time      `gorm:"type:date;not null" json:"end_date"`
	Reason     string         `gorm:"type:text" json:"reason"`
	Status     ApprovalStatus `gorm:"size:20;default:pending" json:"status"`
	ApprovedBy *uint          `json:"approved_by"` // ID của người duyệt
	ApprovedAt *time.Time     `json:"approved_at"`
	CreatedAt  time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt  time.Time      `gorm:"autoUpdateTime" json:"updated_at"`

	// Relationship: Approval thuộc về 1 User (người tạo đơn)
	User User `gorm:"foreignKey:UserID" json:"user,omitempty"`

	// Relationship: Approval được duyệt bởi 1 User (người duyệt)
	Approver *User `gorm:"foreignKey:ApprovedBy" json:"approver,omitempty"`
}

// TableName chỉ định tên bảng trong database
func (Approval) TableName() string {
	return "approvals"
}

// ApprovalResponse là struct trả về cho client
type ApprovalResponse struct {
	ID           uint           `json:"id"`
	UserID       uint           `json:"user_id"`
	UserName     string         `json:"user_name"`
	Type         ApprovalType   `json:"type"`
	TypeLabel    string         `json:"type_label"` // Nhãn tiếng Việt
	StartDate    string         `json:"start_date"` // Format: YYYY-MM-DD
	EndDate      string         `json:"end_date"`
	Days         int            `json:"days"` // Số ngày
	Reason       string         `json:"reason"`
	Status       ApprovalStatus `json:"status"`
	StatusLabel  string         `json:"status_label"` // Nhãn tiếng Việt
	ApprovedBy   *uint          `json:"approved_by"`
	ApproverName string         `json:"approver_name"`
	ApprovedAt   *time.Time     `json:"approved_at"`
	CreatedAt    time.Time      `json:"created_at"`
}

// ToResponse chuyển đổi Approval sang ApprovalResponse
//
// Input: không có
// Output: ApprovalResponse - struct trả về client
func (a *Approval) ToResponse() ApprovalResponse {
	response := ApprovalResponse{
		ID:         a.ID,
		UserID:     a.UserID,
		UserName:   a.User.FullName,
		Type:       a.Type,
		StartDate:  a.StartDate.Format("2006-01-02"),
		EndDate:    a.EndDate.Format("2006-01-02"),
		Reason:     a.Reason,
		Status:     a.Status,
		ApprovedBy: a.ApprovedBy,
		ApprovedAt: a.ApprovedAt,
		CreatedAt:  a.CreatedAt,
	}

	// Tính số ngày
	response.Days = int(a.EndDate.Sub(a.StartDate).Hours()/24) + 1

	// Label tiếng Việt cho type
	switch a.Type {
	case ApprovalTypeLeave:
		response.TypeLabel = "Nghỉ phép"
	case ApprovalTypeOT:
		response.TypeLabel = "Làm thêm giờ"
	}

	// Label tiếng Việt cho status
	switch a.Status {
	case ApprovalStatusPending:
		response.StatusLabel = "Đang chờ duyệt"
	case ApprovalStatusApproved:
		response.StatusLabel = "Đã duyệt"
	case ApprovalStatusRejected:
		response.StatusLabel = "Từ chối"
	}

	// Tên người duyệt
	if a.Approver != nil {
		response.ApproverName = a.Approver.FullName
	}

	return response
}

// CreateApprovalInput là struct để tạo đơn mới
type CreateApprovalInput struct {
	Type      ApprovalType `json:"type" binding:"required,oneof=leave ot"`
	StartDate string       `json:"start_date" binding:"required"` // Format: YYYY-MM-DD
	EndDate   string       `json:"end_date" binding:"required"`
	Reason    string       `json:"reason" binding:"required"`
}

// UpdateApprovalInput là struct để cập nhật trạng thái đơn
type UpdateApprovalInput struct {
	Status ApprovalStatus `json:"status" binding:"required,oneof=approved rejected"`
}

// ApprovalFilter là struct để filter danh sách đơn
type ApprovalFilter struct {
	UserID uint           `form:"user_id"`
	Type   ApprovalType   `form:"type"`
	Status ApprovalStatus `form:"status"`
}
