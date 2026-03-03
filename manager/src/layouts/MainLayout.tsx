import { ReactNode } from 'react';
import { Sidebar } from '../components/Sidebar';
import { useAuth } from '../store/auth.store';

interface MainLayoutProps {
  children: ReactNode;
}

export function MainLayout({ children }: MainLayoutProps) {
  const { user } = useAuth();

  return (
    <div className="min-h-screen flex bg-slate-50">
      <Sidebar />
      <main className="flex-1">
        <header className="bg-white border-b border-slate-200 px-6 py-4">
          <div className="flex items-center justify-between">
            <div>
              <h2 className="text-lg font-semibold text-gray-900">Xin chào, {user?.full_name}</h2>
              <p className="text-sm text-gray-500">Chúc bạn một ngày làm việc hiệu quả</p>
            </div>
            <div className="text-sm text-gray-500">
              {new Date().toLocaleDateString('vi-VN', {
                weekday: 'long',
                year: 'numeric',
                month: 'long',
                day: 'numeric',
              })}
            </div>
          </div>
        </header>
        <div className="p-6">{children}</div>
      </main>
    </div>
  );
}
