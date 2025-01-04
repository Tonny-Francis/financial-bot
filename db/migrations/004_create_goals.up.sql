-- CREATE TABLE GOALS
CREATE TABLE IF NOT EXISTS goals (
    id TEXT PRIMARY KEY NOT NULL,
    user_id TEXT NOT NULL,
    name TEXT NOT NULL,
    value NUMERIC NOT NULL,
    collected NUMERIC NOT NULL,
    status TEXT NOT NULL,
    priority TEXT NOT NULL,
    create_at TEXT NOT NULL,
    update_at TEXT NOT NULL
);

-- CREATE INDEX FOR user_id
CREATE INDEX IF NOT EXISTS idx_user_id ON goals (user_id);

-- CREATE INDEX FOR status
CREATE INDEX IF NOT EXISTS idx_status ON goals (status);

-- CREATE INDEX FOR priority
CREATE INDEX IF NOT EXISTS idx_priority ON goals (priority);