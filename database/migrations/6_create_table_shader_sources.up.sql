CREATE TABLE root.shader_sources(
    id                  BIGSERIAL       NOT NULL        PRIMARY KEY,
    entity_id           BIGSERIAL       NOT NULL        UNIQUE,
    created_at          timestamptz     NOT NULL        DEFAULT now(),
    updated_at          timestamptz     NOT NULL        DEFAULT now(),
    vertex_shader       TEXT,
    geometry_shader     TEXT,
    fragment_shader     TEXT
);

ALTER TABLE
	root.shader_sources
ADD CONSTRAINT
	fk_entity_id
FOREIGN KEY
	(entity_id)
REFERENCES
	root.entities (id);

INSERT INTO
	root.shader_sources(entity_id)
SELECT
	en.id as entity_id
FROM
	root.entities en
WHERE
	en.type = 'SHADER';