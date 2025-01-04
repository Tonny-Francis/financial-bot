-- CREATE TABLE ANNUAL VALUES
CREATE TABLE IF NOT EXISTS annual_values (
    id TEXT PRIMARY KEY NOT NULL,
    user_id TEXT NOT NULL,
    origem TEXT,
    bank TEXT,
    type TEXT,
    value NUMERIC NOT NULL,
    create_at TEXT NOT NULL,
    update_at TEXT NOT NULL
);

-- CREATE INDEX FOR user_id
CREATE INDEX IF NOT EXISTS idx_user_id ON annual_values (user_id);

-- CREATE INDEX FOR origem
CREATE INDEX IF NOT EXISTS idx_origem ON annual_values (origem);

-- CREATE INDEX FOR bank
CREATE INDEX IF NOT EXISTS idx_bank ON annual_values (bank);

-- CREATE INDEX FOR type
CREATE INDEX IF NOT EXISTS idx_type ON annual_values (type);