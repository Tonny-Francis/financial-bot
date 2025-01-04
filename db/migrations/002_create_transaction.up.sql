-- CREATE TABLE TRANSACTIONS
CREATE TABLE IF NOT EXISTS transactions (
    id TEXT PRIMARY KEY NOT NULL,
    user_id TEXT NOT NULL,
    description TEXT NOT NULL,
    type TEXT NOT NULL,
    payment_method TEXT NOT NULL,
    category TEXT NOT NULL,
    card TEXT,
    qtd_installments INTEGER,
    qtd_installments_paid INTEGER,
    value NUMERIC NOT NULL,
    status TEXT NOT NULL,
    create_at TEXT NOT NULL,
    update_at TEXT NOT NULL
);

-- CREATE INDEX FOR user_id
CREATE INDEX IF NOT EXISTS idx_user_id ON transactions (user_id);

-- CREATE INDEX FOR type
CREATE INDEX IF NOT EXISTS idx_type ON transactions (type);

-- CREATE INDEX FOR payment_method
CREATE INDEX IF NOT EXISTS idx_payment_method ON transactions (payment_method);

-- CREATE INDEX FOR category
CREATE INDEX IF NOT EXISTS idx_category ON transactions (category);

-- CREATE INDEX FOR card
CREATE INDEX IF NOT EXISTS idx_card ON transactions (card);

-- CREATE INDEX FOR status
CREATE INDEX IF NOT EXISTS idx_status ON transactions (status);