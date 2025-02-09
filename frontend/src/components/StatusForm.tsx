import React, { FC } from 'react';
import { PingStatus } from '../types/PingStatus';

interface StatusFormProps {
  status: Omit<PingStatus, 'id' | 'links'>;
  onChange: (e: React.ChangeEvent<HTMLInputElement>) => void;
  onSubmit: (e: React.FormEvent) => void;
  onCancel: () => void;
  isEditMode: boolean;
}

const StatusForm: FC<StatusFormProps> = ({ status, onChange, onSubmit, onCancel, isEditMode }) => {
  return (
    <form onSubmit={onSubmit}>
      <div className="form-group">
        <label>IP Address:</label>
        <input type="text" name="ip" value={status.ip} onChange={onChange} />
      </div>
      <div className="form-group">
        <label>Ping Time (ms):</label>
        <input type="number" name="ping_time" value={status.ping_time} onChange={onChange} />
      </div>
      <div className="form-group">
        <label>Last Successful Ping:</label>
        <input type="datetime-local" name="last_success" value={status.last_success} onChange={onChange} />
      </div>
      <div className="form-actions">
        <button type="submit">{isEditMode ? 'Update' : 'Add'}</button>
        {isEditMode && <button type="button" onClick={onCancel}>Cancel</button>}
      </div>
    </form>
  );
};

export default StatusForm;
