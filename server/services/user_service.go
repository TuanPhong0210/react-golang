// ===========================================
// Package services - User Service
// ===========================================
// File này chứa logic xử lý nghiệp vụ cho User
// Bao gồm: CRUD users, change password
// ===========================================

package services

import (
	"errors"

	"github.com/hanbiro/hrms-server/models"
	"github.com/hanbiro/hrms-server/repositories"
	"github.com/hanbiro/hrms-server/utils"
)

// UserService là interface định nghĩa các phương thức
// cho User service
type UserService interface {
	CreateUser(input models.CreateUserInput) (*models.User, error)
	GetUserByID(id uint) (*models.User, error)
	GetAllUsers(page, limit int) ([]models.User, int64, error)
	UpdateUser(id uint, input models.UpdateUserInput) (*models.User, error)
	DeleteUser(id uint) error
	ChangePassword(userID uint, oldPassword, newPassword string) error
}

// userService là implementation của UserService
type userService struct {
	userRepo repositories.UserRepository
}

// NewUserService tạo instance mới của UserService
//
// Input:
//   - userRepo: User repository instance
//
// Output:
//   - UserService interface
func NewUserService(userRepo repositories.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

// CreateUser tạo user mới
//
// Input:
//   - input: Thông tin user cần tạo
//
// Output:
//   - *models.User: User đã tạo
//   - error: Lỗi nếu email đã tồn tại hoặc lỗi khác
func (s *userService) CreateUser(input models.CreateUserInput) (*models.User, error) {
	// Kiểm tra email đã tồn tại chưa
	existingUser, _ := s.userRepo.FindByEmail(input.Email)
	if existingUser != nil {
		return nil, errors.New("email đã được sử dụng")
	}

	// Hash password
	hashedPassword, err := utils.HashPassword(input.Password)
	if err != nil {
		return nil, errors.New("không thể mã hóa mật khẩu")
	}

	// Tạo user mới
	user := &models.User{
		Email:      input.Email,
		Password:   hashedPassword,
		FullName:   input.FullName,
		Phone:      input.Phone,
		Department: input.Department,
		Position:   input.Position,
		RoleID:     input.RoleID,
		IsActive:   true,
	}

	// Lưu vào database
	err = s.userRepo.Create(user)
	if err != nil {
		return nil, errors.New("không thể tạo user")
	}

	// Lấy lại user với thông tin role
	user, _ = s.userRepo.FindByID(user.ID)

	return user, nil
}

// GetUserByID lấy thông tin user theo ID
//
// Input:
//   - id: ID của user
//
// Output:
//   - *models.User: User nếu tìm thấy
//   - error: Lỗi nếu không tìm thấy
func (s *userService) GetUserByID(id uint) (*models.User, error) {
	user, err := s.userRepo.FindByID(id)
	if err != nil {
		return nil, errors.New("không tìm thấy user")
	}
	return user, nil
}

// GetAllUsers lấy danh sách tất cả users với phân trang
//
// Input:
//   - page: Số trang (bắt đầu từ 1)
//   - limit: Số item mỗi trang
//
// Output:
//   - []models.User: Danh sách users
//   - int64: Tổng số users
//   - error: Lỗi nếu có
func (s *userService) GetAllUsers(page, limit int) ([]models.User, int64, error) {
	// Validate pagination params
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	return s.userRepo.FindAll(page, limit)
}

// UpdateUser cập nhật thông tin user
//
// Input:
//   - id: ID của user cần cập nhật
//   - input: Thông tin cần cập nhật
//
// Output:
//   - *models.User: User sau khi cập nhật
//   - error: Lỗi nếu không tìm thấy hoặc lỗi khác
func (s *userService) UpdateUser(id uint, input models.UpdateUserInput) (*models.User, error) {
	// Tìm user
	user, err := s.userRepo.FindByID(id)
	if err != nil {
		return nil, errors.New("không tìm thấy user")
	}

	// Cập nhật các field nếu có giá trị
	if input.FullName != "" {
		user.FullName = input.FullName
	}
	if input.Phone != "" {
		user.Phone = input.Phone
	}
	if input.Department != "" {
		user.Department = input.Department
	}
	if input.Position != "" {
		user.Position = input.Position
	}
	if input.RoleID > 0 {
		user.RoleID = input.RoleID
	}
	if input.IsActive != nil {
		user.IsActive = *input.IsActive
	}

	// Lưu vào database
	err = s.userRepo.Update(user)
	if err != nil {
		return nil, errors.New("không thể cập nhật user")
	}

	// Lấy lại user với thông tin role
	user, _ = s.userRepo.FindByID(user.ID)

	return user, nil
}

// DeleteUser xóa user theo ID
//
// Input:
//   - id: ID của user cần xóa
//
// Output:
//   - error: Lỗi nếu không tìm thấy hoặc lỗi khác
func (s *userService) DeleteUser(id uint) error {
	// Kiểm tra user tồn tại
	_, err := s.userRepo.FindByID(id)
	if err != nil {
		return errors.New("không tìm thấy user")
	}

	// Xóa user
	return s.userRepo.Delete(id)
}

// ChangePassword đổi mật khẩu user
//
// Input:
//   - userID: ID của user
//   - oldPassword: Mật khẩu cũ
//   - newPassword: Mật khẩu mới
//
// Output:
//   - error: Lỗi nếu mật khẩu cũ không đúng hoặc lỗi khác
func (s *userService) ChangePassword(userID uint, oldPassword, newPassword string) error {
	// Tìm user
	user, err := s.userRepo.FindByID(userID)
	if err != nil {
		return errors.New("không tìm thấy user")
	}

	// Kiểm tra mật khẩu cũ
	if !utils.CheckPassword(oldPassword, user.Password) {
		return errors.New("mật khẩu cũ không đúng")
	}

	// Hash mật khẩu mới
	hashedPassword, err := utils.HashPassword(newPassword)
	if err != nil {
		return errors.New("không thể mã hóa mật khẩu")
	}

	// Cập nhật mật khẩu
	user.Password = hashedPassword
	return s.userRepo.Update(user)
}
