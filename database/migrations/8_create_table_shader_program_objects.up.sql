CREATE TABLE root.shader_program_objects(
	id                 BIGSERIAL   NOT NULL PRIMARY KEY,
	shader_program_id  BIGSERIAL   NOT NULL,
	model_entity_id    BIGSERIAL   NOT NULL,
	user_id            BIGSERIAL   NOT NULL,
	created_at         timestamptz NOT NULL DEFAULT now(),
	updated_at         timestamptz NOT NULL DEFAULT now()
);

ALTER TABLE
	root.shader_program_objects
ADD CONSTRAINT
	fk_shader_program_id
FOREIGN KEY
	(shader_program_id)
REFERENCES
	root.shader_programs (id);

ALTER TABLE
	root.shader_program_objects
ADD CONSTRAINT
	fk_model_entity_id
FOREIGN KEY
	(model_entity_id)
REFERENCES
	root.entities (id);