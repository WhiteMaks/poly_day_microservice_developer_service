CREATE TABLE root.shader_programs(
	id               BIGSERIAL   NOT NULL PRIMARY KEY,
	name             varchar(56) NOT NULL UNIQUE,
	shader_entity_id BIGSERIAL   NOT NULL,
	release_id       BIGSERIAL   NOT NULL,
	user_id          BIGSERIAL   NOT NULL,
	created_at       timestamptz NOT NULL DEFAULT now(),
	updated_at       timestamptz NOT NULL DEFAULT now()
);

ALTER TABLE
	root.shader_programs
ADD CONSTRAINT
	fk_shader_entity_id
FOREIGN KEY
	(shader_entity_id)
REFERENCES
	root.entities (id);

ALTER TABLE
	root.shader_programs
ADD CONSTRAINT
	fk_release_id
FOREIGN KEY
	(release_id)
REFERENCES
	root.releases (id);