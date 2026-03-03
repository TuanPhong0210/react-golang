/**
 * ===========================================
 * API Layer - Attendance API
 * ===========================================
 * Các hàm gọi API liên quan đến chấm công
 */

import api from './axios';
import { ApiResponse, PaginatedResponse } from '../types/api';
import { Attendance, CheckInInput, CheckOutInput, AttendanceFilter } from '../types/attendance';

/**
 * Check in
 * @param data - Ghi chú (optional)
 * @returns Bản ghi attendance
 */
export const checkIn = async (data?: CheckInInput): Promise<Attendance> => {
  const response = await api.post<ApiResponse<Attendance>>('/attendance/check-in', data || {});
  return response.data.data!;
};

/**
 * Check out
 * @param data - Ghi chú bổ sung (optional)
 * @returns Bản ghi attendance đã cập nhật
 */
export const checkOut = async (data?: CheckOutInput): Promise<Attendance> => {
  const response = await api.post<ApiResponse<Attendance>>('/attendance/check-out', data || {});
  return response.data.data!;
};

/**
 * Lấy attendance hôm nay
 * @returns Bản ghi attendance (có thể null)
 */
export const getTodayAttendance = async (): Promise<Attendance | null> => {
  const response = await api.get<ApiResponse<Attendance>>('/attendance/today');
  return response.data.data || null;
};

/**
 * Lấy lịch sử chấm công của mình
 * @param startDate - Ngày bắt đầu (YYYY-MM-DD)
 * @param endDate - Ngày kết thúc (YYYY-MM-DD)
 * @returns Danh sách attendance
 */
export const getMyAttendanceHistory = async (
  startDate?: string,
  endDate?: string
): Promise<Attendance[]> => {
  const response = await api.get<ApiResponse<Attendance[]>>('/attendance/history', {
    params: { start_date: startDate, end_date: endDate },
  });
  return response.data.data || [];
};

/**
 * Lấy tất cả attendance (Admin/Manager)
 * @param filter - Điều kiện lọc
 * @param page - Số trang
 * @param limit - Số item mỗi trang
 * @returns Danh sách attendance với pagination
 */
export const getAllAttendances = async (
  filter?: AttendanceFilter,
  page = 1,
  limit = 10
): Promise<PaginatedResponse<Attendance>> => {
  const response = await api.get<PaginatedResponse<Attendance>>('/attendance', {
    params: { ...filter, page, limit },
  });
  return response.data;
};
