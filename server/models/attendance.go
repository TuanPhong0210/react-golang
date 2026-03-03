// ===========================================
// Package models - Model Attendance
// ===========================================
// Định nghĩa struct Attendance (Chấm công)
// Lưu trữ thông tin check-in/check-out của nhân viên
// ===========================================

package models

import (
	"time"
)

// Attendance đại diện cho một bản ghi chấm công
// Mỗi nhân viên chỉ có 1 bản ghi chấm công mỗi ngày
type Attendance struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	UserID    uint       `gorm:"index;not null" json:"user_id"`
	CheckIn   time.Time  `gorm:"not null" json:"check_in"`
	CheckOut  *time.Time `json:"check_out"` // Pointer vì có thể null
	Date      time.Time  `gorm:"type:date;not null" json:"date"`
	Note      string     `gorm:"size:255" json:"note"`
	CreatedAt time.Time  `gorm:"autoCreateTime" json:"created_at"`

	// Relationship: Attendance thuộc về 1 User
	User User `gorm:"foreignKey:UserID" json:"user,omitempty"`
}

// TableName chỉ định tên bảng trong database
func (Attendance) TableName() string {
	return "attendances"
}

// AttendanceResponse là struct trả về cho client
// Bao gồm thông tin user name và tính toán working hours
type AttendanceResponse struct {
	ID           uint       `json:"id"`
	UserID       uint       `json:"user_id"`
	UserName     string     `json:"user_name"`
	CheckIn      time.Time  `json:"check_in"`
	CheckOut     *time.Time `json:"check_out"`
	Date         string     `json:"date"` // Format: YYYY-MM-DD
	Note         string     `json:"note"`
	WorkingHours float64    `json:"working_hours"` // Số giờ làm việc
	CreatedAt    time.Time  `json:"created_at"`
}

// ToResponse chuyển đổi Attendance sang AttendanceResponse
//
// Input: không có
// Output: AttendanceResponse - struct trả về client
func (a *Attendance) ToResponse() AttendanceResponse {
	response := AttendanceResponse{
		ID:        a.ID,
		UserID:    a.UserID,
		UserName:  a.User.FullName,
		CheckIn:   a.CheckIn,
		CheckOut:  a.CheckOut,
		Date:      a.Date.Format("2006-01-02"),
		Note:      a.Note,
		CreatedAt: a.CreatedAt,
	}

	// Tính working hours nếu đã check-out
	if a.CheckOut != nil {
		duration := a.CheckOut.Sub(a.CheckIn)
		response.WorkingHours = duration.Hours()
	}

	return response
}

// CheckInInput là struct để check-in
type CheckInInput struct {
	Note string `json:"note"`
}

// CheckOutInput là struct để check-out
type CheckOutInput struct {
	Note string `json:"note"`
}

// AttendanceFilter là struct để filter danh sách chấm công
type AttendanceFilter struct {
	UserID    uint   `form:"user_id"`
	StartDate string `form:"start_date"` // Format: YYYY-MM-DD
	EndDate   string `form:"end_date"`   // Format: YYYY-MM-DD
}
