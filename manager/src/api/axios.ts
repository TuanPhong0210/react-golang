/**
 * ===========================================
 * API Layer - Axios Instance
 * ===========================================
 * Cấu hình Axios với interceptors để tự động
 * thêm token và xử lý lỗi
 */

import axios from 'axios';

// Tạo axios instance với base URL
const api = axios.create({
  baseURL: '/api', // Vite proxy sẽ forward đến backend
  headers: {
    'Content-Type': 'application/json',
  },
});

/**
 * Request Interceptor
 * Tự động thêm Authorization header với token từ localStorage
 */
api.interceptors.request.use(
  (config) => {
    const token = localStorage.getItem('token');
    if (token) {
      config.headers.Authorization = `Bearer ${token}`;
    }
    return config;
  },
  (error) => {
    return Promise.reject(error);
  }
);

/**
 * Response Interceptor
 * Xử lý các lỗi chung như 401 (Unauthorized)
 */
api.interceptors.response.use(
  (response) => {
    return response;
  },
  (error) => {
    // Nếu 401, redirect về login
    if (error.response?.status === 401) {
      localStorage.removeItem('token');
      localStorage.removeItem('user');
      window.location.href = '/login';
    }
    return Promise.reject(error);
  }
);

export default api;
