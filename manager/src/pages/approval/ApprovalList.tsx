import { ChangeEvent, useEffect, useState } from "react";
import { approveRequest, createApproval, getMyApprovals, getPendingApprovals, rejectRequest } from "../../api/approval.api";
import { Button } from "../../components/Button";
import { Table } from "../../components/Table";
import { Approval, ApprovalType } from "../../types/approval";
import { formatDate } from "../../utils/auth";
import { usePermission } from "../../store/auth.store";

export function ApprovalList() {
  const { canApproveRequests } = usePermission();
  const [myApprovals, setMyApprovals] = useState<Approval[]>([]);
  const [pendingApprovals, setPendingApprovals] = useState<Approval[]>([]);
  const [type, setType] = useState<ApprovalType>("leave");
  const [startDate, setStartDate] = useState("");
  const [endDate, setEndDate] = useState("");
  const [reason, setReason] = useState("");

  const fetchMyApprovals = async () => {
    const data = await getMyApprovals();
    setMyApprovals(data);
  };

  const fetchPending = async () => {
    if (!canApproveRequests) return;
    const response = await getPendingApprovals(1, 10);
    setPendingApprovals(response.data || []);
  };

  useEffect(() => {
    fetchMyApprovals();
    fetchPending();
  }, [canApproveRequests]);

  const handleCreate = async () => {
    if (!startDate || !endDate || !reason) return;
    await createApproval({ type, start_date: startDate, end_date: endDate, reason });
    setStartDate("");
    setEndDate("");
    setReason("");
    fetchMyApprovals();
  };

  const handleApprove = async (id: number) => {
    await approveRequest(id);
    fetchPending();
  };

  const handleReject = async (id: number) => {
    await rejectRequest(id);
    fetchPending();
  };

  return (
    <div className="space-y-6">
      <div className="bg-white rounded-xl border border-slate-100 p-6 shadow-sm">
        <h3 className="text-lg font-semibold text-gray-900">Tạo yêu cầu</h3>
        <div className="mt-4 grid gap-4 md:grid-cols-4">
          <div>
            <label className="block text-sm font-medium text-gray-700 mb-1">Loại đơn</label>
            <select
              className="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-primary-500"
              value={type}
              onChange={(event: ChangeEvent<HTMLSelectElement>) => setType(event.target.value as ApprovalType)}
            >
              <option value="leave">Nghỉ phép</option>
              <option value="ot">Làm thêm giờ</option>
            </select>
          </div>
          <div>
            <label className="block text-sm font-medium text-gray-700 mb-1">Từ ngày</label>
            <input
              type="date"
              className="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-primary-500"
              value={startDate}
              onChange={(event: ChangeEvent<HTMLInputElement>) => setStartDate(event.target.value)}
            />
          </div>
          <div>
            <label className="block text-sm font-medium text-gray-700 mb-1">Đến ngày</label>
            <input
              type="date"
              className="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-primary-500"
              value={endDate}
              onChange={(event: ChangeEvent<HTMLInputElement>) => setEndDate(event.target.value)}
            />
          </div>
          <div className="md:col-span-1">
            <label className="block text-sm font-medium text-gray-700 mb-1">Lý do</label>
            <input
              className="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-primary-500"
              value={reason}
              onChange={(event: ChangeEvent<HTMLInputElement>) => setReason(event.target.value)}
            />
          </div>
        </div>
        <div className="mt-4 flex justify-end">
          <Button onClick={handleCreate}>Gửi yêu cầu</Button>
        </div>
      </div>

      <div className="bg-white rounded-xl border border-slate-100 p-6 shadow-sm">
        <h3 className="text-lg font-semibold text-gray-900">Đơn của tôi</h3>
        <Table<Approval>
          columns={[
            { key: "type_label", title: "Loại" },
            { key: "start_date", title: "Từ", render: (record) => formatDate(record.start_date) },
            { key: "end_date", title: "Đến", render: (record) => formatDate(record.end_date) },
            { key: "days", title: "Số ngày" },
            { key: "status_label", title: "Trạng thái" },
          ]}
          data={myApprovals}
        />
      </div>

      {canApproveRequests && (
        <div className="bg-white rounded-xl border border-slate-100 p-6 shadow-sm">
          <h3 className="text-lg font-semibold text-gray-900">Đơn chờ duyệt</h3>
          <Table<Approval>
            columns={[
              { key: "user_name", title: "Nhân viên" },
              { key: "type_label", title: "Loại" },
              { key: "start_date", title: "Từ", render: (record) => formatDate(record.start_date) },
              { key: "end_date", title: "Đến", render: (record) => formatDate(record.end_date) },
              { key: "days", title: "Số ngày" },
              { key: "reason", title: "Lý do" },
              {
                key: "actions",
                title: "Thao tác",
                render: (record) => (
                  <div className="flex gap-2">
                    <Button size="sm" variant="success" onClick={() => handleApprove(record.id)}>
                      Duyệt
                    </Button>
                    <Button size="sm" variant="danger" onClick={() => handleReject(record.id)}>
                      Từ chối
                    </Button>
                  </div>
                ),
              },
            ]}
            data={pendingApprovals}
          />
        </div>
      )}
    </div>
  );
}
