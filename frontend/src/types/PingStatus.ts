export interface Link {
    href: string;
    rel: string;
    method: string;
  }
  
  export interface PingStatus {
    id: number;
    ip: string;
    ping_time: number;
    last_success: string;
    links: Link[];
  }
  