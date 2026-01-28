CREATE TABLE IF NOT EXISTS payment (
    chat_id BIGINT NOT NULL,
    invoice_payload VARCHAR(255) UNIQUE NOT NULL,
    successful_payment BOOLEAN DEFAULT FALSE NOT NULL,
    time TIMESTAMP DEFAULT NOW() NOT NULL,
    FOREIGN KEY (chat_id) REFERENCES client(chat_id) ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS idx_payment_chat_id ON payment(chat_id);
CREATE INDEX IF NOT EXISTS idx_payment_time ON payment(time);
