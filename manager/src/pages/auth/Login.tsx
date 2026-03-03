import { ChangeEvent, FormEvent, useState } from "react";
import { useNavigate } from "react-router-dom";
import { login } from "../../api/auth.api";
import { Button } from "../../components/Button";
import { Input } from "../../components/Input";
import { useAuth } from "../../store/auth.store";

export function Login() {
  const navigate = useNavigate();
  const { login: saveAuth } = useAuth();
  const [email, setEmail] = useState("admin@company.com");
  const [password, setPassword] = useState("admin123");
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  const handleSubmit = async (event: FormEvent<HTMLFormElement>) => {
    event.preventDefault();
    setError(null);
    setLoading(true);

    try {
      const result = await login(email, password);
      console.log("🚀 ~ result:", result);
      // saveAuth(result.token, result.user);
      // navigate("/dashboard");
    } catch (err) {
      setError("Email hoặc mật khẩu không đúng");
    } finally {
      setLoading(false);
    }
  };

  return (
    <form onSubmit={handleSubmit} className="space-y-4">
      <Input label="Email" type="email" required value={email} onChange={(event: ChangeEvent<HTMLInputElement>) => setEmail(event.target.value)} placeholder="admin@company.com" />
      <Input label="Mật khẩu" type="password" required value={password} onChange={(event: ChangeEvent<HTMLInputElement>) => setPassword(event.target.value)} placeholder="••••••••" />

      {error && <p className="text-sm text-red-500">{error}</p>}

      <Button type="submit" fullWidth isLoading={loading}>
        Đăng nhập
      </Button>

      <div className="text-xs text-gray-500 text-center">Tài khoản demo: admin@company.com / admin123</div>
    </form>
  );
}
