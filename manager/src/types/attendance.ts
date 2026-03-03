/**
 * ===========================================
 * Type Definitions - Attendance
 * ===========================================
 * Định nghĩa các types liên quan đến chấm công
 */

// Attendance response từ API
export interface Attendance {
  id: number;
  user_id: number;
  user_name: string;
  check_in: string;
  check_out: string | null;
  date: string;
  note: string;
  working_hours: number;
  created_at: string;
}

// Input để check-in/check-out
export interface CheckInInput {
  note?: string;
}

export interface CheckOutInput {
  note?: string;
}

// Filter để lọc attendance
export interface AttendanceFilter {
  user_id?: number;
  start_date?: string;
  end_date?: string;
}
