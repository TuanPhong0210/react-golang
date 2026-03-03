-- ===========================================
-- HR Management System - Sample Data
-- ===========================================
-- Dữ liệu mẫu để test hệ thống
-- Password mặc định: "password123" (đã hash bằng bcrypt)
-- ===========================================

-- Thêm vai trò mặc định
INSERT INTO roles (name, description) VALUES
    ('Admin', 'Quản trị viên hệ thống - Full quyền'),
    ('Manager', 'Quản lý - Duyệt đơn, quản lý nhân viên'),
    ('Employee', 'Nhân viên - Chấm công, tạo đơn');

-- Thêm users mẫu
-- Password: "admin123" và "password123" được hash bằng bcrypt
-- Bạn có thể tạo hash mới tại: https://bcrypt-generator.com/
INSERT INTO users (email, password, full_name, phone, department, position, role_id) VALUES
    -- Admin user (password: admin123)
    ('admin@company.com', '$2a$10$rCE8bXmqJHVrVQPQfpLYEOQYlCYZvQaJ9ROvS7rVhGK6zLSBH4Ehi', 'System Admin', '0901234567', 'IT', 'System Administrator', 1),
    -- Manager user (password: password123)
    ('manager@company.com', '$2a$10$N9qo8uLOickgx2ZMRZoMy.Mrq4H7V5mXqR9mQqL0m9LX5dVj3Qm2u', 'Nguyen Van A', '0901234568', 'Human Resources', 'HR Manager', 2),
    -- Employee users (password: password123)
    ('employee1@company.com', '$2a$10$N9qo8uLOickgx2ZMRZoMy.Mrq4H7V5mXqR9mQqL0m9LX5dVj3Qm2u', 'Tran Van B', '0901234569', 'Development', 'Software Engineer', 3),
    ('employee2@company.com', '$2a$10$N9qo8uLOickgx2ZMRZoMy.Mrq4H7V5mXqR9mQqL0m9LX5dVj3Qm2u', 'Le Thi C', '0901234570', 'Development', 'Frontend Developer', 3),
    ('employee3@company.com', '$2a$10$N9qo8uLOickgx2ZMRZoMy.Mrq4H7V5mXqR9mQqL0m9LX5dVj3Qm2u', 'Pham Van D', '0901234571', 'Marketing', 'Marketing Specialist', 3);

-- Thêm dữ liệu chấm công mẫu (cho ngày hôm nay và hôm qua)
INSERT INTO attendances (user_id, check_in, check_out, date, note) VALUES
    (3, CURRENT_DATE - INTERVAL '1 day' + TIME '08:30:00', CURRENT_DATE - INTERVAL '1 day' + TIME '17:30:00', CURRENT_DATE - INTERVAL '1 day', 'Normal day'),
    (4, CURRENT_DATE - INTERVAL '1 day' + TIME '08:45:00', CURRENT_DATE - INTERVAL '1 day' + TIME '18:00:00', CURRENT_DATE - INTERVAL '1 day', NULL),
    (5, CURRENT_DATE - INTERVAL '1 day' + TIME '09:00:00', CURRENT_DATE - INTERVAL '1 day' + TIME '17:45:00', CURRENT_DATE - INTERVAL '1 day', 'Late 30 minutes');

-- Thêm đơn xin nghỉ/OT mẫu
INSERT INTO approvals (user_id, type, start_date, end_date, reason, status, approved_by, approved_at) VALUES
    (3, 'leave', CURRENT_DATE + INTERVAL '7 days', CURRENT_DATE + INTERVAL '8 days', 'Nghỉ phép năm - việc gia đình', 'pending', NULL, NULL),
    (4, 'ot', CURRENT_DATE + INTERVAL '2 days', CURRENT_DATE + INTERVAL '2 days', 'Hoàn thành deadline dự án X', 'approved', 2, CURRENT_TIMESTAMP),
    (5, 'leave', CURRENT_DATE - INTERVAL '5 days', CURRENT_DATE - INTERVAL '3 days', 'Nghỉ ốm', 'approved', 2, CURRENT_TIMESTAMP - INTERVAL '6 days');
