CREATE TABLE IF NOT EXISTS peer (
    host_id SERIAL PRIMARY KEY,
    chat_id BIGINT NOT NULL,
    public_key TEXT,
    expires_at TIMESTAMP NOT NULL,
    FOREIGN KEY (chat_id) REFERENCES client(chat_id) ON DELETE CASCADE
);

-- Уникальность, но разрешает NULL (в PostgreSQL NULL ≠ NULL, поэтому можно несколько NULL)
CREATE UNIQUE INDEX IF NOT EXISTS idx_peer_public_key ON peer(public_key) WHERE public_key IS NOT NULL;

-- Set the sequence to start from 2
ALTER SEQUENCE peer_host_id_seq RESTART WITH 2;

CREATE INDEX IF NOT EXISTS idx_peer_chat_id ON peer(chat_id);
CREATE INDEX IF NOT EXISTS idx_peer_expires_at ON peer(expires_at);
