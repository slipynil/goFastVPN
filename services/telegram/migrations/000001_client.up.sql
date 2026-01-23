CREATE TABLE IF NOT EXISTS client (
    host_id SERIAL PRIMARY KEY,
    telegram_username VARCHAR(255) NOT NULL,
    telegram_id BIGINT UNIQUE NOT NULL,
    status BOOLEAN DEFAULT FALSE NOT NULL,
    expires_at TIMESTAMP NOT NULL
);

-- Set the sequence to start from 2
ALTER SEQUENCE client_host_id_seq RESTART WITH 2;

CREATE INDEX IF NOT EXISTS idx_client_telegram_id ON client(telegram_id);
CREATE INDEX IF NOT EXISTS idx_client_expires_at ON client(expires_at);
CREATE INDEX IF NOT EXISTS idx_client_status ON client(status);
