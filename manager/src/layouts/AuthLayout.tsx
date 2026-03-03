import { ReactNode } from 'react';

interface AuthLayoutProps {
  children: ReactNode;
}

export function AuthLayout({ children }: AuthLayoutProps) {
  return (
    <div className="min-h-screen flex items-center justify-center bg-gradient-to-br from-slate-50 via-blue-50 to-indigo-100">
      <div className="w-full max-w-md">
        <div className="mb-6 text-center">
          <h1 className="text-2xl font-bold text-gray-900">HR Management System</h1>
          <p className="text-sm text-gray-600">Đăng nhập để quản lý nhân sự</p>
        </div>
        <div className="bg-white rounded-2xl shadow-xl border border-slate-100 p-6">
          {children}
        </div>
      </div>
    </div>
  );
}
