// ===========================================
// Package services - Auth Service
// ===========================================
// File này chứa logic xử lý authentication
// Bao gồm: login, validate token
// ===========================================

package services

import (
	"errors"

	"github.com/hanbiro/hrms-server/models"
	"github.com/hanbiro/hrms-server/repositories"
	"github.com/hanbiro/hrms-server/utils"
)

// AuthService là interface định nghĩa các phương thức
// cho authentication
type AuthService interface {
	Login(email, password string) (*models.User, string, error)
	GetCurrentUser(userID uint) (*models.User, error)
}

// authService là implementation của AuthService
type authService struct {
	userRepo repositories.UserRepository
}

// NewAuthService tạo instance mới của AuthService
//
// Input:
//   - userRepo: User repository instance
//
// Output:
//   - AuthService interface
func NewAuthService(userRepo repositories.UserRepository) AuthService {
	return &authService{userRepo: userRepo}
}

// LoginInput là struct chứa thông tin đăng nhập
type LoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// LoginResponse là struct trả về sau khi đăng nhập thành công
type LoginResponse struct {
	Token string              `json:"token"`
	User  models.UserResponse `json:"user"`
}

// Login xác thực user và trả về JWT token
//
// Input:
//   - email: Email đăng nhập
//   - password: Password (plain text)
//
// Output:
//   - *models.User: Thông tin user nếu đăng nhập thành công
//   - string: JWT token
//   - error: Lỗi nếu email/password không đúng
func (s *authService) Login(email, password string) (*models.User, string, error) {
	// Tìm user theo email
	user, err := s.userRepo.FindByEmail(email)
	if err != nil {
		return nil, "", errors.New("email hoặc mật khẩu không đúng")
	}

	// Kiểm tra user có active không
	if !user.IsActive {
		return nil, "", errors.New("tài khoản đã bị vô hiệu hóa")
	}

	// Kiểm tra password
	if !utils.CheckPassword(password, user.Password) {
		return nil, "", errors.New("email hoặc mật khẩu không đúng" + password + user.Password)
	}

	// Tạo JWT token
	token, err := utils.GenerateToken(user.ID, user.Email, user.RoleID, user.Role.Name)
	if err != nil {
		return nil, "", errors.New("không thể tạo token")
	}

	return user, token, nil
}

// GetCurrentUser lấy thông tin user hiện tại từ token
//
// Input:
//   - userID: ID của user (lấy từ JWT claims)
//
// Output:
//   - *models.User: Thông tin user
//   - error: Lỗi nếu không tìm thấy user
func (s *authService) GetCurrentUser(userID uint) (*models.User, error) {
	user, err := s.userRepo.FindByID(userID)
	if err != nil {
		return nil, errors.New("không tìm thấy user")
	}

	if !user.IsActive {
		return nil, errors.New("tài khoản đã bị vô hiệu hóa")
	}

	return user, nil
}
