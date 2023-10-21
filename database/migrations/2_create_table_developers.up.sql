CREATE TABLE root.developers(
    user_id         BIGSERIAL      NOT NULL      PRIMARY KEY,
    created_at      timestamptz    NOT NULL      DEFAULT now()
);