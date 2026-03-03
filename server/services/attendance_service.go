// ===========================================
// Package services - Attendance Service
// ===========================================
// File này chứa logic xử lý nghiệp vụ cho Attendance
// Bao gồm: check-in, check-out, xem lịch sử
// ===========================================

package services

import (
	"errors"
	"time"

	"github.com/hanbiro/hrms-server/models"
	"github.com/hanbiro/hrms-server/repositories"
)

// AttendanceService là interface định nghĩa các phương thức
// cho Attendance service
type AttendanceService interface {
	CheckIn(userID uint, note string) (*models.Attendance, error)
	CheckOut(userID uint, note string) (*models.Attendance, error)
	GetTodayAttendance(userID uint) (*models.Attendance, error)
	GetAttendanceHistory(userID uint, startDate, endDate string) ([]models.Attendance, error)
	GetAllAttendances(filter models.AttendanceFilter, page, limit int) ([]models.Attendance, int64, error)
}

// attendanceService là implementation của AttendanceService
type attendanceService struct {
	attendanceRepo repositories.AttendanceRepository
}

// NewAttendanceService tạo instance mới của AttendanceService
//
// Input:
//   - attendanceRepo: Attendance repository instance
//
// Output:
//   - AttendanceService interface
func NewAttendanceService(attendanceRepo repositories.AttendanceRepository) AttendanceService {
	return &attendanceService{attendanceRepo: attendanceRepo}
}

// CheckIn thực hiện check-in cho user
//
// Input:
//   - userID: ID của user
//   - note: Ghi chú (có thể để trống)
//
// Output:
//   - *models.Attendance: Bản ghi attendance đã tạo
//   - error: Lỗi nếu đã check-in trong ngày
func (s *attendanceService) CheckIn(userID uint, note string) (*models.Attendance, error) {
	now := time.Now()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())

	// Kiểm tra xem đã check-in trong ngày chưa
	existing, _ := s.attendanceRepo.FindByUserAndDate(userID, today)
	if existing != nil {
		return nil, errors.New("bạn đã check-in trong ngày hôm nay")
	}

	// Tạo bản ghi check-in mới
	attendance := &models.Attendance{
		UserID:  userID,
		CheckIn: now,
		Date:    today,
		Note:    note,
	}

	err := s.attendanceRepo.Create(attendance)
	if err != nil {
		return nil, errors.New("không thể tạo bản ghi check-in")
	}

	// Lấy lại với thông tin user
	attendance, _ = s.attendanceRepo.FindByID(attendance.ID)

	return attendance, nil
}

// CheckOut thực hiện check-out cho user
//
// Input:
//   - userID: ID của user
//   - note: Ghi chú bổ sung (có thể để trống)
//
// Output:
//   - *models.Attendance: Bản ghi attendance đã cập nhật
//   - error: Lỗi nếu chưa check-in hoặc đã check-out
func (s *attendanceService) CheckOut(userID uint, note string) (*models.Attendance, error) {
	now := time.Now()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())

	// Tìm bản ghi check-in của ngày hôm nay
	attendance, err := s.attendanceRepo.FindByUserAndDate(userID, today)
	if err != nil {
		return nil, errors.New("bạn chưa check-in trong ngày hôm nay")
	}

	// Kiểm tra đã check-out chưa
	if attendance.CheckOut != nil {
		return nil, errors.New("bạn đã check-out trong ngày hôm nay")
	}

	// Cập nhật check-out
	attendance.CheckOut = &now
	if note != "" {
		attendance.Note = attendance.Note + " | " + note
	}

	err = s.attendanceRepo.Update(attendance)
	if err != nil {
		return nil, errors.New("không thể cập nhật check-out")
	}

	// Lấy lại với thông tin user
	attendance, _ = s.attendanceRepo.FindByID(attendance.ID)

	return attendance, nil
}

// GetTodayAttendance lấy thông tin chấm công của ngày hôm nay
//
// Input:
//   - userID: ID của user
//
// Output:
//   - *models.Attendance: Bản ghi attendance (có thể nil nếu chưa check-in)
//   - error: Lỗi nếu có
func (s *attendanceService) GetTodayAttendance(userID uint) (*models.Attendance, error) {
	now := time.Now()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())

	attendance, err := s.attendanceRepo.FindByUserAndDate(userID, today)
	if err != nil {
		// Không có bản ghi = chưa check-in, không phải lỗi
		return nil, nil
	}

	return attendance, nil
}

// GetAttendanceHistory lấy lịch sử chấm công của user
//
// Input:
//   - userID: ID của user
//   - startDate: Ngày bắt đầu (format: YYYY-MM-DD)
//   - endDate: Ngày kết thúc (format: YYYY-MM-DD)
//
// Output:
//   - []models.Attendance: Danh sách attendance
//   - error: Lỗi nếu có
func (s *attendanceService) GetAttendanceHistory(userID uint, startDate, endDate string) ([]models.Attendance, error) {
	// Parse dates
	start, err := time.Parse("2006-01-02", startDate)
	if err != nil {
		// Mặc định 30 ngày trước
		start = time.Now().AddDate(0, 0, -30)
	}

	end, err := time.Parse("2006-01-02", endDate)
	if err != nil {
		// Mặc định ngày hôm nay
		end = time.Now()
	}

	return s.attendanceRepo.FindByUserID(userID, start, end)
}

// GetAllAttendances lấy danh sách attendance với filter và phân trang
// Dùng cho Admin/Manager xem tất cả
//
// Input:
//   - filter: Điều kiện lọc
//   - page: Số trang
//   - limit: Số item mỗi trang
//
// Output:
//   - []models.Attendance: Danh sách attendance
//   - int64: Tổng số records
//   - error: Lỗi nếu có
func (s *attendanceService) GetAllAttendances(filter models.AttendanceFilter, page, limit int) ([]models.Attendance, int64, error) {
	// Validate pagination params
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	return s.attendanceRepo.FindAll(filter, page, limit)
}
