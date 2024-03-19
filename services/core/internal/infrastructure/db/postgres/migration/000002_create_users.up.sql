CREATE TYPE user_role AS ENUM ('ADMIN','USER');

CREATE TABLE users (
    id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    username TEXT UNIQUE CHECK (LENGTH(username) <= 32), -- Nullable
    password TEXT NOT NULL,
    name TEXT,
    role user_role DEFAULT 'USER' NOT NULL,
    is_active BOOLEAN DEFAULT TRUE NOT NULL,
    created_at timestamptz NOT NULL DEFAULT (now()),
    updated_at timestamptz NOT NULL DEFAULT (now())
);