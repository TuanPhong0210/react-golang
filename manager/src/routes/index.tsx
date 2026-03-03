import { ReactElement } from "react";
import { Navigate, Route, Routes } from "react-router-dom";
import { AuthLayout } from "../layouts/AuthLayout";
import { MainLayout } from "../layouts/MainLayout";
import { Login } from "../pages/auth/Login";
import { Dashboard } from "../pages/dashboard/Dashboard";
import { UserList } from "../pages/users/UserList";
import { UserForm } from "../pages/users/UserForm";
import { AttendancePage } from "../pages/attendance/Attendance";
import { ApprovalList } from "../pages/approval/ApprovalList";
import { useAuth } from "../store/auth.store";
import { UserRole } from "../types/user";

interface RequireAuthProps {
  children: ReactElement;
}

function RequireAuth({ children }: RequireAuthProps) {
  const { isAuthenticated, isLoading } = useAuth();

  if (isLoading) {
    return (
      <div className="flex min-h-screen items-center justify-center">
        <div className="text-sm text-gray-500">Đang tải...</div>
      </div>
    );
  }

  if (!isAuthenticated) {
    return <Navigate to="/login" replace />;
  }

  return children;
}

interface RequireRolesProps {
  roles: UserRole[];
  children: ReactElement;
}

function RequireRoles({ roles, children }: RequireRolesProps) {
  const { user } = useAuth();
  if (!user) return <Navigate to="/login" replace />;

  if (!roles.includes(user.role_name)) {
    return <Navigate to="/dashboard" replace />;
  }

  return children;
}

export function AppRoutes() {
  return (
    <Routes>
      <Route
        path="/"
        element={
          <RequireAuth>
            <Navigate to="/dashboard" replace />
          </RequireAuth>
        }
      />

      <Route
        path="/login"
        element={
          <AuthLayout>
            <Login />
          </AuthLayout>
        }
      />

      <Route
        path="/dashboard"
        element={
          <RequireAuth>
            <MainLayout>
              <Dashboard />
            </MainLayout>
          </RequireAuth>
        }
      />

      <Route
        path="/users"
        element={
          <RequireAuth>
            <RequireRoles roles={["Admin", "Manager"]}>
              <MainLayout>
                <UserList />
              </MainLayout>
            </RequireRoles>
          </RequireAuth>
        }
      />

      <Route
        path="/users/new"
        element={
          <RequireAuth>
            <RequireRoles roles={["Admin"]}>
              <MainLayout>
                <UserForm mode="create" />
              </MainLayout>
            </RequireRoles>
          </RequireAuth>
        }
      />

      <Route
        path="/users/:id"
        element={
          <RequireAuth>
            <RequireRoles roles={["Admin", "Manager"]}>
              <MainLayout>
                <UserForm mode="edit" />
              </MainLayout>
            </RequireRoles>
          </RequireAuth>
        }
      />

      <Route
        path="/attendance"
        element={
          <RequireAuth>
            <MainLayout>
              <AttendancePage />
            </MainLayout>
          </RequireAuth>
        }
      />

      <Route
        path="/approvals"
        element={
          <RequireAuth>
            <MainLayout>
              <ApprovalList />
            </MainLayout>
          </RequireAuth>
        }
      />

      <Route path="*" element={<Navigate to="/dashboard" replace />} />
    </Routes>
  );
}
