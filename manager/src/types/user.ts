/**
 * ===========================================
 * Type Definitions - User
 * ===========================================
 * Định nghĩa các types liên quan đến User
 */

// Role name type
export type UserRole = 'Admin' | 'Manager' | 'Employee';

// Role trong hệ thống
export interface Role {
  id: number;
  name: UserRole;
  description: string;
}

// User response từ API
export interface User {
  id: number;
  email: string;
  full_name: string;
  phone: string;
  department: string;
  position: string;
  role_id: number;
  role_name: UserRole;
  is_active: boolean;
  created_at: string;
  updated_at: string;
}

// Input để tạo user mới
export interface CreateUserInput {
  email: string;
  password: string;
  full_name: string;
  phone?: string;
  department?: string;
  position?: string;
  role_id: number;
}

// Input để cập nhật user
export interface UpdateUserInput {
  full_name?: string;
  phone?: string;
  department?: string;
  position?: string;
  role_id?: number;
  is_active?: boolean;
}

// Input để đổi password
export interface ChangePasswordInput {
  old_password: string;
  new_password: string;
}
