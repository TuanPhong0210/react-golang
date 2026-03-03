/**
 * ===========================================
 * Utils - Auth Helpers
 * ===========================================
 * Các hàm helper cho authentication
 */

/**
 * Lấy token từ localStorage
 */
export const getToken = (): string | null => {
  return localStorage.getItem('token');
};

/**
 * Kiểm tra đã đăng nhập chưa
 */
export const isAuthenticated = (): boolean => {
  return !!getToken();
};

/**
 * Format date từ ISO string sang định dạng Việt Nam
 */
export const formatDate = (dateString: string): string => {
  const date = new Date(dateString);
  return date.toLocaleDateString('vi-VN', {
    year: 'numeric',
    month: '2-digit',
    day: '2-digit',
  });
};

/**
 * Format datetime từ ISO string
 */
export const formatDateTime = (dateString: string): string => {
  const date = new Date(dateString);
  return date.toLocaleString('vi-VN', {
    year: 'numeric',
    month: '2-digit',
    day: '2-digit',
    hour: '2-digit',
    minute: '2-digit',
  });
};

/**
 * Format time từ ISO string
 */
export const formatTime = (dateString: string): string => {
  const date = new Date(dateString);
  return date.toLocaleTimeString('vi-VN', {
    hour: '2-digit',
    minute: '2-digit',
  });
};
