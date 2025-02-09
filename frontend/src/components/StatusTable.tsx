import React, { FC } from 'react';
import { PingStatus } from '../types/PingStatus';
interface StatusTableProps {
  statuses: PingStatus[];
  onDelete: (id: number) => void;
  onEdit: (status: PingStatus) => void;
  onSort: (key: keyof PingStatus) => void;
  sortConfig: { key: keyof PingStatus; direction: 'ascending' | 'descending' } | null;
}

const StatusTable: FC<StatusTableProps> = ({ statuses, onDelete, onEdit, onSort, sortConfig }) => {
  return (
    <table className="status-table">
      <thead>
        <tr>
          {['id', 'ip', 'ping_time', 'last_success', 'actions'].map((header) => (
            <th
              key={header}
              onClick={() => header !== 'actions' ? onSort(header as keyof PingStatus) : null}
              className={header !== 'actions' && sortConfig?.key === header ? `sortable ${sortConfig.direction}` : ''}
            >
              {header === 'ping_time' ? 'Ping (ms)' : 
               header === 'last_success' ? 'Last Successful' : 
               header === 'actions' ? 'Actions' : header.toUpperCase()}
            </th>
          ))}
        </tr>
      </thead>
      <tbody>
        {statuses.map((status) => (
          <tr key={status.id}>
            <td>{status.id}</td>
            <td>{status.ip}</td>
            <td>{status.ping_time.toFixed(2)}</td>
            <td>{new Date(status.last_success).toLocaleString()}</td>
            <td className="actions">
              <button className="edit-btn" onClick={() => onEdit(status)}>↻</button>
              <button className="delete-btn" onClick={() => onDelete(status.id)}>❌</button>
            </td>
          </tr>
        ))}
      </tbody>
    </table>
  );
};

export default StatusTable;
