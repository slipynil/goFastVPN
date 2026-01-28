CREATE TABLE IF NOT EXISTS peer (
    host_id SERIAL PRIMARY KEY,
    chat_id BIGINT NOT NULL,
    expires_at TIMESTAMP NOT NULL,
    FOREIGN KEY (chat_id) REFERENCES client(chat_id) ON DELETE CASCADE
);

-- Set the sequence to start from 2
ALTER SEQUENCE peer_host_id_seq RESTART WITH 2;

CREATE INDEX IF NOT EXISTS idx_peer_chat_id ON peer(chat_id);
CREATE INDEX IF NOT EXISTS idx_peer_expires_at ON peer(expires_at);
