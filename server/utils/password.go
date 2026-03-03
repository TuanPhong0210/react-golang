// ===========================================
// Package utils - Password Utilities
// ===========================================
// File này chứa các hàm liên quan đến password
// Sử dụng bcrypt để hash và validate password
// ===========================================

package utils

import (
	"golang.org/x/crypto/bcrypt"
)

// DefaultCost là độ phức tạp mặc định của bcrypt
// Giá trị cao hơn = bảo mật hơn nhưng chậm hơn
const DefaultCost = 10

// HashPassword hash password với bcrypt
//
// Input:
//   - password: Password dạng plain text
//
// Output:
//   - string: Password đã được hash
//   - error: Lỗi nếu có
//
// Ví dụ:
//
//	hashed, err := HashPassword("mypassword123")
func HashPassword(password string) (string, error) {
	// Chuyển password thành bytes và hash với bcrypt
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// CheckPassword so sánh password plain text với password đã hash
//
// Input:
//   - password: Password dạng plain text cần kiểm tra
//   - hashedPassword: Password đã được hash trong database
//
// Output:
//   - bool: true nếu password khớp, false nếu không
//
// Ví dụ:
//
//	isValid := CheckPassword("mypassword123", hashedFromDB)
func CheckPassword(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
