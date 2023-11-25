-- name: CreateModelSource :one
INSERT INTO
	root.model_sources(entity_id, source_base_64)
VALUES
	($1, $2)
RETURNING *;

-- name: DeleteModelSourceByEntityId :exec
DELETE FROM
	root.model_sources
WHERE
	entity_id = $1;

-- name: UpdateModelSourceByEntityId :exec
UPDATE
	root.model_sources
SET
	source_base_64 = $2,
	updated_at = now()
WHERE
	entity_id = $1;

-- name: GetModelSourceByEntityId :one
SELECT
	*
FROM
	root.model_sources
WHERE
	entity_id = $1;

-- name: CreateShaderSource :one
INSERT INTO
	root.shader_sources(entity_id, vertex_shader, geometry_shader, fragment_shader)
VALUES
	($1, $2, $3, $4)
RETURNING *;

-- name: DeleteShaderSourceByEntityId :exec
DELETE FROM
	root.shader_sources
WHERE
	entity_id = $1;

-- name: UpdateShaderSourceByEntityId :exec
UPDATE
	root.shader_sources
SET
	vertex_shader = $2,
	geometry_shader = $3,
	fragment_shader = $4,
	updated_at = now()
WHERE
	entity_id = $1;

-- name: GetShaderSourceByEntityId :one
SELECT
	*
FROM
	root.shader_sources
WHERE
	entity_id = $1;