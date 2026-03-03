/**
 * ===========================================
 * Component - Table
 * ===========================================
 * Table component đơn giản với Tailwind
 */

import { ReactNode } from 'react';

interface Column<T> {
  key: string;
  title: string;
  render?: (record: T) => ReactNode;
  className?: string;
}

interface TableProps<T> {
  columns: Column<T>[];
  data: T[];
  loading?: boolean;
  emptyText?: string;
}

export function Table<T extends { id: number | string }>({
  columns,
  data,
  loading = false,
  emptyText = 'Không có dữ liệu',
}: TableProps<T>) {
  if (loading) {
    return (
      <div className="flex items-center justify-center py-12">
        <svg className="animate-spin h-8 w-8 text-primary-600" fill="none" viewBox="0 0 24 24">
          <circle className="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" strokeWidth="4" />
          <path className="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4z" />
        </svg>
      </div>
    );
  }

  return (
    <div className="overflow-x-auto">
      <table className="min-w-full divide-y divide-gray-200">
        {/* Header */}
        <thead className="bg-gray-50">
          <tr>
            {columns.map((column) => (
              <th
                key={column.key}
                className={`
                  px-4 py-3 text-left text-xs font-medium
                  text-gray-500 uppercase tracking-wider
                  ${column.className || ''}
                `}
              >
                {column.title}
              </th>
            ))}
          </tr>
        </thead>

        {/* Body */}
        <tbody className="bg-white divide-y divide-gray-200">
          {data.length === 0 ? (
            <tr>
              <td
                colSpan={columns.length}
                className="px-4 py-8 text-center text-gray-500"
              >
                {emptyText}
              </td>
            </tr>
          ) : (
            data.map((record) => (
              <tr key={record.id} className="hover:bg-gray-50">
                {columns.map((column) => (
                  <td
                    key={`${record.id}-${column.key}`}
                    className={`px-4 py-3 text-sm text-gray-900 ${column.className || ''}`}
                  >
                    {column.render
                      ? column.render(record)
                      : (record as Record<string, unknown>)[column.key] as ReactNode}
                  </td>
                ))}
              </tr>
            ))
          )}
        </tbody>
      </table>
    </div>
  );
}
