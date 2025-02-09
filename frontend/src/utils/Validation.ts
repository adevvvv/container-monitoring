export const validateIP = (ip: string): boolean => {
    return /^(?:\d{1,3}\.){3}\d{1,3}$/.test(ip);
  };
  