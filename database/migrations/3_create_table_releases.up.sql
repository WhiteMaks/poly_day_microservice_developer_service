CREATE TABLE root.releases(
    id              BIGSERIAL       NOT NULL        PRIMARY KEY,
    name            VARCHAR(17)     NOT NULL        UNIQUE,
    user_id         BIGSERIAL       NOT NULL,
    created_at      timestamptz     NOT NULL        DEFAULT now(),
    updated_at      timestamptz     NOT NULL        DEFAULT now()
);