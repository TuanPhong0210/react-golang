/**
 * ===========================================
 * Type Definitions - Approval
 * ===========================================
 * Định nghĩa các types liên quan đến phê duyệt
 */

// Loại đơn
export type ApprovalType = 'leave' | 'ot';

// Trạng thái đơn
export type ApprovalStatus = 'pending' | 'approved' | 'rejected';

// Approval response từ API
export interface Approval {
  id: number;
  user_id: number;
  user_name: string;
  type: ApprovalType;
  type_label: string;
  start_date: string;
  end_date: string;
  days: number;
  reason: string;
  status: ApprovalStatus;
  status_label: string;
  approved_by: number | null;
  approver_name: string;
  approved_at: string | null;
  created_at: string;
}

// Input để tạo đơn mới
export interface CreateApprovalInput {
  type: ApprovalType;
  start_date: string;
  end_date: string;
  reason: string;
}

// Filter để lọc approvals
export interface ApprovalFilter {
  user_id?: number;
  type?: ApprovalType;
  status?: ApprovalStatus;
}
