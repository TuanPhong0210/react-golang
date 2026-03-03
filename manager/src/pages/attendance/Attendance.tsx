import { ChangeEvent, useEffect, useState } from "react";
import { checkIn, checkOut, getAllAttendances, getMyAttendanceHistory, getTodayAttendance } from "../../api/attendance.api";
import { Button } from "../../components/Button";
import { Table } from "../../components/Table";
import { Attendance } from "../../types/attendance";
import { formatDate, formatDateTime, formatTime } from "../../utils/auth";
import { usePermission } from "../../store/auth.store";

export function AttendancePage() {
  const { canManageUsers } = usePermission();
  const [today, setToday] = useState<Attendance | null>(null);
  const [history, setHistory] = useState<Attendance[]>([]);
  const [allAttendances, setAllAttendances] = useState<Attendance[]>([]);
  const [loading, setLoading] = useState(false);
  const [note, setNote] = useState("");

  const fetchToday = async () => {
    const data = await getTodayAttendance();
    setToday(data);
  };

  const fetchHistory = async () => {
    const data = await getMyAttendanceHistory();
    setHistory(data);
  };

  const fetchAll = async () => {
    if (!canManageUsers) return;
    const response = await getAllAttendances({}, 1, 10);
    setAllAttendances(response.data || []);
  };

  useEffect(() => {
    fetchToday();
    fetchHistory();
    fetchAll();
  }, [canManageUsers]);

  const handleCheckIn = async () => {
    setLoading(true);
    try {
      await checkIn({ note });
      setNote("");
      fetchToday();
      fetchHistory();
    } finally {
      setLoading(false);
    }
  };

  const handleCheckOut = async () => {
    setLoading(true);
    try {
      await checkOut({ note });
      setNote("");
      fetchToday();
      fetchHistory();
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="space-y-6">
      <div className="bg-white rounded-xl border border-slate-100 p-6 shadow-sm">
        <h3 className="text-lg font-semibold text-gray-900">Chấm công hôm nay</h3>
        <p className="text-sm text-gray-500">{formatDate(new Date().toISOString())}</p>

        <div className="mt-4 grid gap-4 md:grid-cols-3">
          <div className="rounded-lg bg-slate-50 p-4">
            <p className="text-xs text-gray-500">Check-in</p>
            <p className="text-lg font-semibold text-gray-900">{today?.check_in ? formatTime(today.check_in) : "--:--"}</p>
          </div>
          <div className="rounded-lg bg-slate-50 p-4">
            <p className="text-xs text-gray-500">Check-out</p>
            <p className="text-lg font-semibold text-gray-900">{today?.check_out ? formatTime(today.check_out) : "--:--"}</p>
          </div>
          <div className="rounded-lg bg-slate-50 p-4">
            <p className="text-xs text-gray-500">Tổng giờ</p>
            <p className="text-lg font-semibold text-gray-900">{today?.working_hours ? today.working_hours.toFixed(2) : "0"} giờ</p>
          </div>
        </div>

        <div className="mt-4 flex flex-col gap-3 md:flex-row md:items-center">
          <input
            className="flex-1 px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-primary-500"
            placeholder="Ghi chú (nếu có)"
            value={note}
            onChange={(event: ChangeEvent<HTMLInputElement>) => setNote(event.target.value)}
          />
          <div className="flex gap-2">
            <Button onClick={handleCheckIn} isLoading={loading} disabled={!!today}>
              Check-in
            </Button>
            <Button variant="success" onClick={handleCheckOut} isLoading={loading} disabled={!today || !!today.check_out}>
              Check-out
            </Button>
          </div>
        </div>
      </div>

      <div className="bg-white rounded-xl border border-slate-100 p-6 shadow-sm">
        <h3 className="text-lg font-semibold text-gray-900">Lịch sử chấm công</h3>
        <Table<Attendance>
          columns={[
            { key: "date", title: "Ngày", render: (record) => formatDate(record.date) },
            { key: "check_in", title: "Check-in", render: (record) => formatDateTime(record.check_in) },
            { key: "check_out", title: "Check-out", render: (record) => (record.check_out ? formatDateTime(record.check_out) : "-") },
            { key: "working_hours", title: "Giờ làm", render: (record) => record.working_hours.toFixed(2) },
            { key: "note", title: "Ghi chú" },
          ]}
          data={history}
        />
      </div>

      {canManageUsers && (
        <div className="bg-white rounded-xl border border-slate-100 p-6 shadow-sm">
          <h3 className="text-lg font-semibold text-gray-900">Toàn bộ chấm công</h3>
          <Table<Attendance>
            columns={[
              { key: "user_name", title: "Nhân viên" },
              { key: "date", title: "Ngày", render: (record) => formatDate(record.date) },
              { key: "check_in", title: "Check-in", render: (record) => formatTime(record.check_in) },
              { key: "check_out", title: "Check-out", render: (record) => (record.check_out ? formatTime(record.check_out) : "-") },
              { key: "working_hours", title: "Giờ làm", render: (record) => record.working_hours.toFixed(2) },
            ]}
            data={allAttendances}
          />
        </div>
      )}
    </div>
  );
}
