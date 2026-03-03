/**
 * ===========================================
 * API Layer - Approval API
 * ===========================================
 * Các hàm gọi API liên quan đến phê duyệt
 */

import api from './axios';
import { ApiResponse, PaginatedResponse } from '../types/api';
import { Approval, CreateApprovalInput, ApprovalFilter } from '../types/approval';

/**
 * Tạo đơn mới
 * @param data - Thông tin đơn
 * @returns Đơn đã tạo
 */
export const createApproval = async (data: CreateApprovalInput): Promise<Approval> => {
  const response = await api.post<ApiResponse<Approval>>('/approvals', data);
  return response.data.data!;
};

/**
 * Lấy danh sách đơn của mình
 * @returns Danh sách đơn
 */
export const getMyApprovals = async (): Promise<Approval[]> => {
  const response = await api.get<ApiResponse<Approval[]>>('/approvals/my');
  return response.data.data || [];
};

/**
 * Lấy chi tiết đơn
 * @param id - ID của đơn
 * @returns Thông tin đơn
 */
export const getApprovalById = async (id: number): Promise<Approval> => {
  const response = await api.get<ApiResponse<Approval>>(`/approvals/${id}`);
  return response.data.data!;
};

/**
 * Lấy tất cả đơn (Admin/Manager)
 * @param filter - Điều kiện lọc
 * @param page - Số trang
 * @param limit - Số item mỗi trang
 * @returns Danh sách đơn với pagination
 */
export const getAllApprovals = async (
  filter?: ApprovalFilter,
  page = 1,
  limit = 10
): Promise<PaginatedResponse<Approval>> => {
  const response = await api.get<PaginatedResponse<Approval>>('/approvals', {
    params: { ...filter, page, limit },
  });
  return response.data;
};

/**
 * Lấy đơn chờ duyệt (Admin/Manager)
 * @param page - Số trang
 * @param limit - Số item mỗi trang
 * @returns Danh sách đơn pending
 */
export const getPendingApprovals = async (
  page = 1,
  limit = 10
): Promise<PaginatedResponse<Approval>> => {
  const response = await api.get<PaginatedResponse<Approval>>('/approvals/pending', {
    params: { page, limit },
  });
  return response.data;
};

/**
 * Duyệt đơn
 * @param id - ID của đơn
 * @returns Đơn đã duyệt
 */
export const approveRequest = async (id: number): Promise<Approval> => {
  const response = await api.put<ApiResponse<Approval>>(`/approvals/${id}/approve`);
  return response.data.data!;
};

/**
 * Từ chối đơn
 * @param id - ID của đơn
 * @returns Đơn đã từ chối
 */
export const rejectRequest = async (id: number): Promise<Approval> => {
  const response = await api.put<ApiResponse<Approval>>(`/approvals/${id}/reject`);
  return response.data.data!;
};
