import React, { FC } from 'react';
interface StatusMessageProps {
  message: string | null;
  type: 'success' | 'error' | null;
}

const StatusMessage: FC<StatusMessageProps> = ({ message, type }) => {
  if (!message) return null;
  return <div className={`message ${type}`}>{message}</div>;
};

export default StatusMessage;
