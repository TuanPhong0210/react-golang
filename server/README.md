# HR Management System - Backend

Backend API cho hệ thống quản lý nhân sự, được xây dựng với Golang và Gin framework.

## Yêu cầu

- Go 1.21+
- PostgreSQL 14+

## Cài đặt

1. **Clone và vào thư mục server:**
```bash
cd server
```

2. **Cài đặt dependencies:**
```bash
go mod tidy
```

3. **Cấu hình database:**
- Tạo database PostgreSQL tên `hrms_db`
- Copy `.env.example` thành `.env` và chỉnh sửa thông tin kết nối

4. **Chạy server:**
```bash
go run cmd/server/main.go
```

Server sẽ chạy tại `http://localhost:8080`

## Tài khoản mặc định

- Email: `admin@company.com`
- Password: `admin123`

## API Endpoints

### Authentication
- `POST /api/auth/login` - Đăng nhập
- `POST /api/auth/logout` - Đăng xuất
- `GET /api/auth/me` - Thông tin user hiện tại

### Users (Admin/Manager)
- `GET /api/users` - Danh sách users
- `GET /api/users/:id` - Chi tiết user
- `POST /api/users` - Tạo user (Admin)
- `PUT /api/users/:id` - Cập nhật user
- `DELETE /api/users/:id` - Xóa user (Admin)

### Attendance
- `POST /api/attendance/check-in` - Check in
- `POST /api/attendance/check-out` - Check out
- `GET /api/attendance/today` - Attendance hôm nay
- `GET /api/attendance/history` - Lịch sử cá nhân
- `GET /api/attendance` - Tất cả (Admin/Manager)

### Approvals
- `POST /api/approvals` - Tạo đơn
- `GET /api/approvals/my` - Đơn của tôi
- `GET /api/approvals/pending` - Đơn chờ duyệt
- `PUT /api/approvals/:id/approve` - Duyệt đơn
- `PUT /api/approvals/:id/reject` - Từ chối đơn
