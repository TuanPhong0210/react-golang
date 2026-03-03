import { useAuth, usePermission } from '../../store/auth.store';

export function Dashboard() {
  const { user } = useAuth();
  const { isAdmin, isManager } = usePermission();

  const cards = [
    {
      title: 'Vai trò hiện tại',
      value: user?.role_name || 'N/A',
      description: 'Quyền truy cập của bạn',
    },
    {
      title: 'Bộ phận',
      value: user?.department || 'Chưa cập nhật',
      description: 'Thông tin phòng ban',
    },
    {
      title: 'Trạng thái',
      value: user?.is_active ? 'Đang hoạt động' : 'Tạm khóa',
      description: 'Trạng thái tài khoản',
    },
  ];

  return (
    <div className="space-y-6">
      <div className="grid gap-4 md:grid-cols-3">
        {cards.map((card) => (
          <div key={card.title} className="bg-white p-4 rounded-xl border border-slate-100 shadow-sm">
            <p className="text-sm text-gray-500">{card.title}</p>
            <p className="text-2xl font-semibold text-gray-900 mt-1">{card.value}</p>
            <p className="text-xs text-gray-400 mt-1">{card.description}</p>
          </div>
        ))}
      </div>

      <div className="bg-white p-6 rounded-xl border border-slate-100 shadow-sm">
        <h3 className="text-lg font-semibold text-gray-900">Thông tin nhanh</h3>
        <p className="text-sm text-gray-600 mt-2">
          {isAdmin || isManager
            ? 'Bạn có quyền quản lý nhân viên và phê duyệt yêu cầu.'
            : 'Bạn có thể chấm công và tạo yêu cầu nghỉ phép/OT.'}
        </p>
        <div className="mt-4 grid gap-3 sm:grid-cols-2">
          <div className="rounded-lg bg-slate-50 p-4">
            <p className="text-sm font-medium text-gray-700">Hướng dẫn nhanh</p>
            <ul className="text-xs text-gray-500 mt-2 space-y-1">
              <li>• Kiểm tra chấm công tại mục Chấm công.</li>
              <li>• Gửi đơn nghỉ phép tại mục Phê duyệt.</li>
              <li>• Cập nhật thông tin nhân viên trong Quản lý nhân viên.</li>
            </ul>
          </div>
          <div className="rounded-lg bg-indigo-50 p-4">
            <p className="text-sm font-medium text-indigo-800">Trạng thái hệ thống</p>
            <p className="text-xs text-indigo-700 mt-2">
              Backend đang hoạt động. Mọi API đã sẵn sàng.
            </p>
          </div>
        </div>
      </div>
    </div>
  );
}
