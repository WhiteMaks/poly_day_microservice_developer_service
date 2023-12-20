-- name: CreateShaderProgramObject :one
INSERT INTO
	root.shader_program_objects(shader_program_id, model_entity_id, user_id)
VALUES
	($1, $2, $3)
RETURNING *;

-- name: GetShaderProgramObjectsByShaderProgramId :many
SELECT
	*
FROM
	root.shader_program_objects
WHERE
	shader_program_id = $1;

-- name: GetShaderProgramObjectById :one
SELECT
	*
FROM
	root.shader_program_objects
WHERE
	id = $1;

-- name: DeleteShaderProgramObjectById :exec
DELETE FROM
	root.shader_program_objects
WHERE
	id = $1;