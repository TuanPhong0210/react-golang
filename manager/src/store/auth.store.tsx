/**
 * ===========================================
 * Auth Store
 * ===========================================
 * Quản lý state authentication với React Context
 * Lưu trữ thông tin user và token
 */

import { createContext, useContext, useState, useEffect, ReactNode } from 'react';
import { User } from '../types/user';
import { getCurrentUser } from '../api/auth.api';

// Interface cho Auth Context
interface AuthContextType {
  user: User | null;
  token: string | null;
  isLoading: boolean;
  isAuthenticated: boolean;
  login: (token: string, user: User) => void;
  logout: () => void;
  updateUser: (user: User) => void;
}

// Tạo context
const AuthContext = createContext<AuthContextType | undefined>(undefined);

// Provider component
interface AuthProviderProps {
  children: ReactNode;
}

export function AuthProvider({ children }: AuthProviderProps) {
  const [user, setUser] = useState<User | null>(null);
  const [token, setToken] = useState<string | null>(null);
  const [isLoading, setIsLoading] = useState(true);

  // Kiểm tra token trong localStorage khi app khởi động
  useEffect(() => {
    const initAuth = async () => {
      const savedToken = localStorage.getItem('token');
      const savedUser = localStorage.getItem('user');

      if (savedToken && savedUser) {
        try {
          // Gọi API để verify token và lấy user mới nhất
          const currentUser = await getCurrentUser();
          setToken(savedToken);
          setUser(currentUser);
        } catch {
          // Token không hợp lệ, xóa khỏi localStorage
          localStorage.removeItem('token');
          localStorage.removeItem('user');
        }
      }

      setIsLoading(false);
    };

    initAuth();
  }, []);

  // Hàm login - lưu token và user
  const login = (newToken: string, newUser: User) => {
    localStorage.setItem('token', newToken);
    localStorage.setItem('user', JSON.stringify(newUser));
    setToken(newToken);
    setUser(newUser);
  };

  // Hàm logout - xóa token và user
  const logout = () => {
    localStorage.removeItem('token');
    localStorage.removeItem('user');
    setToken(null);
    setUser(null);
  };

  // Hàm update user info
  const updateUser = (newUser: User) => {
    localStorage.setItem('user', JSON.stringify(newUser));
    setUser(newUser);
  };

  const value: AuthContextType = {
    user,
    token,
    isLoading,
    isAuthenticated: !!token && !!user,
    login,
    logout,
    updateUser,
  };

  return <AuthContext.Provider value={value}>{children}</AuthContext.Provider>;
}

// Custom hook để sử dụng auth context
export function useAuth() {
  const context = useContext(AuthContext);
  if (context === undefined) {
    throw new Error('useAuth must be used within an AuthProvider');
  }
  return context;
}

// Helper function kiểm tra quyền
export function usePermission() {
  const { user } = useAuth();

  const isAdmin = user?.role_name === 'Admin';
  const isManager = user?.role_name === 'Manager';
  const isEmployee = user?.role_name === 'Employee';
  const canManageUsers = isAdmin || isManager;
  const canApproveRequests = isAdmin || isManager;

  return {
    isAdmin,
    isManager,
    isEmployee,
    canManageUsers,
    canApproveRequests,
  };
}
