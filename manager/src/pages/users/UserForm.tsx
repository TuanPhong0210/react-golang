import { ChangeEvent, FormEvent, useEffect, useState } from "react";
import { useNavigate, useParams } from "react-router-dom";
import { createUser, getUserById, updateUser } from "../../api/user.api";
import { Button } from "../../components/Button";
import { Input } from "../../components/Input";
import { CreateUserInput, UpdateUserInput } from "../../types/user";

interface UserFormProps {
  mode: "create" | "edit";
}

const roleOptions = [
  { id: 1, name: "Admin" },
  { id: 2, name: "Manager" },
  { id: 3, name: "Employee" },
];

export function UserForm({ mode }: UserFormProps) {
  const navigate = useNavigate();
  const { id } = useParams();
  const [loading, setLoading] = useState(false);
  const [form, setForm] = useState<CreateUserInput>({
    email: "",
    password: "",
    full_name: "",
    phone: "",
    department: "",
    position: "",
    role_id: 3,
  });

  useEffect(() => {
    if (mode === "edit" && id) {
      getUserById(Number(id)).then((user) => {
        setForm({
          email: user.email,
          password: "",
          full_name: user.full_name,
          phone: user.phone || "",
          department: user.department || "",
          position: user.position || "",
          role_id: user.role_id,
        });
      });
    }
  }, [mode, id]);

  const handleSubmit = async (event: FormEvent<HTMLFormElement>) => {
    event.preventDefault();
    setLoading(true);

    try {
      if (mode === "create") {
        await createUser(form);
      } else if (id) {
        const payload: UpdateUserInput = {
          full_name: form.full_name,
          phone: form.phone,
          department: form.department,
          position: form.position,
          role_id: form.role_id,
        };
        await updateUser(Number(id), payload);
      }
      navigate("/users");
    } finally {
      setLoading(false);
    }
  };

  return (
    <form onSubmit={handleSubmit} className="space-y-4 bg-white rounded-xl border border-slate-100 p-6 shadow-sm">
      <h3 className="text-lg font-semibold text-gray-900">{mode === "create" ? "Tạo nhân viên" : "Cập nhật nhân viên"}</h3>

      <div className="grid gap-4 md:grid-cols-2">
        <Input label="Họ tên" required value={form.full_name} onChange={(event: ChangeEvent<HTMLInputElement>) => setForm({ ...form, full_name: event.target.value })} />
        <Input label="Email" type="email" required value={form.email} onChange={(event: ChangeEvent<HTMLInputElement>) => setForm({ ...form, email: event.target.value })} disabled={mode === "edit"} />
        {mode === "create" && (
          <Input label="Mật khẩu" type="password" required value={form.password} onChange={(event: ChangeEvent<HTMLInputElement>) => setForm({ ...form, password: event.target.value })} />
        )}
        <Input label="Số điện thoại" value={form.phone} onChange={(event: ChangeEvent<HTMLInputElement>) => setForm({ ...form, phone: event.target.value })} />
        <Input label="Phòng ban" value={form.department} onChange={(event: ChangeEvent<HTMLInputElement>) => setForm({ ...form, department: event.target.value })} />
        <Input label="Chức vụ" value={form.position} onChange={(event: ChangeEvent<HTMLInputElement>) => setForm({ ...form, position: event.target.value })} />
        <div>
          <label className="block text-sm font-medium text-gray-700 mb-1">Vai trò</label>
          <select
            className="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-primary-500"
            value={form.role_id}
            onChange={(event: ChangeEvent<HTMLSelectElement>) => setForm({ ...form, role_id: Number(event.target.value) })}
          >
            {roleOptions.map((role) => (
              <option key={role.id} value={role.id}>
                {role.name}
              </option>
            ))}
          </select>
        </div>
      </div>

      <div className="flex justify-end gap-2">
        <Button type="button" variant="outline" onClick={() => navigate("/users")}>
          Hủy
        </Button>
        <Button type="submit" isLoading={loading}>
          Lưu
        </Button>
      </div>
    </form>
  );
}
