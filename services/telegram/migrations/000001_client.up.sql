CREATE TABLE IF NOT EXISTS client (
    host_id SERIAL PRIMARY KEY,
    telegram_username VARCHAR(255),
    telegram_id BIGINT,
    expires_at TIMESTAMP
);

-- Set the sequence to start from 2
ALTER SEQUENCE client_host_id_seq RESTART WITH 2;

CREATE INDEX IF NOT EXISTS idx_client_telegram_id ON client(telegram_id);
CREATE INDEX IF NOT EXISTS idx_client_expires_at ON client(expires_at);
