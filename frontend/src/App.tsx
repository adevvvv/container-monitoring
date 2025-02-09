import React, { useState, useEffect } from 'react';
import StatusForm from './components/StatusForm';
import StatusTable from './components/StatusTable';
import StatusMessage from './components/StatusMessage';
import { fetchStatuses, createStatus, updateStatus, deleteStatus } from './services/ApiService';
import { validateIP } from './utils/Validation';
import { PingStatus } from './types/PingStatus';
import './styles/styles.css';

const App: React.FC = () => {
  const [statuses, setStatuses] = useState<PingStatus[]>([]);
  const [newStatus, setNewStatus] = useState<Omit<PingStatus, 'id' | 'links'>>({ ip: '127.0.0.1', ping_time: 100.0, last_success: new Date().toISOString() });
  const [editStatus, setEditStatus] = useState<PingStatus | null>(null);
  const [statusMessage, setStatusMessage] = useState<string | null>(null);
  const [messageType, setMessageType] = useState<'success' | 'error' | null>(null);
  const [sortConfig, setSortConfig] = useState<{ key: keyof PingStatus; direction: 'ascending' | 'descending' } | null>(null);
  const [etag, setEtag] = useState<string>('');

  useEffect(() => {
    fetchStatusesData();
    const interval = setInterval(fetchStatusesData, 30000);
    return () => clearInterval(interval);
  }, [etag]);

  const fetchStatusesData = async () => {
    try {
      const data = await fetchStatuses(etag);
      setStatuses(data);
    } catch (error) {
      showMessage('Error loading data', 'error');
    }
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    if (!validateIP(newStatus.ip)) {
      showMessage('Invalid IP address', 'error');
      return;
    }

    try {
      await createStatus(newStatus, etag);
      showMessage('Status successfully added!', 'success');
      fetchStatusesData();
      setNewStatus({ ip: '127.0.0.1', ping_time: 100.0, last_success: new Date().toISOString() });
    } catch (error) {
      showMessage('Error adding status', 'error');
    }
  };

  const handleUpdate = async (e: React.FormEvent) => {
    e.preventDefault();
    if (!editStatus || !validateIP(editStatus.ip)) {
      showMessage('Invalid data', 'error');
      return;
    }

    try {
      await updateStatus(editStatus, etag);
      showMessage('Status successfully updated!', 'success');
      setEditStatus(null);
      fetchStatusesData();
    } catch (error) {
      showMessage('Error updating status', 'error');
    }
  };

  const handleDelete = async (id: number) => {
    try {
      await deleteStatus(id, etag);
      showMessage('Status successfully deleted!', 'success');
      fetchStatusesData();
    } catch (error) {
      showMessage('Error deleting status', 'error');
    }
  };

  const handleSort = (key: keyof PingStatus) => {
    const direction = sortConfig?.key === key && sortConfig.direction === 'ascending' ? 'descending' : 'ascending';
    setSortConfig({ key, direction });
    setStatuses(prev => [...prev].sort((a, b) => {
      if (a[key] < b[key]) return direction === 'ascending' ? -1 : 1;
      if (a[key] > b[key]) return direction === 'ascending' ? 1 : -1;
      return 0;
    }));
  };

  const showMessage = (message: string, type: 'success' | 'error') => {
    setStatusMessage(message);
    setMessageType(type);
    setTimeout(() => {
      setStatusMessage(null);
      setMessageType(null);
    }, 3000);
  };

  return (
    <div className="container">
      <h1>Ping Status Monitoring</h1>

      <StatusForm
        status={editStatus || newStatus}
        onChange={(e) => setNewStatus(prev => ({ ...prev, [e.target.name]: e.target.value }))}
        onSubmit={editStatus ? handleUpdate : handleSubmit}
        onCancel={() => setEditStatus(null)}
        isEditMode={!!editStatus}
      />

      <StatusMessage message={statusMessage} type={messageType} />

      <StatusTable
        statuses={statuses}
        onDelete={handleDelete}
        onEdit={setEditStatus}
        onSort={handleSort}
        sortConfig={sortConfig}
      />
    </div>
  );
};

export default App;
