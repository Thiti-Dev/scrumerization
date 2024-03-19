CREATE TABLE rooms (
    id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    creator_id UUID NOT NULL REFERENCES users(id) ,
    room_name VARCHAR(32),
    password TEXT,
    is_active BOOLEAN DEFAULT TRUE NOT NULL,
    created_at timestamptz NOT NULL DEFAULT (now()),
    updated_at timestamptz NOT NULL DEFAULT (now())
);