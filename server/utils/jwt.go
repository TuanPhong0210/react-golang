// ===========================================
// Package utils - JWT Token Utilities
// ===========================================
// File này chứa các hàm liên quan đến JWT token
// Bao gồm: tạo token, parse token, validate token
// ===========================================

package utils

import (
	"errors"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/hanbiro/hrms-server/config"
)

// JWTClaims là custom claims cho JWT token
// Chứa thông tin user cần thiết để xác thực
type JWTClaims struct {
	UserID   uint   `json:"user_id"`
	Email    string `json:"email"`
	RoleID   uint   `json:"role_id"`
	RoleName string `json:"role_name"`
	jwt.RegisteredClaims
}

// GenerateToken tạo JWT token cho user
//
// Input:
//   - userID: ID của user
//   - email: Email của user
//   - roleID: ID của role
//   - roleName: Tên role (Admin, Manager, Employee)
//
// Output:
//   - string: JWT token
//   - error: Lỗi nếu có
func GenerateToken(userID uint, email string, roleID uint, roleName string) (string, error) {
	// Lấy secret key từ config
	secretKey := []byte(config.AppConfig.JWTSecret)

	// Lấy thời gian hết hạn từ config
	expiryHours, err := strconv.Atoi(config.AppConfig.JWTExpiryHours)
	if err != nil {
		expiryHours = 24 // Mặc định 24 giờ
	}

	// Tạo claims
	claims := JWTClaims{
		UserID:   userID,
		Email:    email,
		RoleID:   roleID,
		RoleName: roleName,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(expiryHours) * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "hrms-server",
		},
	}

	// Tạo token với thuật toán HS256
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Ký token với secret key
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ParseToken parse và validate JWT token
//
// Input:
//   - tokenString: JWT token string
//
// Output:
//   - *JWTClaims: Claims nếu token hợp lệ
//   - error: Lỗi nếu token không hợp lệ
func ParseToken(tokenString string) (*JWTClaims, error) {
	// Lấy secret key từ config
	secretKey := []byte(config.AppConfig.JWTSecret)

	// Parse token
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Kiểm tra thuật toán signing
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	// Lấy claims từ token
	if claims, ok := token.Claims.(*JWTClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}

// ValidateToken kiểm tra token có hợp lệ không
//
// Input:
//   - tokenString: JWT token string
//
// Output:
//   - bool: true nếu hợp lệ, false nếu không
func ValidateToken(tokenString string) bool {
	_, err := ParseToken(tokenString)
	return err == nil
}
