CREATE TABLE root.entities(
    id                  BIGSERIAL       NOT NULL        PRIMARY KEY,
    name                VARCHAR(35)     NOT NULL,
    user_id             BIGSERIAL       NOT NULL,
    type                VARCHAR(35)     NOT NULL,
    created_at          timestamptz     NOT NULL        DEFAULT now(),
    updated_at          timestamptz     NOT NULL        DEFAULT now(),
    parent_entity_id    BIGSERIAL
);

ALTER TABLE root.entities
	ADD UNIQUE (name, parent_entity_id)