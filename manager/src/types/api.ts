/**
 * ===========================================
 * Type Definitions - API Response
 * ===========================================
 * Định nghĩa các types cho API response
 */

// Response chuẩn từ API
export interface ApiResponse<T = unknown> {
  success: boolean;
  message: string;
  data?: T;
  error?: string;
}

// Response có phân trang
export interface PaginatedResponse<T> {
  success: boolean;
  message: string;
  data: T[];
  pagination: Pagination;
}

// Thông tin phân trang
export interface Pagination {
  page: number;
  limit: number;
  total: number;
  total_pages: number;
}

// Login response
export interface LoginResponse {
  token: string;
  user: import('./user').User;
}
