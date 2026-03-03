// ===========================================
// Package repositories - Attendance Repository
// ===========================================
// File này chứa các hàm thao tác database cho Attendance
// Sử dụng GORM để query và manipulate dữ liệu
// ===========================================

package repositories

import (
	"time"

	"github.com/hanbiro/hrms-server/models"
	"gorm.io/gorm"
)

// AttendanceRepository là interface định nghĩa các phương thức
// cho Attendance repository
type AttendanceRepository interface {
	Create(attendance *models.Attendance) error
	FindByID(id uint) (*models.Attendance, error)
	FindByUserAndDate(userID uint, date time.Time) (*models.Attendance, error)
	FindByUserID(userID uint, startDate, endDate time.Time) ([]models.Attendance, error)
	FindAll(filter models.AttendanceFilter, page, limit int) ([]models.Attendance, int64, error)
	Update(attendance *models.Attendance) error
}

// attendanceRepository là implementation của AttendanceRepository
type attendanceRepository struct {
	db *gorm.DB
}

// NewAttendanceRepository tạo instance mới của AttendanceRepository
//
// Input:
//   - db: GORM database instance
//
// Output:
//   - AttendanceRepository interface
func NewAttendanceRepository(db *gorm.DB) AttendanceRepository {
	return &attendanceRepository{db: db}
}

// Create tạo bản ghi chấm công mới
//
// Input:
//   - attendance: Con trỏ đến Attendance model
//
// Output:
//   - error: Lỗi nếu có (vd: đã check-in trong ngày)
func (r *attendanceRepository) Create(attendance *models.Attendance) error {
	return r.db.Create(attendance).Error
}

// FindByID tìm attendance theo ID
//
// Input:
//   - id: ID của attendance
//
// Output:
//   - *models.Attendance: Attendance nếu tìm thấy
//   - error: Lỗi nếu không tìm thấy
func (r *attendanceRepository) FindByID(id uint) (*models.Attendance, error) {
	var attendance models.Attendance
	err := r.db.Preload("User").First(&attendance, id).Error
	if err != nil {
		return nil, err
	}
	return &attendance, nil
}

// FindByUserAndDate tìm attendance của user trong ngày cụ thể
// Sử dụng để kiểm tra đã check-in chưa
//
// Input:
//   - userID: ID của user
//   - date: Ngày cần kiểm tra
//
// Output:
//   - *models.Attendance: Attendance nếu tìm thấy
//   - error: Lỗi nếu không tìm thấy (có nghĩa chưa check-in)
func (r *attendanceRepository) FindByUserAndDate(userID uint, date time.Time) (*models.Attendance, error) {
	var attendance models.Attendance
	// Format date để so sánh chỉ ngày, không tính giờ
	dateStr := date.Format("2006-01-02")
	err := r.db.Preload("User").
		Where("user_id = ? AND DATE(date) = ?", userID, dateStr).
		First(&attendance).Error

	if err != nil {
		return nil, err
	}
	return &attendance, nil
}

// FindByUserID tìm tất cả attendance của user trong khoảng thời gian
//
// Input:
//   - userID: ID của user
//   - startDate: Ngày bắt đầu
//   - endDate: Ngày kết thúc
//
// Output:
//   - []models.Attendance: Danh sách attendance
//   - error: Lỗi nếu có
func (r *attendanceRepository) FindByUserID(userID uint, startDate, endDate time.Time) ([]models.Attendance, error) {
	var attendances []models.Attendance
	err := r.db.Preload("User").
		Where("user_id = ? AND date BETWEEN ? AND ?", userID, startDate, endDate).
		Order("date DESC").
		Find(&attendances).Error

	if err != nil {
		return nil, err
	}
	return attendances, nil
}

// FindAll lấy danh sách attendance với filter và phân trang
//
// Input:
//   - filter: Điều kiện lọc (user_id, start_date, end_date)
//   - page: Số trang
//   - limit: Số item mỗi trang
//
// Output:
//   - []models.Attendance: Danh sách attendance
//   - int64: Tổng số records
//   - error: Lỗi nếu có
func (r *attendanceRepository) FindAll(filter models.AttendanceFilter, page, limit int) ([]models.Attendance, int64, error) {
	var attendances []models.Attendance
	var total int64

	query := r.db.Model(&models.Attendance{})

	// Áp dụng filter
	if filter.UserID > 0 {
		query = query.Where("user_id = ?", filter.UserID)
	}
	if filter.StartDate != "" {
		query = query.Where("date >= ?", filter.StartDate)
	}
	if filter.EndDate != "" {
		query = query.Where("date <= ?", filter.EndDate)
	}

	// Đếm tổng
	query.Count(&total)

	// Query với pagination
	offset := (page - 1) * limit
	err := query.Preload("User").
		Order("date DESC, check_in DESC").
		Offset(offset).
		Limit(limit).
		Find(&attendances).Error

	if err != nil {
		return nil, 0, err
	}

	return attendances, total, nil
}

// Update cập nhật attendance (dùng khi check-out)
//
// Input:
//   - attendance: Con trỏ đến Attendance model với data mới
//
// Output:
//   - error: Lỗi nếu có
func (r *attendanceRepository) Update(attendance *models.Attendance) error {
	return r.db.Save(attendance).Error
}
