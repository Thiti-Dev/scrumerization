CREATE TABLE topics (
    id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    room_id UUID NOT NULL REFERENCES rooms(id) ,
    name VARCHAR(60) NOT NULL,
    avg_score NUMERIC(4,1), -- total avg score generates right after the topic finishes
    is_active BOOLEAN DEFAULT TRUE NOT NULL,
    created_at timestamptz NOT NULL DEFAULT (now()),
    updated_at timestamptz NOT NULL DEFAULT (now())
);