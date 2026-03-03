/**
 * ===========================================
 * API Layer - User API
 * ===========================================
 * Các hàm gọi API liên quan đến user management
 */

import api from './axios';
import { ApiResponse, PaginatedResponse } from '../types/api';
import { User, CreateUserInput, UpdateUserInput, ChangePasswordInput } from '../types/user';

/**
 * Lấy danh sách users
 * @param page - Số trang (default: 1)
 * @param limit - Số item mỗi trang (default: 10)
 * @returns Danh sách users với pagination
 */
export const getUsers = async (page = 1, limit = 10): Promise<PaginatedResponse<User>> => {
  const response = await api.get<PaginatedResponse<User>>('/users', {
    params: { page, limit },
  });
  return response.data;
};

/**
 * Lấy thông tin user theo ID
 * @param id - ID của user
 * @returns Thông tin user
 */
export const getUserById = async (id: number): Promise<User> => {
  const response = await api.get<ApiResponse<User>>(`/users/${id}`);
  return response.data.data!;
};

/**
 * Tạo user mới
 * @param data - Thông tin user cần tạo
 * @returns User đã tạo
 */
export const createUser = async (data: CreateUserInput): Promise<User> => {
  const response = await api.post<ApiResponse<User>>('/users', data);
  return response.data.data!;
};

/**
 * Cập nhật user
 * @param id - ID của user
 * @param data - Thông tin cần cập nhật
 * @returns User sau khi cập nhật
 */
export const updateUser = async (id: number, data: UpdateUserInput): Promise<User> => {
  const response = await api.put<ApiResponse<User>>(`/users/${id}`, data);
  return response.data.data!;
};

/**
 * Xóa user
 * @param id - ID của user cần xóa
 */
export const deleteUser = async (id: number): Promise<void> => {
  await api.delete(`/users/${id}`);
};

/**
 * Đổi mật khẩu
 * @param data - Mật khẩu cũ và mới
 */
export const changePassword = async (data: ChangePasswordInput): Promise<void> => {
  await api.put('/users/change-password', data);
};
