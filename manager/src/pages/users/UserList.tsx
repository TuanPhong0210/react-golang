import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import { deleteUser, getUsers } from "../../api/user.api";
import { Button } from "../../components/Button";
import { Table } from "../../components/Table";
import { User } from "../../types/user";
import { formatDate } from "../../utils/auth";

export function UserList() {
  const navigate = useNavigate();
  const [users, setUsers] = useState<User[]>([]);
  const [loading, setLoading] = useState(false);
  const [page, setPage] = useState(1);
  const [totalPages, setTotalPages] = useState(1);

  const fetchUsers = async (targetPage = page) => {
    setLoading(true);
    try {
      const response = await getUsers(targetPage, 10);
      setUsers(response.data || []);
      setTotalPages(response.pagination.total_pages || 1);
    } finally {
      setLoading(false);
    }
  };

  const handleDelete = async (id: number) => {
    if (!window.confirm("Bạn chắc chắn muốn xóa nhân viên này?")) return;
    await deleteUser(id);
    fetchUsers();
  };

  useEffect(() => {
    fetchUsers();
  }, [page]);

  return (
    <div className="space-y-4">
      <div className="flex items-center justify-between">
        <div>
          <h3 className="text-lg font-semibold text-gray-900">Quản lý nhân viên</h3>
          <p className="text-sm text-gray-500">Danh sách nhân viên trong hệ thống</p>
        </div>
        <Button onClick={() => navigate("/users/new")}>Tạo nhân viên</Button>
      </div>

      <div className="bg-white rounded-xl border border-slate-100 shadow-sm p-4">
        <Table<User>
          columns={[
            { key: "full_name", title: "Họ tên" },
            { key: "email", title: "Email" },
            { key: "department", title: "Phòng ban" },
            { key: "role_name", title: "Vai trò" },
            {
              key: "created_at",
              title: "Ngày tạo",
              render: (record) => formatDate(record.created_at),
            },
            {
              key: "actions",
              title: "Thao tác",
              render: (record) => (
                <div className="flex gap-2">
                  <Button size="sm" variant="outline" onClick={() => navigate(`/users/${record.id}`)}>
                    Sửa
                  </Button>
                  <Button size="sm" variant="danger" onClick={() => handleDelete(record.id)}>
                    Xóa
                  </Button>
                </div>
              ),
            },
          ]}
          data={users}
          loading={loading}
        />
      </div>

      <div className="flex items-center justify-between text-sm text-gray-500">
        <span>
          Trang {page} / {totalPages}
        </span>
        <div className="flex gap-2">
          <Button size="sm" variant="outline" disabled={page <= 1} onClick={() => setPage(page - 1)}>
            Trước
          </Button>
          <Button size="sm" variant="outline" disabled={page >= totalPages} onClick={() => setPage(page + 1)}>
            Sau
          </Button>
        </div>
      </div>
    </div>
  );
}
