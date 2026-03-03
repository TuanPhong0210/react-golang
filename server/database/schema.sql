-- ===========================================
-- HR Management System - Database Schema
-- ===========================================
-- Hệ thống quản lý nhân sự với PostgreSQL
-- Bao gồm: roles, users, attendances, approvals
-- ===========================================

-- Xóa bảng cũ nếu tồn tại (theo thứ tự dependency)
DROP TABLE IF EXISTS approvals CASCADE;
DROP TABLE IF EXISTS attendances CASCADE;
DROP TABLE IF EXISTS users CASCADE;
DROP TABLE IF EXISTS roles CASCADE;

-- ===========================================
-- Bảng roles - Phân quyền người dùng
-- ===========================================
-- Chứa các vai trò: Admin, Manager, Employee
CREATE TABLE roles (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) UNIQUE NOT NULL,
    description VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- ===========================================
-- Bảng users - Thông tin nhân viên
-- ===========================================
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL, -- Hash bằng bcrypt
    full_name VARCHAR(100) NOT NULL,
    phone VARCHAR(20),
    department VARCHAR(100),
    position VARCHAR(100),
    role_id INTEGER REFERENCES roles(id) ON DELETE SET NULL,
    is_active BOOLEAN DEFAULT true,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- ===========================================
-- Bảng attendances - Chấm công
-- ===========================================
-- Lưu thông tin check-in/check-out hàng ngày
CREATE TABLE attendances (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
    check_in TIMESTAMP NOT NULL,
    check_out TIMESTAMP,
    date DATE NOT NULL,
    note VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    -- Mỗi user chỉ có 1 record chấm công mỗi ngày
    UNIQUE(user_id, date)
);

-- ===========================================
-- Bảng approvals - Phê duyệt (nghỉ phép, OT)
-- ===========================================
CREATE TABLE approvals (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
    type VARCHAR(50) NOT NULL, -- 'leave' hoặc 'ot'
    start_date DATE NOT NULL,
    end_date DATE NOT NULL,
    reason TEXT,
    status VARCHAR(20) DEFAULT 'pending', -- 'pending', 'approved', 'rejected'
    approved_by INTEGER REFERENCES users(id) ON DELETE SET NULL,
    approved_at TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- ===========================================
-- Indexes để tối ưu query
-- ===========================================
CREATE INDEX idx_users_email ON users(email);
CREATE INDEX idx_users_role ON users(role_id);
CREATE INDEX idx_attendances_user ON attendances(user_id);
CREATE INDEX idx_attendances_date ON attendances(date);
CREATE INDEX idx_approvals_user ON approvals(user_id);
CREATE INDEX idx_approvals_status ON approvals(status);
