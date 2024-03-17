CREATE TABLE users (
    id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    username TEXT UNIQUE CHECK (LENGTH(username) <= 32), -- Nullable
    password VARCHAR(64) NOT NULL,
    name TEXT,
    is_active BOOLEAN DEFAULT TRUE NOT NULL,
    created_at timestamptz NOT NULL DEFAULT (now()),
    updated_at timestamptz NOT NULL DEFAULT (now())
);