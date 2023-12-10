-- name: CreateShaderProgram :one
INSERT INTO
	root.shader_programs(name, user_id, release_id, shader_entity_id)
VALUES
	($1, $2, $3, $4)
RETURNING *;

-- name: GetShaderProgramsByReleaseId :many
SELECT
	*
FROM
	root.shader_programs
WHERE
	release_id = $1;

-- name: GetShaderProgramById :one
SELECT
	*
FROM
	root.shader_programs
WHERE
	id = $1;

-- name: DeleteShaderProgramById :exec
DELETE FROM
	root.shader_programs
WHERE
	id = $1;