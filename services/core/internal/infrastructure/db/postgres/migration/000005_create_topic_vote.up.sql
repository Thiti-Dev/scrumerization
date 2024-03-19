CREATE TABLE topic_votes (
    topic_id UUID NOT NULL REFERENCES topics(id),
    user_id UUID NOT NULL REFERENCES users(id),
    voted INTEGER NOT NULL CHECK (voted >= 0 AND voted <= 100),
    voted_description TEXT, -- optional explanation
    created_at timestamptz NOT NULL DEFAULT (now()),
    updated_at timestamptz NOT NULL DEFAULT (now()),

    CONSTRAINT topic_id_to_user_id_constraint UNIQUE (topic_id, user_id) -- Each user is allowed to vote only once
);