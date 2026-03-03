// ===========================================
// Package models - Model Role
// ===========================================
// Định nghĩa struct Role và các phương thức liên quan
// Role: Admin, Manager, Employee
// ===========================================

package models

import (
	"time"
)

// Role đại diện cho vai trò của người dùng trong hệ thống
// Có 3 vai trò chính:
// - Admin: Quản trị viên, full quyền
// - Manager: Quản lý, duyệt đơn, quản lý nhân viên
// - Employee: Nhân viên, chấm công, tạo đơn
type Role struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Name        string    `gorm:"size:50;uniqueIndex;not null" json:"name"`
	Description string    `gorm:"size:255" json:"description"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`

	// Relationship: Role có nhiều Users
	Users []User `gorm:"foreignKey:RoleID" json:"users,omitempty"`
}

// TableName chỉ định tên bảng trong database
// GORM sẽ sử dụng tên này thay vì convention mặc định
func (Role) TableName() string {
	return "roles"
}

// RoleNames là các tên role hợp lệ trong hệ thống
// Sử dụng để validate role
var RoleNames = []string{"Admin", "Manager", "Employee"}

// IsAdmin kiểm tra xem role có phải là Admin không
func (r *Role) IsAdmin() bool {
	return r.Name == "Admin"
}

// IsManager kiểm tra xem role có phải là Manager không
func (r *Role) IsManager() bool {
	return r.Name == "Manager"
}

// IsEmployee kiểm tra xem role có phải là Employee không
func (r *Role) IsEmployee() bool {
	return r.Name == "Employee"
}

// CanManageUsers kiểm tra xem role có quyền quản lý users không
// Admin và Manager đều có quyền này
func (r *Role) CanManageUsers() bool {
	return r.Name == "Admin" || r.Name == "Manager"
}

// CanApproveRequests kiểm tra xem role có quyền duyệt đơn không
// Admin và Manager đều có quyền này
func (r *Role) CanApproveRequests() bool {
	return r.Name == "Admin" || r.Name == "Manager"
}
