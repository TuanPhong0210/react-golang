// ===========================================
// Package repositories - User Repository
// ===========================================
// File này chứa các hàm thao tác database cho User
// Sử dụng GORM để query và manipulate dữ liệu
// ===========================================

package repositories

import (
	"github.com/hanbiro/hrms-server/models"
	"gorm.io/gorm"
)

// UserRepository là interface định nghĩa các phương thức
// cho User repository
type UserRepository interface {
	Create(user *models.User) error
	FindByID(id uint) (*models.User, error)
	FindByEmail(email string) (*models.User, error)
	FindAll(page, limit int) ([]models.User, int64, error)
	Update(user *models.User) error
	Delete(id uint) error
}

// userRepository là implementation của UserRepository
type userRepository struct {
	db *gorm.DB
}

// NewUserRepository tạo instance mới của UserRepository
//
// Input:
//   - db: GORM database instance
//
// Output:
//   - UserRepository interface
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

// Create tạo user mới trong database
//
// Input:
//   - user: Con trỏ đến User model
//
// Output:
//   - error: Lỗi nếu có (vd: email trùng)
func (r *userRepository) Create(user *models.User) error {
	return r.db.Create(user).Error
}

// FindByID tìm user theo ID
//
// Input:
//   - id: ID của user cần tìm
//
// Output:
//   - *models.User: User nếu tìm thấy
//   - error: Lỗi nếu không tìm thấy hoặc lỗi DB
func (r *userRepository) FindByID(id uint) (*models.User, error) {
	var user models.User
	// Preload Role để lấy thông tin role
	err := r.db.Preload("Role").First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// FindByEmail tìm user theo email
//
// Input:
//   - email: Email của user cần tìm
//
// Output:
//   - *models.User: User nếu tìm thấy
//   - error: Lỗi nếu không tìm thấy
func (r *userRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.db.Preload("Role").Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// FindAll lấy danh sách tất cả users với phân trang
//
// Input:
//   - page: Số trang (bắt đầu từ 1)
//   - limit: Số item mỗi trang
//
// Output:
//   - []models.User: Danh sách users
//   - int64: Tổng số users
//   - error: Lỗi nếu có
func (r *userRepository) FindAll(page, limit int) ([]models.User, int64, error) {
	var users []models.User
	var total int64

	// Đếm tổng số records
	r.db.Model(&models.User{}).Count(&total)

	// Tính offset
	offset := (page - 1) * limit

	// Query với pagination và sắp xếp theo created_at DESC
	err := r.db.Preload("Role").
		Order("created_at DESC").
		Offset(offset).
		Limit(limit).
		Find(&users).Error

	if err != nil {
		return nil, 0, err
	}

	return users, total, nil
}

// Update cập nhật thông tin user
//
// Input:
//   - user: Con trỏ đến User model với data mới
//
// Output:
//   - error: Lỗi nếu có
func (r *userRepository) Update(user *models.User) error {
	return r.db.Save(user).Error
}

// Delete xóa user theo ID (soft delete nếu dùng gorm.DeletedAt)
//
// Input:
//   - id: ID của user cần xóa
//
// Output:
//   - error: Lỗi nếu có
func (r *userRepository) Delete(id uint) error {
	return r.db.Delete(&models.User{}, id).Error
}
