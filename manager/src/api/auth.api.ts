/**
 * ===========================================
 * API Layer - Auth API
 * ===========================================
 * Các hàm gọi API liên quan đến authentication
 */

import api from './axios';
import { ApiResponse, LoginResponse } from '../types/api';
import { User } from '../types/user';

/**
 * Đăng nhập
 * @param email - Email đăng nhập
 * @param password - Mật khẩu
 * @returns Token và thông tin user
 */
export const login = async (email: string, password: string): Promise<LoginResponse> => {
  const response = await api.post<ApiResponse<LoginResponse>>('/auth/login', {
    email,
    password,
  });
  return response.data.data!;
};

/**
 * Đăng xuất
 */
export const logout = async (): Promise<void> => {
  await api.post('/auth/logout');
  localStorage.removeItem('token');
  localStorage.removeItem('user');
};

/**
 * Lấy thông tin user hiện tại
 * @returns Thông tin user
 */
export const getCurrentUser = async (): Promise<User> => {
  const response = await api.get<ApiResponse<User>>('/auth/me');
  return response.data.data!;
};
