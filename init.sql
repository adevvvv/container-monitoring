CREATE TABLE IF NOT EXISTS ping_status (
    id SERIAL PRIMARY KEY,
    ip TEXT NOT NULL CHECK (ip ~ '^\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}$'),
    ping_time DOUBLE PRECISION NOT NULL,
    last_success TIMESTAMP WITH TIME ZONE NOT NULL
);
