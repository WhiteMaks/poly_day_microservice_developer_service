CREATE TABLE root.model_sources(
    id              BIGSERIAL       NOT NULL        PRIMARY KEY,
    entity_id       BIGSERIAL       NOT NULL        UNIQUE,
    created_at      timestamptz     NOT NULL        DEFAULT now(),
    updated_at      timestamptz     NOT NULL        DEFAULT now(),
    source_base_64  TEXT
);

ALTER TABLE
	root.model_sources
ADD CONSTRAINT
	fk_entity_id
FOREIGN KEY
	(entity_id)
REFERENCES
	root.entities (id);

INSERT INTO
	root.model_sources(entity_id)
SELECT
	en.id as entity_id
FROM
	root.entities en
WHERE
	en.type = 'MODEL';