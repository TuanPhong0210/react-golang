// ===========================================
// Package models - Model User
// ===========================================
// Định nghĩa struct User và các phương thức liên quan
// User là thực thể chính đại diện cho nhân viên
// ===========================================

package models

import (
	"time"
)

// User đại diện cho một nhân viên trong hệ thống
// Chứa thông tin cá nhân, tài khoản, và liên kết với Role
type User struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	Email      string    `gorm:"size:255;uniqueIndex;not null" json:"email"`
	Password   string    `gorm:"size:255;not null" json:"-"` // json:"-" ẩn password khi trả về JSON
	FullName   string    `gorm:"size:100;not null" json:"full_name"`
	Phone      string    `gorm:"size:20" json:"phone"`
	Department string    `gorm:"size:100" json:"department"`
	Position   string    `gorm:"size:100" json:"position"`
	RoleID     uint      `gorm:"index" json:"role_id"`
	IsActive   bool      `gorm:"default:true" json:"is_active"`
	CreatedAt  time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime" json:"updated_at"`

	// Relationship: User thuộc về 1 Role
	Role Role `gorm:"foreignKey:RoleID" json:"role,omitempty"`

	// Relationship: User có nhiều Attendances
	Attendances []Attendance `gorm:"foreignKey:UserID" json:"attendances,omitempty"`

	// Relationship: User có nhiều Approvals
	Approvals []Approval `gorm:"foreignKey:UserID" json:"approvals,omitempty"`
}

// TableName chỉ định tên bảng trong database
func (User) TableName() string {
	return "users"
}

// UserResponse là struct trả về cho client
// Không chứa password và thêm thông tin role name
type UserResponse struct {
	ID         uint      `json:"id"`
	Email      string    `json:"email"`
	FullName   string    `json:"full_name"`
	Phone      string    `json:"phone"`
	Department string    `json:"department"`
	Position   string    `json:"position"`
	RoleID     uint      `json:"role_id"`
	RoleName   string    `json:"role_name"`
	IsActive   bool      `json:"is_active"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

// ToResponse chuyển đổi User sang UserResponse
// Loại bỏ password và thêm role name
//
// Input: không có
// Output: UserResponse - struct an toàn để trả về client
func (u *User) ToResponse() UserResponse {
	return UserResponse{
		ID:         u.ID,
		Email:      u.Email,
		FullName:   u.FullName,
		Phone:      u.Phone,
		Department: u.Department,
		Position:   u.Position,
		RoleID:     u.RoleID,
		RoleName:   u.Role.Name,
		IsActive:   u.IsActive,
		CreatedAt:  u.CreatedAt,
		UpdatedAt:  u.UpdatedAt,
	}
}

// CreateUserInput là struct để tạo user mới
type CreateUserInput struct {
	Email      string `json:"email" binding:"required,email"`
	Password   string `json:"password" binding:"required,min=6"`
	FullName   string `json:"full_name" binding:"required"`
	Phone      string `json:"phone"`
	Department string `json:"department"`
	Position   string `json:"position"`
	RoleID     uint   `json:"role_id" binding:"required"`
}

// UpdateUserInput là struct để cập nhật user
type UpdateUserInput struct {
	FullName   string `json:"full_name"`
	Phone      string `json:"phone"`
	Department string `json:"department"`
	Position   string `json:"position"`
	RoleID     uint   `json:"role_id"`
	IsActive   *bool  `json:"is_active"` // Pointer để phân biệt false và không gửi
}
