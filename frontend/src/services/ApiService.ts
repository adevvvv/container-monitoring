import { PingStatus } from '../types/PingStatus';

export const fetchStatuses = async (etag: string): Promise<PingStatus[]> => {
  const headers: HeadersInit = {};
  if (etag) {
    headers['If-None-Match'] = etag;
  }

  const response = await fetch('http://localhost:8080/api/v1/status', { headers });

  if (response.status === 304) return [];

  if (!response.ok) {
    throw new Error(`Server error: ${response.status}`);
  }

  const data = await response.json();
  return data.data;
};

export const createStatus = async (status: Omit<PingStatus, 'id' | 'links'>, etag: string): Promise<PingStatus> => {
  const response = await fetch('http://localhost:8080/api/v1/status', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
      'If-Match': etag,
    },
    body: JSON.stringify(status),
  });

  if (!response.ok) {
    throw new Error('Error adding status');
  }

  return await response.json();
};

export const updateStatus = async (status: PingStatus, etag: string): Promise<PingStatus> => {
  const response = await fetch(`http://localhost:8080/api/v1/status/${status.id}`, {
    method: 'PUT',
    headers: {
      'Content-Type': 'application/json',
      'If-Match': etag,
    },
    body: JSON.stringify(status),
  });

  if (!response.ok) {
    throw new Error('Error updating status');
  }

  return await response.json();
};

export const deleteStatus = async (id: number, etag: string): Promise<void> => {
  const response = await fetch(`http://localhost:8080/api/v1/status/${id}`, {
    method: 'DELETE',
    headers: {
      'If-Match': etag,
    },
  });

  if (!response.ok) {
    throw new Error('Error deleting status');
  }
};
